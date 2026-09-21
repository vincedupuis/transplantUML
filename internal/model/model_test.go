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
		{"bad kind", func(sm *StateMachine) { sm.States[0].Kind = "weird" }, `unknown kind "weird"`},
		{"bad source", func(sm *StateMachine) { sm.Transitions[0].Source = "zzz" }, `unknown source "zzz"`},
		{"bad target", func(sm *StateMachine) { sm.Transitions[0].Targets = []string{"zzz"} }, `unknown target "zzz"`},
		{"cycle", func(sm *StateMachine) { sm.States[1].Parent = "b1" }, `forms a cycle`},
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
}

func names(states []*State) string {
	out := make([]string, len(states))
	for i, s := range states {
		out[i] = s.Name
	}
	return strings.Join(out, ",")
}
