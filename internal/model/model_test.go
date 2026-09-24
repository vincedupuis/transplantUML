package model

import (
	"strings"
	"testing"
)

func sample() *StateMachine {
	return &StateMachine{
		Initial: "a",
		States: []*State{
			{Name: "a", Kind: Normal},
			{Name: "b", Kind: Normal, Initial: "b1"},
			{Name: "b1", Parent: "b", Kind: Normal},
			{Name: "h", Parent: "b", Kind: HistoryDeep},
			{Name: "f", Parent: "b", Kind: Final},
		},
		Transitions: []*Transition{
			{Source: "a", Targets: []string{"b"}, Event: "go"},
			{Source: "h", Targets: []string{"b1"}},
			{Source: "b1", Targets: []string{"f", "a"}, Event: "x"},
		},
	}
}

func TestValidateOK(t *testing.T) {
	if err := sample().Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateErrors(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*StateMachine)
		want   string
	}{
		{"duplicate", func(sm *StateMachine) { sm.States = append(sm.States, &State{Name: "a", Kind: Normal}) }, `duplicate state name "a"`},
		{"unknown parent", func(sm *StateMachine) { sm.States[2].Parent = "nope" }, `unknown parent "nope"`},
		{"final parent", func(sm *StateMachine) { sm.States[2].Parent = "f" }, `cannot have children`},
		{"root history", func(sm *StateMachine) { sm.States[3].Parent = "" }, `history states must be nested`},
		{"bad initial", func(sm *StateMachine) { sm.States[1].Initial = "zzz" }, `unknown initial state "zzz"`},
		{"root initial", func(sm *StateMachine) { sm.Initial = "zzz" }, `unknown initial state "zzz"`},
		{"initial not a child", func(sm *StateMachine) { sm.States[1].Initial = "a" }, `initial state "a" is not one of its children`},
		{"root initial nested", func(sm *StateMachine) { sm.Initial = "b1" }, `initial state "b1" is not a top-level state`},
		{"bad kind", func(sm *StateMachine) { sm.States[0].Kind = "weird" }, `unknown kind "weird"`},
		{"bad source", func(sm *StateMachine) { sm.Transitions[0].Source = "zzz" }, `unknown source "zzz"`},
		{"bad target", func(sm *StateMachine) { sm.Transitions[0].Targets = []string{"zzz"} }, `unknown target "zzz"`},
		{"cycle", func(sm *StateMachine) { sm.States[1].Parent = "b1" }, `forms a cycle`},
		{"choice with actions", func(sm *StateMachine) { sm.States[0].Kind = Choice; sm.States[0].OnEntry = []string{"x"} }, `cannot have entry/exit actions`},
		{"final with do", func(sm *StateMachine) { sm.States[4].Do = []string{"x"} }, `cannot have do activities`},
		{"choice with trigger", func(sm *StateMachine) { sm.States[0].Kind = Choice }, `cannot have a trigger`},
		{"final source", func(sm *StateMachine) { sm.Transitions[0].Source = "f" }, `final state "f" cannot have outgoing transitions`},
		{"terminate source", func(sm *StateMachine) { sm.States[0].Kind = Terminate }, `terminate state "a" cannot have outgoing transitions`},
		{"submachine on parallel", func(sm *StateMachine) { sm.States[0].Kind = Parallel; sm.States[0].Submachine = "m" }, `only normal states can reference a submachine`},
		{"event and after", func(sm *StateMachine) { sm.Transitions[0].After = "5s" }, `has both an event and a time trigger`},
		{"bad transition kind", func(sm *StateMachine) { sm.Transitions[0].Kind = "sideways" }, `unknown kind "sideways"`},
		{"unreached junction", func(sm *StateMachine) { connect(sm, Junction, 0, 1) }, `a junction needs at least one incoming and one outgoing transition, has 0 and 1`},
		{"dead-end choice", func(sm *StateMachine) { connect(sm, Choice, 1, 0) }, `a choice needs at least one incoming and one outgoing transition, has 1 and 0`},
		{"fork to one", func(sm *StateMachine) { connect(sm, Fork, 1, 1) }, `a fork needs exactly one incoming transition and at least two outgoing, has 1 and 1`},
		{"fork from two", func(sm *StateMachine) { connect(sm, Fork, 2, 2) }, `a fork needs exactly one incoming transition and at least two outgoing, has 2 and 2`},
		{"join from one", func(sm *StateMachine) { connect(sm, Join, 1, 1) }, `a join needs at least two incoming transitions and exactly one outgoing, has 1 and 1`},
		{"join to two", func(sm *StateMachine) { connect(sm, Join, 2, 2) }, `a join needs at least two incoming transitions and exactly one outgoing, has 2 and 2`},
		{"two history defaults", func(sm *StateMachine) {
			sm.Transitions = append(sm.Transitions, &Transition{Source: "h", Targets: []string{"f"}})
		}, `a history state has at most one outgoing transition, its default, has 2`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sm := sample()
			c.mutate(sm)
			err := sm.Validate()
			if err == nil || !strings.Contains(err.Error(), c.want) {
				t.Fatalf("want error containing %q, got %v", c.want, err)
			}
		})
	}
}

