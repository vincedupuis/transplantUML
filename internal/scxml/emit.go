package scxml

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/beevik/etree"
	"github.com/vincedupuis/transplantUML/internal/model"
)

// Emitter writes the model out as SCXML. What SCXML expresses natively is
// written natively (the connector pseudo-states as transient states, time
// triggers as a delayed <send> cancelled on exit, submachines and do
// activities as <invoke>); the rest is recorded in the tpuml extension
// namespace and reported as a warning, since an SCXML engine will not honour
// it. Executable content, which the parser flattens to strings, comes back as
// <script> bodies unless the string is raw XML, which is re-inserted as is.
type Emitter struct{}

func (Emitter) Emit(sm *model.StateMachine) ([]byte, model.Warnings, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	root := doc.CreateElement("scxml")

	e := &emitter{sm: sm, emitted: map[string]bool{}}
	e.datamodel(root, sm.Variables)
	e.note(root, sm.Note)
	if err := e.children(root, ""); err != nil {
		return nil, nil, err
	}
	if err := e.leftovers(); err != nil {
		return nil, nil, err
	}

	// Attributes are written in creation order; the walk above may have used
	// the extension namespace, which is declared only when that happened.
	root.CreateAttr("xmlns", namespace)
	if e.usesExt {
		root.CreateAttr("xmlns:"+extPrefix, ExtNamespace)
	}
	root.CreateAttr("version", "1.0")
	setAttr(root, "name", sm.Name)
	setAttr(root, "initial", sm.Initial)

	canonical(doc)
	doc.Indent(2)
	var buf bytes.Buffer
	if _, err := doc.WriteTo(&buf); err != nil {
		return nil, nil, fmt.Errorf("writing XML: %w", err)
	}
	return buf.Bytes(), e.warn, nil
}

// stateTags maps each kind to its SCXML element. Kinds SCXML has no element
// for become transient <state>s (or a <final> for terminate) tagged with
// tpuml:kind so that the parser reads them back.
var stateTags = map[model.StateKind]string{
	model.Normal:         "state",
	model.Parallel:       "parallel",
	model.Final:          "final",
	model.Terminate:      "final",
	model.HistoryShallow: "history",
	model.HistoryDeep:    "history",
	model.Choice:         "state",
	model.Junction:       "state",
	model.Fork:           "state",
	model.Join:           "state",
	model.EntryPoint:     "state",
	model.ExitPoint:      "state",
}

// emitter walks the flat model as a tree, recording what it wrote so that
// anything the walk cannot reach is reported rather than silently dropped.
type emitter struct {
	sm      *model.StateMachine
	emitted map[string]bool
	usesExt bool
	warn    model.Warnings
}

// children writes every direct child of parent (and their descendants) under el.
func (e *emitter) children(el *etree.Element, parent string) error {
	for _, s := range e.sm.Children(parent) {
		tag, ok := stateTags[s.Kind]
		if !ok {
			return fmt.Errorf("state %q: cannot emit unknown kind %q", s.Name, s.Kind)
		}
		e.emitted[s.Name] = true
		if e.sm.IsReference(s.Name) {
			e.reference(s)
			continue
		}
		child := el.CreateElement(tag)
		setAttr(child, "id", s.Name)
		if s.IsDeepHistory() {
			child.CreateAttr("type", "deep") // shallow is the SCXML default
		}
		setAttr(child, "initial", s.Initial)
		e.pseudo(child, s)
		e.ext(child, "stereotype", s.Stereotype)
		e.ext(child, "invariant", s.Invariant)
		if len(s.Defer) > 0 {
			e.ext(child, "defer", strings.Join(s.Defer, " "))
			e.warn.Addf("state %q: SCXML has no deferred events; written as tpuml:defer, which engines ignore", s.Name)
		}
		e.note(child, s.Note)
		e.aboutNotes(child, s)
		e.datamodel(child, s.Variables)

		transitions := e.sm.OutgoingTransitions(s.Name)
		onentry, onexit := append([]string(nil), s.OnEntry...), append([]string(nil), s.OnExit...)
		for _, delay := range delays(transitions) {
			id := s.Name + ".after." + eventSafe(delay)
			send := etree.NewElement("send")
			send.CreateAttr("event", afterEvent(delay))
			send.CreateAttr(delayAttr(delay), delay)
			send.CreateAttr("id", id)
			cancel := etree.NewElement("cancel")
			cancel.CreateAttr("sendid", id)
			onentry = append(onentry, rawXML(send))
			onexit = append(onexit, rawXML(cancel))
		}
		actions(child, "onentry", onentry)
		actions(child, "onexit", onexit)
		e.invokes(child, s)

		for _, t := range scxmlTransitions(s, transitions) {
			tr := child.CreateElement("transition")
			if t.After != "" {
				tr.CreateAttr("event", afterEvent(t.After))
			} else {
				setAttr(tr, "event", t.Event)
			}
			setAttr(tr, "cond", t.Cond)
			setAttr(tr, "target", strings.Join(e.targets(t), " "))
			switch {
			case t.IsInternal():
				tr.CreateAttr("type", "internal")
			case t.IsLocal():
				e.ext(tr, "kind", string(model.Local))
				e.warn.Addf("transition %s -> %s: SCXML has no local transitions; written as an external one tagged tpuml:kind=\"local\"", t.Source, strings.Join(t.Targets, " "))
			}
			e.note(tr, t.Note)
			executable(tr, t.Actions)
		}
		if err := e.children(child, s.Name); err != nil {
			return err
		}
	}
	return nil
}

