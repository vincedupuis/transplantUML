// Package scxml reads and writes SCXML (State Chart XML) documents: Parser
// turns a document into the model, Emitter (emit.go) writes the model back
// out as the same subset of SCXML.
//
// Supported: nested <state>, <parallel>, <final>, <history type="shallow|deep">,
// the initial attribute and the <initial> element, <transition> with multiple
// targets and type="internal", and executable content in <onentry>, <onexit>
// and <transition>. Executable content is stored as text: <script> bodies
// verbatim, common elements (<log>, <assign>, <raise>, <send>, <cancel>)
// rendered to a short readable form, anything else as its XML source.
// <datamodel>, <invoke> and <donedata> are ignored.
package scxml

import (
	"errors"
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/vincedupuis/transplantUML/internal/model"
)

type Parser struct{}

func (Parser) Parse(src []byte) (*model.StateMachine, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(src); err != nil {
		return nil, fmt.Errorf("parsing XML: %w", err)
	}
	root := doc.SelectElement("scxml")
	if root == nil {
		return nil, errors.New("no <scxml> root element found")
	}

	sm := &model.StateMachine{
		Name:        root.SelectAttrValue("name", ""),
		Initial:     initialOf(root),
		States:      make([]*model.State, 0),
		Transitions: make([]*model.Transition, 0),
	}
	walk(sm, root, "")
	return sm, nil
}

var stateKinds = map[string]model.StateKind{
	"state":    model.Normal,
	"parallel": model.Parallel,
	"final":    model.Final,
	"history":  model.HistoryShallow,
}

// walk appends every state-like child of el (and their descendants) to sm,
// together with the transitions declared directly under them.
func walk(sm *model.StateMachine, el *etree.Element, parent string) {
	for _, child := range el.ChildElements() {
		kind, ok := stateKinds[child.Tag]
		if !ok {
			continue
		}
		if kind == model.HistoryShallow && child.SelectAttrValue("type", "shallow") == "deep" {
			kind = model.HistoryDeep
		}
		st := &model.State{
			Name:   child.SelectAttrValue("id", ""),
			Parent: parent,
			Kind:   kind,
		}
		if kind == model.Normal {
			st.Initial = initialOf(child)
		}
		for _, e := range child.SelectElements("onentry") {
			st.OnEntry = append(st.OnEntry, executableContent(e)...)
		}
		for _, e := range child.SelectElements("onexit") {
			st.OnExit = append(st.OnExit, executableContent(e)...)
		}
		sm.States = append(sm.States, st)

		for _, t := range child.SelectElements("transition") {
			sm.Transitions = append(sm.Transitions, &model.Transition{
				Source:   st.Name,
				Targets:  targets(t.SelectAttrValue("target", "")),
				Event:    t.SelectAttrValue("event", ""),
				Cond:     t.SelectAttrValue("cond", ""),
				Actions:  executableContent(t),
				Internal: t.SelectAttrValue("type", "external") == "internal",
			})
		}
		walk(sm, child, st.Name)
	}
}

// targets splits a transition's target attribute. A targetless transition gets
// a nil list rather than an empty one, so that the model still compares equal
// after a round trip through a format that omits empty lists.
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
// <transition> element to one string per action.
func executableContent(el *etree.Element) []string {
	var out []string
	for _, c := range el.ChildElements() {
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