// connect adds a pseudostate "p" of the given kind to the sample, reached from
// "a" by in transitions and leaving for "a" by out.
func connect(sm *StateMachine, kind StateKind, in, out int) {
	sm.States = append(sm.States, &State{Name: "p", Kind: kind})
	for range in {
		sm.Transitions = append(sm.Transitions, &Transition{Source: "a", Targets: []string{"p"}, Event: "e"})
	}
	for range out {
		sm.Transitions = append(sm.Transitions, &Transition{Source: "p", Targets: []string{"a"}})
	}
}

// A fork may give all its targets to one transition, as SCXML does, and a
// pseudostate may be reached by being its parent's initial child.
func TestValidatePseudostateCounts(t *testing.T) {
	sm := sample()
	connect(sm, Fork, 1, 0)
	sm.Transitions = append(sm.Transitions, &Transition{Source: "p", Targets: []string{"a", "b"}})
	sm.States = append(sm.States, &State{Name: "c", Parent: "b", Kind: Choice})
	sm.States[1].Initial = "c"
	sm.Transitions = append(sm.Transitions, &Transition{Source: "c", Targets: []string{"b1"}})
	if err := sm.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAccessors(t *testing.T) {
	sm := sample()
	if got := names(sm.RootStates()); got != "a,b" {
		t.Errorf("RootStates = %s", got)
	}
	if got := names(sm.Children("b")); got != "b1,h,f" {
		t.Errorf("Children(b) = %s", got)
	}
	if got := names(sm.HistoryOf("b")); got != "h" {
		t.Errorf("HistoryOf(b) = %s", got)
	}
	if sm.InitialOf("") != "a" || sm.InitialOf("b") != "b1" || sm.InitialOf("a") != "" {
		t.Errorf("InitialOf wrong")
	}
	if n := len(sm.OutgoingTransitions("b1")); n != 1 {
		t.Errorf("OutgoingTransitions(b1) = %d", n)
	}
	if n := len(sm.IncomingTransitions("a")); n != 1 {
		t.Errorf("IncomingTransitions(a) = %d", n)
	}
	if sm.State("nope") != nil {
		t.Errorf("State(nope) should be nil")
	}
	if got := strings.Join(sm.Ancestors("b1"), ","); got != "b" {
		t.Errorf("Ancestors(b1) = %s", got)
	}
	if got := sm.Ancestors("a"); len(got) != 0 {
		t.Errorf("Ancestors(a) = %v", got)
	}
}

func TestCommonAncestor(t *testing.T) {
	sm := &StateMachine{States: []*State{
		{Name: "a"}, {Name: "b"}, {Name: "b1", Parent: "b"}, {Name: "b2", Parent: "b"}, {Name: "b11", Parent: "b1"},
	}}
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"a"}, ""}, {[]string{"b1"}, "b"}, {[]string{"b11"}, "b1"},
		{[]string{"b11", "b2"}, "b"}, {[]string{"b11", "b1"}, "b"}, {[]string{"b1", "b11"}, "b"},
		{[]string{"b11", "a"}, ""}, {[]string{"b11", "b11"}, "b1"}, {[]string{"b", "b1"}, ""}, {nil, ""},
	}
	for _, c := range cases {
		if got := sm.CommonAncestor(c.in...); got != c.want {
			t.Errorf("CommonAncestor(%v) = %q, want %q", c.in, got, c.want)
		}
	}
	if got := sm.ScopeOf(&Transition{Source: "b11", Targets: []string{"b2"}}); got != "b" {
		t.Errorf("ScopeOf(b11 -> b2) = %q", got)
	}
}

func TestPredicates(t *testing.T) {
	for _, kind := range []StateKind{Choice, Junction, Fork, Join, EntryPoint, ExitPoint} {
		s := &State{Kind: kind}
		if !s.IsConnector() || !s.IsPseudo() {
			t.Errorf("%s must be a connector pseudo-state", kind)
		}
	}
	for _, kind := range []StateKind{Normal, Parallel} {
		if s := (&State{Kind: kind}); s.IsPseudo() || s.IsConnector() {
			t.Errorf("%s must not be a pseudo-state", kind)
		}
	}
	if s := (&State{Kind: Terminate}); !s.IsPseudo() || s.IsConnector() || !s.IsTerminate() {
		t.Errorf("terminate predicates wrong")
	}
	tr := &Transition{}
	if !tr.IsExternal() || tr.IsLocal() || tr.IsInternal() {
		t.Errorf("an empty kind is external")
	}
	if tr := (&Transition{Kind: External}); !tr.IsExternal() {
		t.Errorf("explicit external")
	}
	if got := (&Transition{Event: "go"}).Trigger(); got != "go" {
		t.Errorf("Trigger = %q", got)
	}
	if got := (&Transition{After: "5s"}).Trigger(); got != "after(5s)" {
		t.Errorf("Trigger = %q", got)
	}
}

func names(states []*State) string {
	out := make([]string, len(states))
	for i, s := range states {
		out[i] = s.Name
	}
	return strings.Join(out, ",")
}
