// Package model defines the format-neutral state machine representation that
// every parser produces and every emitter or template consumes.
//
// The model is deliberately a superset of what the supported input formats can
// express: pseudo-states (initial, history, final) and parallel regions are
// first-class states, and entry/exit/transition actions are preserved, so that
// converting to another structured format is lossless.
package model

import (
	"errors"
	"fmt"
	"slices"
)

// StateKind classifies a State.
type StateKind string

const (
	Normal         StateKind = "normal"
	Parallel       StateKind = "parallel"
	Final          StateKind = "final"
	HistoryShallow StateKind = "history-shallow"
	HistoryDeep    StateKind = "history-deep"
)

// StateMachine is the root of the model. Hierarchy is expressed through
// State.Parent rather than nesting; the root has the empty name "".
type StateMachine struct {
	Name        string        `json:"name,omitempty"`
	Initial     string        `json:"initial,omitempty"`
	States      []*State      `json:"states"`
	Transitions []*Transition `json:"transitions"`
}

type State struct {
	Name    string    `json:"name"`
	Parent  string    `json:"parent,omitempty"` // "" means top level
	Kind    StateKind `json:"kind"`
	Initial string    `json:"initial,omitempty"` // compound states only
	OnEntry []string  `json:"onEntry,omitempty"`
	OnExit  []string  `json:"onExit,omitempty"`
}

type Transition struct {
	Source   string   `json:"source"`
	Targets  []string `json:"targets,omitempty"` // empty means targetless (stays in Source)
	Event    string   `json:"event,omitempty"`
	Cond     string   `json:"cond,omitempty"`
	Actions  []string `json:"actions,omitempty"`
	Internal bool     `json:"internal,omitempty"`
}

func (s *State) IsNormal() bool      { return s.Kind == Normal }
func (s *State) IsParallel() bool    { return s.Kind == Parallel }
func (s *State) IsFinal() bool       { return s.Kind == Final }
func (s *State) IsDeepHistory() bool { return s.Kind == HistoryDeep }
func (s *State) IsHistory() bool     { return s.Kind == HistoryShallow || s.Kind == HistoryDeep }

// IsPseudo reports whether the state is a history or final pseudo-state, i.e.
// something a diagram draws as a symbol rather than a box.
func (s *State) IsPseudo() bool { return s.IsHistory() || s.IsFinal() }

// State returns the state with the given name, or nil.
func (sm *StateMachine) State(name string) *State {
	for _, s := range sm.States {
		if s.Name == name {
			return s
		}
	}
	return nil
}

// Children returns the direct children of parent, in document order.
// Use "" for the top-level states.
func (sm *StateMachine) Children(parent string) []*State {
	out := make([]*State, 0)
	for _, s := range sm.States {
		if s.Parent == parent {
			out = append(out, s)
		}
	}
	return out
}

// RootStates is shorthand for Children("").
func (sm *StateMachine) RootStates() []*State { return sm.Children("") }

// HistoryOf returns the history pseudo-states declared directly under parent.
func (sm *StateMachine) HistoryOf(parent string) []*State {
	out := make([]*State, 0)
	for _, s := range sm.Children(parent) {
		if s.IsHistory() {
			out = append(out, s)
		}
	}
	return out
}

// InitialOf returns the initial child of the named state, or of the machine
// itself when name is "". It returns "" when there is none.
func (sm *StateMachine) InitialOf(name string) string {
	if name == "" {
		return sm.Initial
	}
	if s := sm.State(name); s != nil {
		return s.Initial
	}
	return ""
}

// OutgoingTransitions returns the transitions whose source is the named state.
func (sm *StateMachine) OutgoingTransitions(source string) []*Transition {
	out := make([]*Transition, 0)
	for _, t := range sm.Transitions {
		if t.Source == source {
			out = append(out, t)
		}
	}
	return out
}

// IncomingTransitions returns the transitions that list the named state as a target.
func (sm *StateMachine) IncomingTransitions(target string) []*Transition {
	out := make([]*Transition, 0)
	for _, t := range sm.Transitions {
		if slices.Contains(t.Targets, target) {
			out = append(out, t)
		}
	}
	return out
}

// Validate checks the structural integrity of the model and returns every
// problem found joined into one error, or nil.
func (sm *StateMachine) Validate() error {
	var errs []error
	fail := func(format string, args ...any) { errs = append(errs, fmt.Errorf(format, args...)) }

	byName := make(map[string]*State, len(sm.States))
	for _, s := range sm.States {
		if s.Name == "" {
			fail("state with empty name (parent %q)", s.Parent)
			continue
		}
		if _, dup := byName[s.Name]; dup {
			fail("duplicate state name %q", s.Name)
			continue
		}
		byName[s.Name] = s
	}

	for _, s := range sm.States {
		if s.Name == "" || byName[s.Name] != s {
			continue // already reported
		}
		switch s.Kind {
		case Normal, Parallel, Final, HistoryShallow, HistoryDeep:
		default:
			fail("state %q: unknown kind %q", s.Name, s.Kind)
		}
		if s.Parent != "" {
			p, ok := byName[s.Parent]
			switch {
			case !ok:
				fail("state %q: unknown parent %q", s.Name, s.Parent)
			case p.IsPseudo():
				fail("state %q: parent %q is a %s state and cannot have children", s.Name, s.Parent, p.Kind)
			}
		} else if s.IsHistory() {
			fail("state %q: history states must be nested in a state", s.Name)
		}
		if s.Initial != "" {
			if !s.IsNormal() {
				fail("state %q: only normal states may declare an initial state", s.Name)
			} else if _, ok := byName[s.Initial]; !ok {
				fail("state %q: unknown initial state %q", s.Name, s.Initial)
			}
		}
		if len(s.OnEntry)+len(s.OnExit) > 0 && s.IsHistory() {
			fail("state %q: history states cannot have entry/exit actions", s.Name)
		}
	}

	// Detect parent cycles: walking up from any state must reach the root.
	for _, s := range sm.States {
		seen := map[string]bool{}
		for cur := s; cur != nil && cur.Parent != ""; cur = byName[cur.Parent] {
			if seen[cur.Name] {
				fail("state %q: parent chain forms a cycle", s.Name)
				break
			}
			seen[cur.Name] = true
		}
	}

	if sm.Initial != "" {
		if _, ok := byName[sm.Initial]; !ok {
			fail("unknown initial state %q", sm.Initial)
		}
	}

	for i, t := range sm.Transitions {
		if t.Source == "" {
			fail("transition #%d: empty source", i)
		} else if _, ok := byName[t.Source]; !ok {
			fail("transition #%d: unknown source %q", i, t.Source)
		}
		for _, tg := range t.Targets {
			if _, ok := byName[tg]; !ok {
				fail("transition #%d (%s): unknown target %q", i, t.Source, tg)
			}
		}
	}

	return errors.Join(errs...)
}
