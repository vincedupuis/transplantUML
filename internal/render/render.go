// Package render executes a Go text/template against the model.
//
// Templates get the *model.StateMachine as the root "." and can call these
// functions in addition to the sprig library (https://masterminds.github.io/sprig/):
//
//	States                      -> []*State   (all of them, in model order)
//	Transitions                 -> []*Transition
//	State name                  -> *State or nil
//	Children parent             -> []*State   ("" for the top level)
//	RootStates                  -> []*State
//	HistoryOf parent            -> []*State
//	IsReference name            -> bool       (an entry or exit point of a submachine state)
//	InitialOf name              -> string     ("" for the machine's initial)
//	InitialActionsOf name       -> []string   (effect of that initial transition)
//	Ancestors name              -> []string   (parent, grandparent, ... nearest first)
//	CommonAncestor name...      -> string     (innermost state containing them all, "" for the top level)
//	ScopeOf transition          -> string     (innermost state containing its source and targets)
//	OutgoingTransitions source  -> []*Transition
//	IncomingTransitions target  -> []*Transition
//	include "name" data         -> string     (like template, but pipeable, e.g. | indent 4)
//	prefix p s                  -> p+s, or "" when s is empty
//	surround p s q              -> p+s+q, or "" when s is empty
//	joinNonEmpty sep s...       -> s joined by sep, skipping empty strings
//	warn format args...         -> "" (records a warning for the user, printf-style)
//	file name                   -> a mark: the text after it, up to the next mark, is the file name
//
// A template calls warn when it cannot draw something the model holds, so the
// user learns what the output leaves out; Render returns the warnings. A
// template that writes several files starts each with file; Files splits the
// output at those marks.
package render

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/vincedupuis/transplantUML/internal/model"
)

func Render(sm *model.StateMachine, tmplSrc string) (string, model.Warnings, error) {
	tmpl := template.New("fsm")
	var warnings model.Warnings

	funcs := template.FuncMap{
		"warn": func(format string, args ...any) string {
			warnings.Addf(format, args...)
			return ""
		},
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
		"file": func(name string) (string, error) {
			if name == "" || name != filepath.Base(name) || name == "." || name == ".." {
				return "", fmt.Errorf("file %q: a file name has no folder", name)
			}
			return fileMark + name + fileMark, nil
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
		"States":              func() []*model.State { return sm.States },
		"Transitions":         func() []*model.Transition { return sm.Transitions },
		"State":               sm.State,
		"Children":            sm.Children,
		"RootStates":          sm.RootStates,
		"HistoryOf":           sm.HistoryOf,
		"IsReference":         sm.IsReference,
		"InitialOf":           sm.InitialOf,
		"InitialActionsOf":    sm.InitialActionsOf,
		"Ancestors":           sm.Ancestors,
		"CommonAncestor":      sm.CommonAncestor,
		"ScopeOf":             sm.ScopeOf,
		"OutgoingTransitions": sm.OutgoingTransitions,
		"IncomingTransitions": sm.IncomingTransitions,
	}
	tmpl.Funcs(sprig.FuncMap()).Funcs(funcs)

	if _, err := tmpl.Parse(tmplSrc); err != nil {
		return "", nil, fmt.Errorf("parsing template: %w", err)
	}
	var out bytes.Buffer
	if err := tmpl.Execute(&out, sm); err != nil {
		return "", nil, fmt.Errorf("executing template: %w", err)
	}
	return out.String(), warnings, nil
}

// fileMark surrounds a file name in the output of a template that writes
// several files. A NUL byte appears in no text a template means to write.
const fileMark = "\x00fsm-file\x00"

// File is one of the files a template writes.
type File struct {
	Name    string
	Content string
}

// Files splits the output of Render at the marks the template's file calls
// left, or returns nil when there are none. Only white space may come before
// the first mark, and a name may not be used twice.
func Files(out string) ([]File, error) {
	parts := strings.Split(out, fileMark)
	if len(parts) == 1 {
		return nil, nil
	}
	if strings.TrimSpace(parts[0]) != "" {
		return nil, fmt.Errorf("the template writes text before its first file")
	}
	var files []File
	seen := map[string]bool{}
	for i := 1; i+1 < len(parts); i += 2 {
		name := parts[i]
		if seen[name] {
			return nil, fmt.Errorf("the template writes the file %q twice", name)
		}
		seen[name] = true
		files = append(files, File{Name: name, Content: strings.TrimLeft(parts[i+1], "\n")})
	}
	return files, nil
}
