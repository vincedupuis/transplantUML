// Package scxml reads and writes SCXML (State Chart XML) documents: Parser
// turns a document into the model, Emitter (emit.go) writes the model out as
// SCXML.
//
// Natively supported: nested <state>, <parallel>, <final>, <history
// type="shallow|deep">, the initial attribute and the <initial> element,
// <transition> with multiple targets and type="internal", executable content
// in <onentry>, <onexit> and <transition>, <datamodel>/<data> on the machine
// and on states, and <invoke> (a do activity, or a submachine state when it
// invokes another SCXML document).
//
// Guards and actions hold names only, which the generated code calls as
// functions (model.IsName, model.IsCondition): an action is a <script> whose
// body is a name, a do activity an <invoke> whose src is a name, and a cond or
// tpuml:invariant names joined by not, and, or and parentheses. Any other
// executable content is an error.
//
// Two SCXML idioms are recognised as UML concepts: a transient state (no
// content, only eventless transitions) is a choice when it branches on
// guards and a fork when its one transition has several targets; and a
// <send delay> in <onentry> whose event a transition of the same state
// consumes is a time trigger, after(delay).
//
// What SCXML has no element for is read from the tpuml extension namespace
// (ExtNamespace), which the W3C schema permits on every element and which
// SCXML engines ignore: the tpuml:kind attribute names any model.StateKind
// (choice, junction, fork, join, entry-point, exit-point, terminate, or
// normal to suppress the idiom detection) or, on a transition, local (the
// same as SCXML's type="internal" with targets nested in the source);
// tpuml:defer lists deferred events; tpuml:invariant and tpuml:stereotype
// annotate a state; and a <tpuml:note> child holds a note, which on a state
// may say with about="entry|exit|do|invariant|defer" (and event="…" for
// defer) that it describes one of the state's behaviours, its invariant or a
// deferred event instead.
package scxml

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/vincedupuis/transplantUML/internal/model"
)

const (
	namespace = "http://www.w3.org/2005/07/scxml"

	// ExtNamespace is the namespace of the tpuml extension attributes and
	// elements that carry what SCXML cannot express.
	ExtNamespace = "https://github.com/vincedupuis/transplantUML"
	extPrefix    = "tpuml"
)

type Parser struct{}

// Parse reads an SCXML document. Warnings name the elements the model has no
// place for and that were therefore dropped.
func (Parser) Parse(src []byte) (*model.StateMachine, model.Warnings, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(src); err != nil {
		return nil, nil, fmt.Errorf("parsing XML: %w", err)
	}
	root := doc.SelectElement("scxml")
	if root == nil {
		return nil, nil, errors.New("no <scxml> root element found")
	}

	p := &parser{sm: &model.StateMachine{
		Name:        root.SelectAttrValue("name", ""),
		States:      make([]*model.State, 0),
		Transitions: make([]*model.Transition, 0),
	}}
	p.sm.Initial, _ = initialOf(root) // <scxml> takes no <initial> element, so no actions
	p.sm.Variables = p.datamodel(root)
	p.sm.Note = note(root)
	p.walk(root, "")
	localTransitions(p.sm)
	for _, c := range root.ChildElements() {
		if _, ok := stateKinds[c.Tag]; !ok && !known(c, "datamodel") {
			p.unsupported(c, "at the top level")
		}
	}
	if err := errors.Join(p.errs...); err != nil {
		return nil, p.warn, err
	}
	return p.sm, p.warn, nil
}

type parser struct {
	sm   *model.StateMachine
	warn model.Warnings
	errs []error
}

func (p *parser) fail(format string, args ...any) {
	p.errs = append(p.errs, fmt.Errorf(format, args...))
}

// unsupported reports an element the model has no place for. An SCXML
// element would change what the machine does, so it is an error. An element
// of another namespace, which SCXML engines ignore, is dropped with a warning,
// in case it was meant to be an extension element.
func (p *parser) unsupported(el *etree.Element, where string) {
	if uri := el.NamespaceURI(); uri == "" || uri == namespace {
		p.fail("<%s> %s is not supported", el.Tag, where)
		return
	}
	p.warn.Addf("<%s:%s> %s belongs to %q, which SCXML engines ignore; it was dropped", el.Space, el.Tag, where, el.NamespaceURI())
}

var stateKinds = map[string]model.StateKind{
	"state":    model.Normal,
	"parallel": model.Parallel,
	"final":    model.Final,
	"history":  model.HistoryShallow,
}

// walk appends every state-like child of el (and their descendants) to sm,
// together with the transitions declared directly under them.
func (p *parser) walk(el *etree.Element, parent string) {
	for _, child := range el.ChildElements() {
		if _, ok := stateKinds[child.Tag]; ok {
			p.state(child, parent)
		}
	}
}

