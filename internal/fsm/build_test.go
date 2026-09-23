package fsm

import (
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/vincedupuis/transplantUML/internal/model"
)

func build(t *testing.T, src string) *model.StateMachine {
	t.Helper()
	sm, warnings, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(warnings) > 0 {
		t.Fatalf("unexpected warnings: %v", warnings)
	}
	if err := sm.Validate(); err != nil {
		t.Fatalf("validate: %v", err)
	}
	return sm
}

// coffee.fsm exercises every goto form, so the model it builds pins down how
// each one resolves.
func TestBuildCoffee(t *testing.T) {
	src, err := os.ReadFile("testdata/coffee.fsm")
	if err != nil {
		t.Fatal(err)
	}
	sm := build(t, string(src))

	if sm.Name != "coffee" || sm.Initial != "idle" {
		t.Errorf("machine = %q initial %q, want coffee/idle", sm.Name, sm.Initial)
	}
	for _, want := range []struct {
		name, parent string
		kind         model.StateKind
	}{
		{"idle", "", model.Normal},
		{"brewing", "", model.Normal},
		{"heating", "brewing", model.Normal},
		{"pouring", "brewing", model.Normal},
		{"brewing.H", "brewing", model.HistoryShallow}, // goto H
		{"final", "", model.Final},                     // goto final, in brewing's own scope
	} {
		s := sm.State(want.name)
		if s == nil {
			t.Errorf("missing state %q", want.name)
			continue
		}
		if s.Parent != want.parent || s.Kind != want.kind {
			t.Errorf("state %q: parent %q kind %q, want %q/%q", s.Name, s.Parent, s.Kind, want.parent, want.kind)
		}
	}
	if got := sm.State("brewing").Initial; got != "heating" {
		t.Errorf("brewing initial = %q, want heating", got)
	}
	if got := sm.State("idle").OnEntry; !slices.Equal(got, []string{"beep"}) {
		t.Errorf("idle entry = %v, want [beep]", got)
	}
	if got := sm.State("brewing").OnExit; !slices.Equal(got, []string{"log", "reset"}) {
		t.Errorf("brewing exit = %v, want [log reset]", got)
	}
	if got := sm.State("brewing").Do; !slices.Equal(got, []string{"heat"}) {
		t.Errorf("brewing do = %v, want [heat]", got)
	}
	if got := sm.State("pouring").Defer; !slices.Equal(got, []string{"coin"}) {
		t.Errorf("pouring defer = %v, want [coin]", got)
	}

	for _, want := range []struct {
		source, event, target, cond string
	}{
		{"idle", "coin", "brewing", ""},                           // bare sibling path
		{"brewing", "cancel", "brewing", ""},                      // goto .
		{"brewing", "error", "final", ""},                         // goto final
		{"brewing", "resume", "brewing.H", ""},                    // goto H
		{"heating", "hot", "pouring", ""},                         // ../pouring
		{"pouring", "done", "idle", "not empty and (cup or mug)"}, // /idle
		{"pouring", "tick", "", ""},                               // targetless
	} {
		i := slices.IndexFunc(sm.Transitions, func(t *model.Transition) bool {
			return t.Source == want.source && t.Event == want.event
		})
		if i < 0 {
			t.Errorf("missing transition %s on %s", want.source, want.event)
			continue
		}
		tr := sm.Transitions[i]
		target := ""
		if len(tr.Targets) > 0 {
			target = tr.Targets[0]
		}
		if target != want.target || tr.Cond != want.cond {
			t.Errorf("%s on %s: target %q cond %q, want %q/%q", want.source, want.event, target, tr.Cond, want.target, want.cond)
		}
	}
}

// A goto names its target outright, wherever it sits and whether or not it is
// declared before the transition that names it.
func TestBuildTargets(t *testing.T) {
	const src = `fsm m {
		initial state a {
			initial state b { initial state c {} on e goto c }
			on f goto b
		}
		state d { on g goto c }
	}`
	sm := build(t, src)
	for _, want := range []struct{ source, target string }{
		{"b", "c"}, // a child
		{"a", "b"}, // a child declared by name alone
		{"d", "c"}, // a grandchild of another branch, declared further up
	} {
		i := slices.IndexFunc(sm.Transitions, func(t *model.Transition) bool { return t.Source == want.source })
		if i < 0 || !slices.Equal(sm.Transitions[i].Targets, []string{want.target}) {
			t.Errorf("from %s: want target %q, got %v", want.source, want.target, sm.Transitions[i].Targets)
		}
	}
}