// reference warns about an entry or exit point of a submachine state, which
// is left out: an invoked SCXML machine starts in its own initial state and
// reports only when it is done.
func (e *emitter) reference(s *model.State) {
	if s.Kind == model.EntryPoint {
		e.warn.Addf("state %q: SCXML cannot enter an invoked machine through its entry point; transitions to it enter %q instead", s.Name, s.Parent)
		return
	}
	e.warn.Addf("state %q: SCXML cannot leave an invoked machine through its exit point; the transitions leaving it are not written", s.Name)
}

// targets returns the transition's targets, with each entry point of a
// submachine state, which is not written, replaced by that state.
func (e *emitter) targets(t *model.Transition) []string {
	out := make([]string, len(t.Targets))
	for i, tg := range t.Targets {
		out[i] = tg
		if e.sm.IsReference(tg) {
			out[i] = e.sm.State(tg).Parent
		}
	}
	return out
}

// scxmlTransitions rewrites the transitions leaving s into the forms an SCXML
// engine runs the way UML means them. An engine takes the first enabled
// transition in document order, so a fork's transitions become one transition
// to all their targets, and a branch guarded by "else" loses its guard and
// moves after its siblings.
func scxmlTransitions(s *model.State, transitions []*model.Transition) []*model.Transition {
	if s.Kind == model.Fork && len(transitions) > 1 {
		fork := &model.Transition{Source: s.Name}
		for _, t := range transitions {
			fork.Targets = append(fork.Targets, t.Targets...)
			fork.Actions = append(fork.Actions, t.Actions...)
		}
		return []*model.Transition{fork}
	}
	out := make([]*model.Transition, 0, len(transitions))
	var otherwise []*model.Transition
	for _, t := range transitions {
		if t.Cond != "else" {
			out = append(out, t)
			continue
		}
		unguarded := *t
		unguarded.Cond = ""
		otherwise = append(otherwise, &unguarded)
	}
	return append(out, otherwise...)
}

// pseudo tags the kinds SCXML has no element for and warns where the SCXML
// stand-in does not behave like the UML original.
func (e *emitter) pseudo(el *etree.Element, s *model.State) {
	switch s.Kind {
	case model.Choice, model.Junction, model.Fork, model.EntryPoint, model.ExitPoint:
		// A transient state entered and left in the same step: same behaviour.
	case model.Join:
		e.warn.Addf("state %q: SCXML cannot join regions; the first region to reach it leaves the parallel state", s.Name)
	case model.Terminate:
		e.warn.Addf("state %q: SCXML has no terminate; written as a <final> state, which runs exit actions", s.Name)
	default:
		return
	}
	e.ext(el, "kind", string(s.Kind))
}

// invokes writes the submachine reference and the do activities as <invoke>.
func (e *emitter) invokes(el *etree.Element, s *model.State) {
	if s.Submachine != "" {
		el.CreateElement("invoke").CreateAttr("src", s.Submachine)
	}
	for _, do := range s.Do {
		if child, ok := parseElement(do); ok {
			el.AddChild(child)
			continue
		}
		inv := el.CreateElement("invoke")
		if src, typ, ok := parseInvoke(do); ok {
			setAttr(inv, "type", typ)
			inv.CreateAttr("src", src)
			continue
		}
		inv.CreateAttr("type", extPrefix+":do")
		inv.CreateElement("content").SetText(do)
		e.warn.Addf("state %q: SCXML cannot run the do activity %q; written as <invoke> content", s.Name, do)
	}
}

// parseInvoke reads the parser's compact form of an <invoke>: "invoke(src)"
// or "invoke(src, type)". The type is the last comma-separated part when it
// looks like a type (no spaces or parentheses); anything else is all source.
func parseInvoke(s string) (src, typ string, ok bool) {
	if !strings.HasPrefix(s, "invoke(") || !strings.HasSuffix(s, ")") {
		return "", "", false
	}
	src = strings.TrimSuffix(strings.TrimPrefix(s, "invoke("), ")")
	if i := strings.LastIndex(src, ", "); i >= 0 && !strings.ContainsAny(src[i+2:], " ()") {
		src, typ = src[:i], src[i+2:]
	}
	return src, typ, src != ""
}

