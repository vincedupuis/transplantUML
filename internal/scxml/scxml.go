// Package scxml reads and writes SCXML (State Chart XML) documents: Parser
// turns a document into the model, Emitter (emit.go) writes the model out as
// SCXML.
//
// Natively supported: nested <state>, <parallel>, <final>, <history
// type="shallow|deep">, the initial attribute and the <initial> element,
// <transition> with multiple targets and type="internal", executable content
// in <onentry>, <onexit> and <transition>, <datamodel>/<data> on the machine
// and on states, and <invoke> (a do activity, or a submachine state when it
// invokes another SCXML document). Executable content is stored as text:
// <script> bodies verbatim, common elements (<log>, <assign>, <raise>, <send>,
// <cancel>) rendered to a short readable form, anything else as its XML.
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
// normal to suppress the idiom detection) or, on a transition, local;
// tpuml:defer lists deferred events; tpuml:invariant and tpuml:stereotype
// annotate a state; and a <tpuml:note> child holds a note.
package scxml

import (
	"errors"
	"fmt"
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
		Initial:     initialOf(root),
		States:      make([]*model.State, 0),
		Transitions: make([]*model.Transition, 0),
	}}
	p.sm.Variables = p.datamodel(root)
	p.sm.Note = note(root)
	p.walk(root, "")
	for _, c := range root.ChildElements() {
		if _, ok := stateKinds[c.Tag]; !ok && !known(c, "datamodel") {
			p.warn.Addf("<%s> at the top level is not supported and was dropped", c.Tag)
		}
	}
	return p.sm, p.warn, nil
}

type parser struct {
	sm   *model.StateMachine
	warn model.Warnings
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
		st.Initial = initialOf(el)
	}
	st.Variables = p.datamodel(el)
	for _, inv := range el.SelectElements("invoke") {
		p.invoke(st, inv)
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
			Actions: executableContent(t, nil),
			Note:    note(t),
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
		st.OnEntry = append(st.OnEntry, executableContent(e, skip)...)
	}
	for _, e := range el.SelectElements("onexit") {
		st.OnExit = append(st.OnExit, executableContent(e, skip)...)
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
		p.warn.Addf("state %q: <%s> is not supported and was dropped", st.Name, c.Tag)
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

// invoke maps an <invoke> element: invoking another SCXML document is a
// submachine state, anything else is a do activity.
func (p *parser) invoke(st *model.State, inv *etree.Element) {
	typ := inv.SelectAttrValue("type", "")
	src := joinNonEmpty(" ", inv.SelectAttrValue("src", ""), inv.SelectAttrValue("srcexpr", ""))
	isSCXML := typ == "" || typ == "scxml" || strings.HasPrefix(typ, "http://www.w3.org/TR/scxml")
	if isSCXML {
		typ = "" // the default, not worth repeating
	}
	simple := len(inv.ChildElements()) == 0 && src != ""
	switch {
	case simple && isSCXML && st.Submachine == "":
		st.Submachine = src
	case simple:
		st.Do = append(st.Do, "invoke("+joinNonEmpty(", ", src, typ)+")")
	default:
		st.Do = append(st.Do, rawXML(inv))
	}
}

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
// first state-like child in document order (per the SCXML spec).
func initialOf(el *etree.Element) string {
	if v := el.SelectAttrValue("initial", ""); v != "" {
		return v
	}
	if init := el.SelectElement("initial"); init != nil {
		if t := init.SelectElement("transition"); t != nil {
			return t.SelectAttrValue("target", "")
		}
	}
	for _, child := range el.ChildElements() {
		if _, ok := stateKinds[child.Tag]; ok && child.Tag != "history" {
			return child.SelectAttrValue("id", "")
		}
	}
	return ""
}

// executableContent converts the children of an <onentry>, <onexit> or
// <transition> element to one string per action, leaving out the elements in
// skip (those already consumed as a time trigger) and tpuml annotations.
func executableContent(el *etree.Element, skip map[*etree.Element]bool) []string {
	var out []string
	for _, c := range el.ChildElements() {
		if skip[c] || isExt(c) {
			continue
		}
		attr := func(name string) string { return c.SelectAttrValue(name, "") }
		switch c.Tag {
		case "script":
			if src := attr("src"); src != "" {
				out = append(out, "script("+src+")")
			} else {
				out = append(out, strings.TrimSpace(c.Text()))
			}
		case "log":
			out = append(out, "log("+joinNonEmpty(": ", attr("label"), attr("expr"))+")")
		case "assign":
			value := attr("expr")
			if value == "" {
				value = strings.TrimSpace(c.Text())
			}
			out = append(out, attr("location")+" = "+value)
		case "raise":
			out = append(out, "raise "+attr("event"))
		case "send":
			out = append(out, "send "+joinNonEmpty(" ", attr("event"), attr("eventexpr")))
		case "cancel":
			out = append(out, "cancel "+joinNonEmpty(" ", attr("sendid"), attr("sendidexpr")))
		default:
			out = append(out, rawXML(c))
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

// note returns the text of el's <tpuml:note> child, one trimmed line per
// source line, or "".
func note(el *etree.Element) string {
	for _, c := range el.ChildElements() {
		if c.Tag != "note" || !isExt(c) {
			continue
		}
		var lines []string
		for _, line := range strings.Split(strings.TrimSpace(c.Text()), "\n") {
			lines = append(lines, strings.TrimSpace(line))
		}
		return strings.Join(lines, "\n")
	}
	return ""
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

// rawXML renders an element tpuml has no model field for. The result is
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