// A time trigger fires on the delay rather than on an event, so it fills
// After and leaves Event empty; the rest of the clause reads as usual. The
// delay is a time value or, like a guard or an effect, the name of something
// the generated code calls for it.
func TestBuildTimeTrigger(t *testing.T) {
	sm := build(t, "fsm m { initial state a { after(1.5s) [hot] / log goto b } state b { after(250ms) / tick } state c { after(retryDelay) goto a } }")
	for _, want := range []struct {
		source, after, target, cond string
		actions                     []string
	}{
		{"a", "1.5s", "b", "hot", []string{"log"}},
		{"b", "250ms", "", "", []string{"tick"}},
		{"c", "retryDelay", "a", "", nil},
	} {
		i := slices.IndexFunc(sm.Transitions, func(t *model.Transition) bool { return t.Source == want.source })
		if i < 0 {
			t.Errorf("missing transition from %s", want.source)
			continue
		}
		tr := sm.Transitions[i]
		target := ""
		if len(tr.Targets) > 0 {
			target = tr.Targets[0]
		}
		if tr.Event != "" || tr.After != want.after || target != want.target || tr.Cond != want.cond || !slices.Equal(tr.Actions, want.actions) {
			t.Errorf("from %s: event %q after %q target %q cond %q actions %v, want after %q target %q cond %q actions %v",
				want.source, tr.Event, tr.After, target, tr.Cond, tr.Actions, want.after, want.target, want.cond, want.actions)
		}
	}
}

// A behaviour clause and a deferred event may each be written more than once
// in a state, and every line adds to the list the model keeps.
func TestBuildBehaviours(t *testing.T) {
	sm := build(t, "fsm m { initial state a { entry / dim do / poll do / refresh exit / stop on pause / defer on resume / defer } }")
	a := sm.State("a")
	for _, want := range []struct {
		what string
		got  []string
		list []string
	}{
		{"entry", a.OnEntry, []string{"dim"}},
		{"do", a.Do, []string{"poll", "refresh"}},
		{"exit", a.OnExit, []string{"stop"}},
		{"defer", a.Defer, []string{"pause", "resume"}},
	} {
		if !slices.Equal(want.got, want.list) {
			t.Errorf("a %s = %v, want %v", want.what, want.got, want.list)
		}
	}
	// A deferred event is not a transition.
	if got := sm.OutgoingTransitions("a"); len(got) != 0 {
		t.Errorf("transitions = %v, want none", got)
	}
}

func TestBuildErrors(t *testing.T) {
	cases := map[string]string{
		"fsm m { initial state a {} initial state b {} }":         `already starts in "a"`,
		"fsm m { state a {} state a {} }":                         `the machine already has a state called "a"`,
		"fsm m { state p { state a {} } state q { state a {} } }": `the machine already has a state called "a"`,
		"fsm m { state a { on e goto nope } }":                    `the machine has no state called "nope"`,
		"fsm m { state a { on e goto H } }":                       `state "a" has no children, so it has no history`,
		"fsm m { on e goto a state a {} }":                        `the machine itself has no behaviour`,
		"fsm m { after(5s) goto a state a {} }":                   `put "after(5s)" inside a state`,
	}
	for src, want := range cases {
		_, _, err := Parser{}.Parse([]byte(src))
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("%s\n  want error containing %q, got %v", src, want, err)
		}
	}
}

// A scope with children but no initial state builds, since the model allows
// it; the user is told rather than stopped.
func TestBuildNoInitial(t *testing.T) {
	_, warnings, err := Parser{}.Parse([]byte("fsm m { state a { state b {} state c {} } }"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	want := []string{"the machine has no initial state", `state "a" has no initial state`}
	for _, w := range want {
		if !slices.ContainsFunc(warnings, func(got string) bool { return strings.Contains(got, w) }) {
			t.Errorf("missing warning %q, got %v", w, warnings)
		}
	}
}