func (p *parser) state(el *etree.Element, parent string) {
	kind := stateKinds[el.Tag]
	if kind == model.HistoryShallow && el.SelectAttrValue("type", "shallow") == "deep" {
		kind = model.HistoryDeep
	}
	st := &model.State{
		Name:       el.SelectAttrValue("id", ""),
		Parent:     parent,
		Kind:       kind,
		Defer:      targets(ext(el, "defer")),
		Invariant:  ext(el, "invariant"),
		Stereotype: ext(el, "stereotype"),
		Note:       note(el),
	}
	if kind == model.Normal {
		var init *etree.Element
		st.Initial, init = initialOf(el)
		if init != nil {
			st.InitialActions = p.actions(st.Name, "<initial>", init, nil)
		}
	}
	st.Variables = p.datamodel(el)
	p.aboutNotes(st, el)
	var invokeID string // of the submachine, whose done.invoke event is the state's completion
	for _, inv := range el.SelectElements("invoke") {
		if p.invoke(st, inv) {
			invokeID = inv.SelectAttrValue("id", "")
		}
	}

	timers := timers(el)
	skip := map[*etree.Element]bool{}
	transitions := make([]*model.Transition, 0)
	for _, t := range el.SelectElements("transition") {
		tr := &model.Transition{
			Source:  st.Name,
			Targets: targets(t.SelectAttrValue("target", "")),
			Event:   t.SelectAttrValue("event", ""),
			Cond:    t.SelectAttrValue("cond", ""),
			Actions: p.actions(st.Name, "<transition>", t, nil),
			Note:    note(t),
		}
		if completes(st, tr.Event, invokeID) {
			tr.Event = ""
		}
		if timer, ok := timers[tr.Event]; ok {
			tr.Event, tr.After = "", timer.delay
			skip[timer.send] = true
			if timer.cancel != nil {
				skip[timer.cancel] = true
			}
		}
		switch {
		case t.SelectAttrValue("type", "external") == "internal":
			tr.Kind = model.Internal
		case ext(t, "kind") == string(model.Local):
			tr.Kind = model.Local
		}
		transitions = append(transitions, tr)
	}
	for _, e := range el.SelectElements("onentry") {
		st.OnEntry = append(st.OnEntry, p.actions(st.Name, "<onentry>", e, skip)...)
	}
	for _, e := range el.SelectElements("onexit") {
		st.OnExit = append(st.OnExit, p.actions(st.Name, "<onexit>", e, skip)...)
	}

	if k := ext(el, "kind"); k != "" {
		st.Kind = model.StateKind(k)
	} else if kind == model.Normal && st.Initial == "" {
		st.Kind = connectorKind(st, transitions)
	}

	for _, c := range el.ChildElements() {
		if _, ok := stateKinds[c.Tag]; ok || known(c, "onentry", "onexit", "transition", "initial", "datamodel", "invoke") {
			continue
		}
		p.unsupported(c, fmt.Sprintf("in state %q", st.Name))
	}

	p.sm.States = append(p.sm.States, st)
	p.sm.Transitions = append(p.sm.Transitions, transitions...)
	p.walk(el, st.Name)
}

// connectorKind recognises the SCXML idiom for UML connector pseudo-states: a
// state with no content of its own that is left as soon as it is entered.
// Several guarded eventless transitions make it a choice; one eventless
// transition with several targets makes it a fork.
func connectorKind(st *model.State, transitions []*model.Transition) model.StateKind {
	if len(st.OnEntry)+len(st.OnExit)+len(st.Do)+len(st.Defer)+len(st.Variables) > 0 || st.Submachine != "" {
		return model.Normal
	}
	guarded := false
	for _, t := range transitions {
		if t.Event != "" || t.After != "" || len(t.Targets) == 0 || !t.IsExternal() {
			return model.Normal
		}
		guarded = guarded || t.Cond != ""
	}
	switch {
	case len(transitions) == 1 && len(transitions[0].Targets) > 1 && !guarded:
		return model.Fork
	case len(transitions) > 1 && guarded:
		return model.Choice
	}
	return model.Normal
}

// completes reports whether event is UML's completion event of st. SCXML
// raises done.state on a compound state when it completes and done.invoke
// when an invoked machine does. A bare done.invoke matches every invoke of
// the state, so it means the submachine's only when nothing else is invoked.
func completes(st *model.State, event, invokeID string) bool {
	switch {
	case event == "done.state."+st.Name:
		return true
	case st.Submachine == "":
		return false
	case invokeID != "" && event == "done.invoke."+invokeID:
		return true
	}
	return len(st.Do) == 0 && (event == "done.invoke" || event == "done.invoke.*")
}

