// Package model defines the format-neutral state machine representation that
// every parser produces and every emitter or template consumes.
//
// The model follows UML state machines rather than any one document format:
// states, regions, the pseudo-states (initial, history, choice, junction, fork,
// join, entry/exit points, terminate), entry/exit/do behaviours, deferred
// events, and transitions with triggers, guards and effects. A format that
// cannot express one of these still parses it (SCXML uses a small extension
// vocabulary) or emits an approximation with a warning, so the user can keep
// one source document and generate whatever outputs they need from it.
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
	Parallel       StateKind = "parallel" // orthogonal state; its children are the regions
	Final          StateKind = "final"
	HistoryShallow StateKind = "history-shallow"
	HistoryDeep    StateKind = "history-deep"
	Choice         StateKind = "choice"      // dynamic branch: guards evaluated on arrival
	Junction       StateKind = "junction"    // static branch/merge point
	Fork           StateKind = "fork"        // one incoming, one outgoing transition per region
	Join           StateKind = "join"        // one incoming per region, one outgoing
	EntryPoint     StateKind = "entry-point" // named entry port on a compound state
	ExitPoint      StateKind = "exit-point"  // named exit port on a compound state
	Terminate      StateKind = "terminate"   // ends the whole machine, no exit actions
)

// TransitionKind is the UML transition kind. The empty string means external.
type TransitionKind string

const (
	External TransitionKind = "external" // exits and re-enters the source, even on a self-transition
	Local    TransitionKind = "local"    // stays inside the source compound state
	Internal TransitionKind = "internal" // no state change, no entry/exit behaviour
)

// StateMachine is the root of the model. Hierarchy is expressed through
// State.Parent rather than nesting; the root has the empty name "".
type StateMachine struct {
	Name        string        `json:"name,omitempty"`
	Initial     string        `json:"initial,omitempty"`
	Variables   []Variable    `json:"variables,omitempty"` // context attributes guards and actions refer to
	Note        string        `json:"note,omitempty"`
	States      []*State      `json:"states"`
	Transitions []*Transition `json:"transitions"`
}

// Variable is a context attribute of the state machine (or of one state).
type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"` // initial value expression
}

type State struct {
	Name       string     `json:"name"`
	Parent     string     `json:"parent,omitempty"` // "" means top level
	Kind       StateKind  `json:"kind"`
	Initial    string     `json:"initial,omitempty"`    // compound states only
	OnEntry    []string   `json:"onEntry,omitempty"`    // entry behaviour
	OnExit     []string   `json:"onExit,omitempty"`     // exit behaviour
	Do         []string   `json:"do,omitempty"`         // do activity, runs while the state is active
	Defer      []string   `json:"defer,omitempty"`      // events queued rather than consumed while here
	Submachine string     `json:"submachine,omitempty"` // referenced state machine (submachine state)
	Invariant  string     `json:"invariant,omitempty"`  // condition that holds while the state is active
	Variables  []Variable `json:"variables,omitempty"`  // attributes scoped to this state
	Stereotype string     `json:"stereotype,omitempty"` // «stereotype» shown on diagrams
	Note       string     `json:"note,omitempty"`
}

type Transition struct {
	Source  string         `json:"source"`
	Targets []string       `json:"targets,omitempty"` // empty means targetless (stays in Source)
	Event   string         `json:"event,omitempty"`   // trigger; empty with no After means a completion transition
	After   string         `json:"after,omitempty"`   // time trigger: fires after this delay in Source (e.g. "5s")
	Cond    string         `json:"cond,omitempty"`    // guard
	Actions []string       `json:"actions,omitempty"` // effect
	Kind    TransitionKind `json:"kind,omitempty"`    // "" means external
	Note    string         `json:"note,omitempty"`
}

func (s *State) IsNormal() bool      { return s.Kind == Normal }
func (s *State) IsParallel() bool    { return s.Kind == Parallel }
func (s *State) IsFinal() bool       { return s.Kind == Final }
func (s *State) IsTerminate() bool   { return s.Kind == Terminate }
func (s *State) IsDeepHistory() bool { return s.Kind == HistoryDeep }
func (s *State) IsHistory() bool     { return s.Kind == HistoryShallow || s.Kind == HistoryDeep }

// IsConnector reports whether the state is a transient routing pseudo-state:
// choice, junction, fork, join, entry point or exit point. The machine never
// rests in one; it is entered and left in the same step.
func (s *State) IsConnector() bool {
	switch s.Kind {
	case Choice, Junction, Fork, Join, EntryPoint, ExitPoint:
		return true
	}
	return false
}

// IsPseudo reports whether the state is a pseudo-state or a final state, i.e.
// anything that cannot have children or behaviours of its own and that a
// diagram draws as a symbol rather than a box.
func (s *State) IsPseudo() bool { return !s.IsNormal() && !s.IsParallel() }

func (t *Transition) IsExternal() bool { return t.Kind == "" || t.Kind == External }
func (t *Transition) IsLocal() bool    { return t.Kind == Local }
func (t *Transition) IsInternal() bool { return t.Kind == Internal }

// Trigger is the transition's trigger as UML writes it: the event name, or
// "after(delay)" for a time trigger, or "" for a completion transition.
func (t *Transition) Trigger() string {
	if t.After != "" {
		return "after(" + t.After + ")"
	}
	return t.Event
}

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

