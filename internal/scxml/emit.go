package scxml

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/vincedupuis/transplantUML/internal/model"
)

// Emitter writes the model out as SCXML. What SCXML expresses natively is
// written natively (the connector pseudo-states as transient states, time
// triggers as a delayed <send> cancelled on exit, submachines and do
// activities as <invoke>, a local transition as type="internal"); what SCXML
// runs differently is an error, and the notes, stereotypes, invariants and
// kinds it has no element for are recorded in the tpuml extension namespace.
// Each action, a name, is a <script> holding it, and each do activity an
// <invoke> whose src is its name.
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
	if len(sm.InitialActions) > 0 {
		e.fail("the machine's initial transition: SCXML has no <initial> element on <scxml>, so its actions %s cannot run", strings.Join(sm.InitialActions, ", "))
	}
	if err := errors.Join(e.errs...); err != nil {
		return nil, e.warn, err
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
	errs    []error
}

// fail records what SCXML cannot do the way the model means it; Emit then
// fails with every such error.
func (e *emitter) fail(format string, args ...any) {
	e.errs = append(e.errs, fmt.Errorf(format, args...))
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
		e.initial(child, s)
		e.pseudo(child, s)
		e.ext(child, "stereotype", s.Stereotype)
		e.ext(child, "invariant", s.Invariant)
		if len(s.Defer) > 0 {
			e.fail("state %q: SCXML has no deferred events", s.Name)
		}
		e.note(child, s.Note)
		e.aboutNotes(child, s)
		e.datamodel(child, s.Variables)

		transitions := e.sm.OutgoingTransitions(s.Name)
		timed := delays(transitions)
		if len(s.OnEntry)+len(timed) > 0 {
			onentry := child.CreateElement("onentry")
			executable(onentry, s.OnEntry)
			for _, delay := range timed {
				send := onentry.CreateElement("send")
				send.CreateAttr("event", afterEvent(delay))
				send.CreateAttr(delayAttr(delay), delay)
				send.CreateAttr("id", timerID(s, delay))
			}
		}
		if len(s.OnExit)+len(timed) > 0 {
			onexit := child.CreateElement("onexit")
			executable(onexit, s.OnExit)
			for _, delay := range timed {
				onexit.CreateElement("cancel").CreateAttr("sendid", timerID(s, delay))
			}
		}
		e.invokes(child, s)

		for _, t := range scxmlTransitions(s, transitions) {
			tr := child.CreateElement("transition")
			switch {
			case t.After != "":
				tr.CreateAttr("event", afterEvent(t.After))
			case t.Event == "":
				setAttr(tr, "event", e.completion(s))
			default:
				tr.CreateAttr("event", t.Event)
			}
			setAttr(tr, "cond", t.Cond)
			setAttr(tr, "target", strings.Join(e.targets(t), " "))
			switch {
			case t.IsInternal():
				tr.CreateAttr("type", "internal")
			case t.IsLocal() && nestedIn(e.sm, s, t.Targets):
				tr.CreateAttr("type", "internal") // SCXML's name for a local transition
			case t.IsLocal():
				e.fail("transition %s -> %s: SCXML keeps a transition inside its source only when the source is compound and every target lies inside it", t.Source, strings.Join(t.Targets, " "))
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

// initial writes the child s starts in: as the initial attribute, or as an
// <initial> element when the initial transition has an effect.
func (e *emitter) initial(el *etree.Element, s *model.State) {
	if len(s.InitialActions) == 0 {
		setAttr(el, "initial", s.Initial)
		return
	}
	tr := el.CreateElement("initial").CreateElement("transition")
	tr.CreateAttr("target", s.Initial)
	executable(tr, s.InitialActions)
}

// completion returns the event SCXML raises when s completes, which UML's
// completion transitions wait for. A submachine state completes when the
// machine it invokes does, a compound or parallel state when it reaches its
// final state. A simple state completes on entry, when an eventless transition
// is taken, unless a do activity keeps it busy.
func (e *emitter) completion(s *model.State) string {
	switch {
	case s.Submachine != "":
		return "done.invoke." + invokeID(s)
	case s.IsParallel() || (s.IsNormal() && len(e.sm.Children(s.Name)) > 0):
		return "done.state." + s.Name
	case len(s.Do) > 0:
		e.fail("state %q: SCXML would take its completion transition without waiting for the do activity to end", s.Name)
	}
	return ""
}

// invokeID names the <invoke> of a submachine state, so that a transition can
// wait for its done.invoke event.
func invokeID(s *model.State) string { return s.Name + ".submachine" }

// reference fails on an entry or exit point of a submachine state: an
// invoked SCXML machine starts in its own initial state and reports only when
// it is done.
func (e *emitter) reference(s *model.State) {
	if s.Kind == model.EntryPoint {
		e.fail("state %q: SCXML cannot enter an invoked machine through its entry point", s.Name)
		return
	}
	e.fail("state %q: SCXML cannot leave an invoked machine through its exit point", s.Name)
}

// nestedIn reports whether every target lies inside the compound state s.
func nestedIn(sm *model.StateMachine, s *model.State, targets []string) bool {
	if len(targets) == 0 {
		return false
	}
	for _, tg := range targets {
		if !slices.Contains(sm.Ancestors(tg), s.Name) {
			return false
		}
	}
	return true
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

// pseudo tags the kinds SCXML has no element for, and fails where the SCXML
// stand-in does not behave like the UML original.
func (e *emitter) pseudo(el *etree.Element, s *model.State) {
	switch s.Kind {
	case model.Choice, model.Junction, model.Fork, model.EntryPoint, model.ExitPoint:
		// A transient state entered and left in the same step: same behaviour.
	case model.Join:
		e.fail("state %q: SCXML cannot join regions; the first region to reach it would leave the parallel state", s.Name)
	case model.Terminate:
		if s.Parent != "" || len(s.OnExit) > 0 {
			e.fail("state %q: SCXML has no terminate, and a <final> state that is not at the top level, or that has exit actions, does not end the machine the same way", s.Name)
		} else {
			e.warn.Addf("state %q: SCXML has no terminate; written as a <final> state at the top level, which ends the machine", s.Name)
		}
	default:
		return
	}
	e.ext(el, "kind", string(s.Kind))
}

// invokes writes the submachine reference and the do activities as <invoke>.
func (e *emitter) invokes(el *etree.Element, s *model.State) {
	if s.Submachine != "" {
		inv := el.CreateElement("invoke")
		inv.CreateAttr("id", invokeID(s))
		inv.CreateAttr("src", s.Submachine)
	}
	for _, do := range s.Do {
		inv := el.CreateElement("invoke")
		inv.CreateAttr("type", doType)
		inv.CreateAttr("src", do)
	}
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

// timerID names the <send> of a time trigger of s, which its <cancel> refers to.
func timerID(s *model.State, delay string) string { return s.Name + ".after." + eventSafe(delay) }

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

// executable writes each action, a name, as a <script> holding it.
func executable(el *etree.Element, list []string) {
	for _, a := range list {
		el.CreateElement("script").SetText(a)
	}
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