// invoke maps an <invoke> element: invoking another SCXML document is a
// submachine state, anything else is a do activity, named by its src. It
// reports whether the element made st a submachine state.
func (p *parser) invoke(st *model.State, inv *etree.Element) bool {
	typ := inv.SelectAttrValue("type", "")
	src := joinNonEmpty(" ", inv.SelectAttrValue("src", ""), inv.SelectAttrValue("srcexpr", ""))
	isSCXML := typ == "" || typ == "scxml" || strings.HasPrefix(typ, "http://www.w3.org/TR/scxml")
	switch {
	case len(inv.ChildElements()) > 0 || src == "":
		p.fail("state %q: an <invoke> with content or without src is not supported; a do activity is an <invoke> whose src is a name", st.Name)
	case isSCXML && st.Submachine == "":
		st.Submachine = src
		return true
	case isSCXML:
		p.fail("state %q: a state runs one submachine, but it invokes %q and %q", st.Name, st.Submachine, src)
	default:
		if typ != doType {
			p.warn.Addf("state %q: the type %q of the do activity %q is not kept", st.Name, typ, src)
		}
		st.Do = append(st.Do, src)
	}
	return false
}

// doType is the <invoke> type the emitter writes for a do activity.
const doType = extPrefix + ":do"

// timer is a <send delay> in <onentry> and its <cancel> in <onexit>: the SCXML
// idiom for a UML time trigger.
type timer struct {
	delay        string
	send, cancel *etree.Element
}

func timers(el *etree.Element) map[string]*timer {
	byEvent := map[string]*timer{}
	byID := map[string]*timer{}
	for _, onentry := range el.SelectElements("onentry") {
		for _, s := range onentry.SelectElements("send") {
			ev := s.SelectAttrValue("event", "")
			delay := joinNonEmpty(" ", s.SelectAttrValue("delay", ""), s.SelectAttrValue("delayexpr", ""))
			if ev == "" || delay == "" {
				continue
			}
			t := &timer{delay: delay, send: s}
			byEvent[ev] = t
			if id := s.SelectAttrValue("id", ""); id != "" {
				byID[id] = t
			}
		}
	}
	for _, onexit := range el.SelectElements("onexit") {
		for _, c := range onexit.SelectElements("cancel") {
			if t, ok := byID[c.SelectAttrValue("sendid", "")]; ok {
				t.cancel = c
			}
		}
	}
	return byEvent
}

// datamodel reads the <datamodel> children of el as variables.
func (p *parser) datamodel(el *etree.Element) []model.Variable {
	var out []model.Variable
	for _, dm := range el.SelectElements("datamodel") {
		for _, d := range dm.SelectElements("data") {
			v := model.Variable{Name: d.SelectAttrValue("id", "")}
			switch {
			case d.SelectAttrValue("expr", "") != "":
				v.Value = d.SelectAttrValue("expr", "")
			case d.SelectAttrValue("src", "") != "":
				v.Value = "src(" + d.SelectAttrValue("src", "") + ")"
			case len(d.ChildElements()) > 0:
				parts := make([]string, 0, len(d.ChildElements()))
				for _, c := range d.ChildElements() {
					parts = append(parts, rawXML(c))
				}
				v.Value = strings.Join(parts, "")
			default:
				v.Value = strings.TrimSpace(d.Text())
			}
			out = append(out, v)
		}
	}
	return out
}

// targets splits a space-separated attribute (transition targets, deferred
// events). An empty attribute gets a nil list rather than an empty one, so
// that the model still compares equal after a round trip through a format
// that omits empty lists.
func targets(attr string) []string {
	if fields := strings.Fields(attr); len(fields) > 0 {
		return fields
	}
	return nil
}

// initialOf resolves the initial child of a <scxml> or <state> element: the
// initial attribute, else the <initial> element's transition target, else the
// first state-like child in document order (per the SCXML spec). It also
// returns the <initial> element's transition, whose executable content is
// UML's effect of the initial transition.
func initialOf(el *etree.Element) (string, *etree.Element) {
	if v := el.SelectAttrValue("initial", ""); v != "" {
		return v, nil
	}
	if init := el.SelectElement("initial"); init != nil {
		if t := init.SelectElement("transition"); t != nil {
			return t.SelectAttrValue("target", ""), t
		}
	}
	for _, child := range el.ChildElements() {
		if _, ok := stateKinds[child.Tag]; ok && child.Tag != "history" {
			return child.SelectAttrValue("id", ""), nil
		}
	}
	return "", nil
}