// Ancestors returns the names of the state's parent, grandparent and so on up
// to the top level, nearest first. Unknown or top-level states have none.
func (sm *StateMachine) Ancestors(name string) []string {
	out := make([]string, 0)
	seen := map[string]bool{}
	for s := sm.State(name); s != nil && s.Parent != "" && !seen[s.Parent]; s = sm.State(s.Parent) {
		seen[s.Parent] = true
		out = append(out, s.Parent)
	}
	return out
}

// CommonAncestor returns the innermost state that contains every named state,
// or "" when only the top level does. A state does not contain itself, so
// the common ancestor of a single state is its parent.
func (sm *StateMachine) CommonAncestor(names ...string) string {
	if len(names) == 0 {
		return ""
	}
	common := sm.Ancestors(names[0])
	for _, name := range names[1:] {
		theirs := sm.Ancestors(name)
		for len(common) > 0 && !slices.Contains(theirs, common[0]) {
			common = common[1:]
		}
	}
	if len(common) == 0 {
		return ""
	}
	return common[0]
}

// ScopeOf returns the innermost state that contains the transition's source
// and all its targets, or "" for the top level. See CommonAncestor.
func (sm *StateMachine) ScopeOf(t *Transition) string {
	return sm.CommonAncestor(append([]string{t.Source}, t.Targets...)...)
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
		case Normal, Parallel, Final, HistoryShallow, HistoryDeep,
			Choice, Junction, Fork, Join, EntryPoint, ExitPoint, Terminate:
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
			init, ok := byName[s.Initial]
			switch {
			case !s.IsNormal():
				fail("state %q: only normal states may declare an initial state", s.Name)
			case !ok:
				fail("state %q: unknown initial state %q", s.Name, s.Initial)
			case init.Parent != s.Name:
				fail("state %q: initial state %q is not one of its children", s.Name, s.Initial)
			}
		}
		if s.IsPseudo() && !s.IsFinal() && len(s.OnEntry)+len(s.OnExit) > 0 {
			fail("state %q: %s states cannot have entry/exit actions", s.Name, s.Kind)
		}
		if s.IsPseudo() && len(s.Do)+len(s.Defer)+len(s.Variables) > 0 {
			fail("state %q: %s states cannot have do activities, deferred events or variables", s.Name, s.Kind)
		}
		if s.Submachine != "" && !s.IsNormal() {
			fail("state %q: only normal states can reference a submachine", s.Name)
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
		if init, ok := byName[sm.Initial]; !ok {
			fail("unknown initial state %q", sm.Initial)
		} else if init.Parent != "" {
			fail("initial state %q is not a top-level state", sm.Initial)
		}
	}

	for i, t := range sm.Transitions {
		if t.Source == "" {
			fail("transition #%d: empty source", i)
		} else if src, ok := byName[t.Source]; !ok {
			fail("transition #%d: unknown source %q", i, t.Source)
		} else if src.IsFinal() || src.IsTerminate() {
			fail("transition #%d: %s state %q cannot have outgoing transitions", i, src.Kind, t.Source)
		} else if src.IsConnector() && (t.Event != "" || t.After != "") {
			fail("transition #%d: transitions leaving %s state %q cannot have a trigger", i, src.Kind, t.Source)
		}
		for _, tg := range t.Targets {
			if _, ok := byName[tg]; !ok {
				fail("transition #%d (%s): unknown target %q", i, t.Source, tg)
			}
		}
		if t.Event != "" && t.After != "" {
			fail("transition #%d (%s): has both an event and a time trigger", i, t.Source)
		}
		switch t.Kind {
		case "", External, Local, Internal:
		default:
			fail("transition #%d (%s): unknown kind %q", i, t.Source, t.Kind)
		}
	}

	// UML's constraints on how many transitions a pseudostate joins. A fork
	// counts targets, since a document may give it one multi-target transition.
	// Being a state's initial child counts as the one incoming transition.
	for _, s := range sm.States {
		if byName[s.Name] != s {
			continue
		}
		in := len(sm.IncomingTransitions(s.Name))
		if sm.Initial == s.Name || (s.Parent != "" && byName[s.Parent] != nil && byName[s.Parent].Initial == s.Name) {
			in++
		}
		outgoing := sm.OutgoingTransitions(s.Name)
		out := len(outgoing)
		switch s.Kind {
		case Choice, Junction:
			if in < 1 || out < 1 {
				fail("state %q: a %s needs at least one incoming and one outgoing transition, has %d and %d", s.Name, s.Kind, in, out)
			}
		case Fork:
			out = 0
			for _, t := range outgoing {
				out += len(t.Targets)
			}
			if in != 1 || out < 2 {
				fail("state %q: a fork needs exactly one incoming transition and at least two outgoing, has %d and %d", s.Name, in, out)
			}
		case Join:
			if in < 2 || out != 1 {
				fail("state %q: a join needs at least two incoming transitions and exactly one outgoing, has %d and %d", s.Name, in, out)
			}
		case HistoryShallow, HistoryDeep:
			if out > 1 {
				fail("state %q: a history state has at most one outgoing transition, its default, has %d", s.Name, out)
			}
		}
	}

	return errors.Join(errs...)
}