func (e *emitter) datamodel(el *etree.Element, vars []model.Variable) {
	if len(vars) == 0 {
		return
	}
	dm := el.CreateElement("datamodel")
	for _, v := range vars {
		d := dm.CreateElement("data")
		d.CreateAttr("id", v.Name)
		switch {
		case strings.HasPrefix(v.Value, "src(") && strings.HasSuffix(v.Value, ")"):
			d.CreateAttr("src", strings.TrimSuffix(strings.TrimPrefix(v.Value, "src("), ")"))
		case v.Value != "":
			d.CreateAttr("expr", v.Value)
		}
	}
}

// note writes a <tpuml:note> under el and returns it, or nil when there is
// no text.
func (e *emitter) note(el *etree.Element, text string) *etree.Element {
	if text == "" {
		return nil
	}
	e.usesExt = true
	n := el.CreateElement(extPrefix + ":note")
	n.SetText(text)
	return n
}

// aboutNotes writes the notes on what the state lists beside its own note,
// each saying what it is about. The parser reads them back; see aboutNotes
// there.
func (e *emitter) aboutNotes(el *etree.Element, s *model.State) {
	about := func(what, text string) *etree.Element {
		n := e.note(el, text)
		if n != nil {
			n.CreateAttr("about", what)
		}
		return n
	}
	about("entry", s.EntryNote)
	about("exit", s.ExitNote)
	about("do", s.DoNote)
	about("invariant", s.InvariantNote)
	written := map[string]bool{}
	for _, ev := range s.Defer {
		if written[ev] {
			continue
		}
		written[ev] = true
		if n := about("defer", s.DeferNotes[ev]); n != nil {
			n.CreateAttr("event", ev)
		}
	}
}

// ext sets a tpuml extension attribute, unless the value is empty.
func (e *emitter) ext(el *etree.Element, key, value string) {
	if value != "" {
		e.usesExt = true
		el.CreateAttr(extPrefix+":"+key, value)
	}
}

// delays lists the distinct time triggers among transitions, in order.
func delays(transitions []*model.Transition) []string {
	var out []string
	seen := map[string]bool{}
	for _, t := range transitions {
		if t.After != "" && !seen[t.After] {
			seen[t.After] = true
			out = append(out, t.After)
		}
	}
	return out
}

// afterEvent names the event a time trigger's <send> raises.
func afterEvent(delay string) string { return "after." + eventSafe(delay) }

// timeValue is the CSS2 time SCXML's delay attribute takes.
var timeValue = regexp.MustCompile(`^[0-9]+(\.[0-9]+)?(ms|s)$`)

// delayAttr names the attribute a <send> carries the delay in: delay for a
// time value, delayexpr for anything else, which the engine evaluates.
func delayAttr(delay string) string {
	if timeValue.MatchString(delay) {
		return "delay"
	}
	return "delayexpr"
}

var unsafeEventChars = regexp.MustCompile(`[^A-Za-z0-9_.-]`)

func eventSafe(s string) string { return unsafeEventChars.ReplaceAllString(s, "_") }

// leftovers reports the states the walk never reached and the transitions that
// hang off them. A valid model has none; see model.Validate.
func (e *emitter) leftovers() error {
	var errs []error
	for _, s := range e.sm.States {
		if !e.emitted[s.Name] {
			errs = append(errs, fmt.Errorf("state %q: not reachable from the top level (parent %q)", s.Name, s.Parent))
		}
	}
	for i, t := range e.sm.Transitions {
		if !e.emitted[t.Source] {
			errs = append(errs, fmt.Errorf("transition #%d: unknown source %q", i, t.Source))
		}
	}
	return errors.Join(errs...)
}

// actions writes an <onentry>/<onexit> wrapper, unless there is nothing to put in it.
func actions(el *etree.Element, tag string, list []string) {
	if len(list) > 0 {
		executable(el.CreateElement(tag), list)
	}
}

// executable turns flattened action strings back into executable content:
// a string that is XML (the parser keeps unknown elements that way) is
// inserted as an element, anything else becomes a <script> body, which the
// parser reads back verbatim.
func executable(el *etree.Element, list []string) {
	for _, a := range list {
		if child, ok := parseElement(a); ok {
			el.AddChild(child)
			continue
		}
		el.CreateElement("script").SetText(a)
	}
}

func parseElement(s string) (*etree.Element, bool) {
	if !strings.HasPrefix(strings.TrimSpace(s), "<") {
		return nil, false
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromString(s); err != nil {
		return nil, false
	}
	root := doc.Root()
	return root, root != nil
}

// canonical escapes only what XML requires, leaving quotes and apostrophes
// readable. Used by the emitter and by the parser's rawXML.
func canonical(doc *etree.Document) {
	doc.WriteSettings.CanonicalText = true
	doc.WriteSettings.CanonicalAttrVal = true
}

func setAttr(el *etree.Element, key, value string) {
	if value != "" {
		el.CreateAttr(key, value)
	}
}
