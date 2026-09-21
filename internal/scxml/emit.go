package scxml

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/vincedupuis/transplantUML/internal/model"
)

const namespace = "http://www.w3.org/2005/07/scxml"

// Emitter writes the model back out as SCXML. It is the inverse of Parser for
// everything the model holds structurally; executable content, which the
// parser flattens to strings, comes back as <script> bodies unless the string
// is raw XML, so model -> SCXML -> model is stable even though the document is
// not byte-identical to the one it came from.
type Emitter struct{}

func (Emitter) Emit(sm *model.StateMachine) ([]byte, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	root := doc.CreateElement("scxml")
	root.CreateAttr("xmlns", namespace)
	root.CreateAttr("version", "1.0")
	setAttr(root, "name", sm.Name)
	setAttr(root, "initial", sm.Initial)

	e := &emitter{sm: sm, emitted: map[string]bool{}}
	if err := e.children(root, ""); err != nil {
		return nil, err
	}
	if err := e.leftovers(); err != nil {
		return nil, err
	}

	canonical(doc)
	doc.Indent(2)
	var buf bytes.Buffer
	if _, err := doc.WriteTo(&buf); err != nil {
		return nil, fmt.Errorf("writing XML: %w", err)
	}
	return buf.Bytes(), nil
}

var stateTags = map[model.StateKind]string{
	model.Normal:         "state",
	model.Parallel:       "parallel",
	model.Final:          "final",
	model.HistoryShallow: "history",
	model.HistoryDeep:    "history",
}

// emitter walks the flat model as a tree, recording what it wrote so that
// anything the walk cannot reach is reported rather than silently dropped.
type emitter struct {
	sm      *model.StateMachine
	emitted map[string]bool
}

// children writes every direct child of parent (and their descendants) under el.
func (e *emitter) children(el *etree.Element, parent string) error {
	for _, s := range e.sm.Children(parent) {
		tag, ok := stateTags[s.Kind]
		if !ok {
			return fmt.Errorf("state %q: cannot emit unknown kind %q", s.Name, s.Kind)
		}
		child := el.CreateElement(tag)
		e.emitted[s.Name] = true
		setAttr(child, "id", s.Name)
		if s.IsDeepHistory() {
			child.CreateAttr("type", "deep") // shallow is the SCXML default
		}
		setAttr(child, "initial", s.Initial)
		actions(child, "onentry", s.OnEntry)
		actions(child, "onexit", s.OnExit)
		for _, t := range e.sm.OutgoingTransitions(s.Name) {
			tr := child.CreateElement("transition")
			setAttr(tr, "event", t.Event)
			setAttr(tr, "cond", t.Cond)
			setAttr(tr, "target", strings.Join(t.Targets, " "))
			if t.Internal {
				tr.CreateAttr("type", "internal")
			}
			executable(tr, t.Actions)
		}
		if err := e.children(child, s.Name); err != nil {
			return err
		}
	}
	return nil
}

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
