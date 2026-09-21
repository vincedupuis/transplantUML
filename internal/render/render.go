// Package render executes a Go text/template against the model.
//
// Templates get the *model.StateMachine as the root "." and can call these
// functions in addition to the sprig library (https://masterminds.github.io/sprig/):
//
//	State name                  -> *State or nil
//	Children parent             -> []*State   ("" for the top level)
//	RootStates                  -> []*State
//	HistoryOf parent            -> []*State
//	InitialOf name              -> string     ("" for the machine's initial)
//	OutgoingTransitions source  -> []*Transition
//	IncomingTransitions target  -> []*Transition
//	include "name" data         -> string     (like template, but pipeable, e.g. | indent 4)
//	prefix p s                  -> p+s, or "" when s is empty
//	surround p s q              -> p+s+q, or "" when s is empty
//	joinNonEmpty sep s...       -> s joined by sep, skipping empty strings
package render

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/vincedupuis/transplantUML/internal/model"
)

func Render(sm *model.StateMachine, tmplSrc string) (string, error) {
	tmpl := template.New("tpuml")

	funcs := template.FuncMap{
		"include": func(name string, data any) (string, error) {
			var buf bytes.Buffer
			if err := tmpl.ExecuteTemplate(&buf, name, data); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
		"prefix": func(prefix, value string) string {
			if value == "" {
				return ""
			}
			return prefix + value
		},
		"surround": func(prefix, value, suffix string) string {
			if value == "" {
				return ""
			}
			return prefix + value + suffix
		},
		"joinNonEmpty": func(sep string, values ...string) string {
			kept := values[:0:0]
			for _, v := range values {
				if v != "" {
					kept = append(kept, v)
				}
			}
			return strings.Join(kept, sep)
		},
		"State":               sm.State,
		"Children":            sm.Children,
		"RootStates":          sm.RootStates,
		"HistoryOf":           sm.HistoryOf,
		"InitialOf":           sm.InitialOf,
		"OutgoingTransitions": sm.OutgoingTransitions,
		"IncomingTransitions": sm.IncomingTransitions,
	}
	tmpl.Funcs(sprig.FuncMap()).Funcs(funcs)

	if _, err := tmpl.Parse(tmplSrc); err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}
	var out bytes.Buffer
	if err := tmpl.Execute(&out, sm); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}
	return out.String(), nil
}