// actions reads the executable content of an <onentry>, <onexit> or
// <transition> element of the state named state, one action per <script>,
// leaving out the elements in skip (those already consumed as a time trigger)
// and tpuml annotations. Any other element is an error; model.Validate checks
// that each script holds a name.
func (p *parser) actions(state, in string, el *etree.Element, skip map[*etree.Element]bool) []string {
	var out []string
	for _, c := range el.ChildElements() {
		switch {
		case skip[c] || isExt(c):
		case c.Tag == "script" && c.SelectAttrValue("src", "") == "":
			out = append(out, strings.TrimSpace(c.Text()))
		case c.Tag == "script":
			p.fail("state %q: <script src> in %s is not supported; an action is a <script> holding a name", state, in)
		default:
			p.fail("state %q: <%s> in %s is not supported; an action is a <script> holding a name", state, c.Tag, in)
		}
	}
	return out
}

// ext returns the value of the tpuml extension attribute name on el, or "".
func ext(el *etree.Element, name string) string {
	for i := range el.Attr {
		if a := &el.Attr[i]; a.Key == name && a.Space != "" && a.NamespaceURI() == ExtNamespace {
			return a.Value
		}
	}
	return ""
}

// isExt reports whether el is a tpuml extension element.
func isExt(el *etree.Element) bool {
	return el.Space != "" && el.NamespaceURI() == ExtNamespace
}

// note returns the text of el's own <tpuml:note> child, or "". A note with
// an about attribute describes something the state lists instead; see
// aboutNotes.
func note(el *etree.Element) string {
	for _, c := range notes(el) {
		if c.SelectAttr("about") == nil {
			return noteText(c)
		}
	}
	return ""
}

// aboutNotes reads the notes on what a state lists, which SCXML writes as
// attributes (tpuml:invariant, tpuml:defer) or as several elements that make
// up one UML behaviour (<onentry>, <onexit>, <invoke>), so none has an
// element of its own to hold a note. They sit beside the state's own note,
// saying what they are about.
func (p *parser) aboutNotes(st *model.State, el *etree.Element) {
	for _, c := range notes(el) {
		about := c.SelectAttr("about")
		if about == nil {
			continue
		}
		text := noteText(c)
		switch about.Value {
		case "entry":
			st.EntryNote = text
		case "exit":
			st.ExitNote = text
		case "do":
			st.DoNote = text
		case "invariant":
			st.InvariantNote = text
		case "defer":
			ev := c.SelectAttrValue("event", "")
			if ev == "" {
				p.fail("state %q: a note about a deferred event names no event", st.Name)
				continue
			}
			if st.DeferNotes == nil {
				st.DeferNotes = map[string]string{}
			}
			st.DeferNotes[ev] = text
		default:
			p.fail("state %q: a note about %q is not supported; it is about entry, exit, do, invariant or defer", st.Name, about.Value)
		}
	}
}

// notes returns el's <tpuml:note> children.
func notes(el *etree.Element) []*etree.Element {
	var out []*etree.Element
	for _, c := range el.ChildElements() {
		if c.Tag == "note" && isExt(c) {
			out = append(out, c)
		}
	}
	return out
}

// noteText returns a note's text, one trimmed line per source line.
func noteText(el *etree.Element) string {
	var lines []string
	for _, line := range strings.Split(strings.TrimSpace(el.Text()), "\n") {
		lines = append(lines, strings.TrimSpace(line))
	}
	return strings.Join(lines, "\n")
}

// localTransitions reads SCXML's type="internal" on a transition with a
// target: UML's local transition when the source is compound and every target
// lies inside it, an external one otherwise, as SCXML runs it. Without a
// target it is UML's internal transition.
func localTransitions(sm *model.StateMachine) {
	for _, t := range sm.Transitions {
		if !t.IsInternal() || len(t.Targets) == 0 {
			continue
		}
		t.Kind = model.Local
		for _, tg := range t.Targets {
			if !slices.Contains(sm.Ancestors(tg), t.Source) {
				t.Kind = model.External
			}
		}
	}
}

// known reports whether el is one of the named SCXML elements or a tpuml
// extension element, i.e. something the parser handles elsewhere.
func known(el *etree.Element, tags ...string) bool {
	if isExt(el) {
		return true
	}
	for _, t := range tags {
		if el.Tag == t {
			return true
		}
	}
	return false
}

// rawXML renders an element the model has no field for. The result is
// normalized (indentation removed, canonical escaping) so that the same
// element always yields the same string, whatever the source document looked
// like and however many times it went through the emitter.
func rawXML(el *etree.Element) string {
	doc := etree.NewDocument()
	doc.SetRoot(el.Copy())
	doc.Unindent()
	canonical(doc)
	s, err := doc.WriteToString()
	if err != nil {
		return "<" + el.Tag + ">"
	}
	return strings.TrimSpace(s)
}

func joinNonEmpty(sep string, values ...string) string {
	kept := values[:0:0]
	for _, v := range values {
		if v != "" {
			kept = append(kept, v)
		}
	}
	return strings.Join(kept, sep)
}
