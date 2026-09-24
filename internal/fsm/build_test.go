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
		kind                        model.TransitionKind
	}{
		{"idle", "coin", "brewing", "", ""},                           // a sibling
		{"brewing", "cancel", "brewing", "", ""},                      // itself, exited and re-entered
		{"brewing", "restart", "heating", "", model.Local},            // goto local
		{"brewing", "error", "final", "", ""},                         // goto final
		{"brewing", "resume", "brewing.H", "", ""},                    // goto H
		{"heating", "hot", "pouring", "", ""},                         // a sibling
		{"pouring", "done", "idle", "not empty and (cup or mug)", ""}, // out of the parent
		{"pouring", "tick", "", "", model.Internal},                   // no goto
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
		if target != want.target || tr.Cond != want.cond || tr.Kind != want.kind {
			t.Errorf("%s on %s: target %q cond %q kind %q, want %q/%q/%q", want.source, want.event, target, tr.Cond, tr.Kind, want.target, want.cond, want.kind)
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

// A goto is external, as UML's default is, unless it says local, which reaches
// any state inside the source, its history included. A clause without a goto
// is internal.
func TestBuildTransitionKinds(t *testing.T) {
	const src = `fsm m {
		initial state a {
			initial state b { initial state c {} }
			on ext goto c
			on loc goto local c
			on hist goto local H*
			on self goto a
			on stay / log
		}
	}`
	sm := build(t, src)
	for _, want := range []struct {
		event, target string
		kind          model.TransitionKind
	}{
		{"ext", "c", ""},
		{"loc", "c", model.Local},
		{"hist", "a.H-deep", model.Local},
		{"self", "a", ""},
		{"stay", "", model.Internal},
	} {
		i := slices.IndexFunc(sm.Transitions, func(t *model.Transition) bool { return t.Event == want.event })
		if i < 0 {
			t.Errorf("missing transition on %s", want.event)
			continue
		}
		tr := sm.Transitions[i]
		if strings.Join(tr.Targets, " ") != want.target || tr.Kind != want.kind {
			t.Errorf("on %s: targets %v kind %q, want %q/%q", want.event, tr.Targets, tr.Kind, want.target, want.kind)
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

// An invariant reaches the model as the document writes it, spacing
// included, on a state, a parallel state and a submachine state alike.
func TestBuildInvariant(t *testing.T) {
	sm := build(t, `fsm m {
		initial state a { invariant [cashLoaded  and not (jammed or empty)] on e goto p }
		parallel state p { invariant [powered] region r { initial submachine s { invariant [ready] } } }
	}`)
	for name, want := range map[string]string{
		"a": "cashLoaded  and not (jammed or empty)",
		"p": "powered",
		"s": "ready",
	} {
		if got := sm.State(name).Invariant; got != want {
			t.Errorf("%s invariant = %q, want %q", name, got, want)
		}
	}
	// An invariant is not a transition.
	if got := sm.OutgoingTransitions("p"); len(got) != 0 {
		t.Errorf("p transitions = %v, want none", got)
	}
}

// shop.fsm declares every state kind, so the model it builds pins down what
// each declaration and each goto becomes.
func TestBuildStateKinds(t *testing.T) {
	src, err := os.ReadFile("../../example/shop.fsm")
	if err != nil {
		t.Fatal(err)
	}
	sm := build(t, string(src))

	for _, want := range []struct {
		name, parent string
		kind         model.StateKind
	}{
		{"reorder", "", model.EntryPoint}, // a point of the machine itself
		{"express", "checkout", model.EntryPoint},
		{"cancelled", "checkout", model.ExitPoint},
		{"route", "checkout", model.Choice},
		{"paid", "checkout", model.Junction},
		{"split", "", model.Fork},
		{"support", "", model.Normal}, // a submachine state
		{"shipping", "", model.Parallel},
		{"warehouse", "shipping", model.Normal}, // a region
		{"packing", "warehouse", model.Normal},
		{"merge", "", model.Join},
		{"terminate", "", model.Terminate},                 // goto terminate, beside browsing
		{"checkout.H", "checkout", model.HistoryShallow},   // goto checkout.H
		{"checkout.H-deep", "checkout", model.HistoryDeep}, // goto checkout.H*
		{"final", "", model.Final},
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
	if got := sm.State("checkout").Stereotype; got != "secure" {
		t.Errorf("checkout stereotype = %q, want secure", got)
	}
	// A submachine state is named after the machine it refers to.
	if s := sm.State("support"); s.Submachine != "support" || !slices.Equal(s.OnEntry, []string{"openChat"}) {
		t.Errorf("support: submachine %q entry %v, want support/[openChat]", s.Submachine, s.OnEntry)
	}
	// Every region is entered, so the parallel state starts in none of them;
	// each region starts in its own initial child.
	if got := sm.State("shipping").Initial; got != "" {
		t.Errorf("shipping initial = %q, want none", got)
	}
	if got := sm.State("warehouse").Initial; got != "packing" {
		t.Errorf("warehouse initial = %q, want packing", got)
	}

	for _, want := range []struct {
		source, cond, target string
		actions              []string
	}{
		{"reorder", "", "checkout", []string{"loadBasket"}},
		{"express", "", "paying", []string{"useSavedCard"}},
		{"cancelled", "", "browsing", nil},
		{"route", "large", "review", nil},
		{"route", "else", "paying", nil}, // UML's catch-all guard
		{"paid", "", "split", []string{"receipt"}},
		{"split", "", "packing", nil}, // one transition per fork line
		{"split", "", "invoicing", []string{"notify"}},
		{"support", "", "browsing", nil}, // leaves when the submachine completes
		{"packed", "", "merge", nil},     // a completion transition
		{"sent", "paidInFull", "merge", nil},
		{"merge", "", "done", []string{"close"}},
	} {
		i := slices.IndexFunc(sm.Transitions, func(t *model.Transition) bool {
			return t.Source == want.source && slices.Equal(t.Targets, []string{want.target})
		})
		if i < 0 {
			t.Errorf("missing transition %s -> %s", want.source, want.target)
			continue
		}
		tr := sm.Transitions[i]
		if tr.Trigger() != "" || tr.Cond != want.cond || !slices.Equal(tr.Actions, want.actions) {
			t.Errorf("%s -> %s: trigger %q cond %q actions %v, want none/%q/%v",
				want.source, want.target, tr.Trigger(), tr.Cond, tr.Actions, want.cond, want.actions)
		}
	}
}

// goto final and goto terminate end the region the transition's source sits
// in, and an exit point leads out of its state, so for one of those they end
// the region around that state.
func TestBuildEndings(t *testing.T) {
	sm := build(t, `fsm m {
		initial state a {
			exit point out goto final
			initial state b { on e goto terminate on f goto final }
		}
	}`)
	for _, want := range []struct {
		source, target, parent string
	}{
		{"out", "final", ""},
		{"b", "a.terminate", "a"},
		{"b", "a.final", "a"},
	} {
		i := slices.IndexFunc(sm.Transitions, func(t *model.Transition) bool {
			return t.Source == want.source && slices.Equal(t.Targets, []string{want.target})
		})
		if i < 0 {
			t.Errorf("missing transition %s -> %s", want.source, want.target)
			continue
		}
		if got := sm.State(want.target).Parent; got != want.parent {
			t.Errorf("%s: parent %q, want %q", want.target, got, want.parent)
		}
	}
}

// Line breaks carry no meaning, so a deferred event followed by a goto is two
// clauses: the deferral and a completion transition.
func TestBuildTwoClausesOnOneLine(t *testing.T) {
	sm := build(t, "fsm m { initial state s { on e / defer goto t } state t {} }")
	if got := sm.State("s").Defer; !slices.Equal(got, []string{"e"}) {
		t.Errorf("s defer = %v, want [e]", got)
	}
	if got := sm.OutgoingTransitions("s"); len(got) != 1 || got[0].Trigger() != "" || !slices.Equal(got[0].Targets, []string{"t"}) {
		t.Errorf("transitions = %v, want one completion transition to t", got)
	}
}

// Every declaration may carry a stereotype, which reaches the model without
// its brackets.
func TestBuildStereotype(t *testing.T) {
	sm := build(t, `fsm m {
		initial state a <<s1>> { on e goto f }
		fork f <<s8>> { goto sub goto x }
		parallel state p <<s2>> {
			entry point in <<s3>> goto x
			region r <<s4>> { initial submachine sub <<s5>> { goto c } choice c <<s6>> goto k }
			region r2 { initial state x { goto k } }
		}
		join k <<s9>> goto j
		junction j <<s7>> goto a
	}`)
	for name, want := range map[string]string{
		"a": "s1", "p": "s2", "in": "s3", "r": "s4", "sub": "s5",
		"c": "s6", "j": "s7", "f": "s8", "k": "s9",
	} {
		if got := sm.State(name).Stereotype; got != want {
			t.Errorf("%s stereotype = %q, want %q", name, got, want)
		}
	}
}

// A submachine state may start its scope, which then starts in the machine it
// refers to.
func TestBuildInitialSubmachine(t *testing.T) {
	sm := build(t, "fsm m { initial state s { initial submachine pay { goto s } } }")
	if got := sm.State("s").Initial; got != "pay" {
		t.Errorf("s initial = %q, want pay", got)
	}
	if got := sm.State("pay"); got.Parent != "s" || got.Submachine != "pay" {
		t.Errorf("pay: parent %q submachine %q, want s/pay", got.Parent, got.Submachine)
	}
}

// A choice or junction with a single branch may write it on the declaring
// line. It holds that one branch only, so the next clause belongs to the
// enclosing state.
func TestBuildOneLineBranch(t *testing.T) {
	sm := build(t, "fsm m { initial state s { on e goto c choice c [g] / a goto t [h] goto t } state t {} }")
	if got := sm.OutgoingTransitions("c"); len(got) != 1 || got[0].Cond != "g" || !slices.Equal(got[0].Actions, []string{"a"}) {
		t.Errorf("c transitions = %v, want one guarded by g doing a", got)
	}
	if got := sm.OutgoingTransitions("s"); len(got) != 2 || got[1].Cond != "h" {
		t.Errorf("s transitions = %v, want e and a completion transition guarded by h", got)
	}
}

func TestBuildErrors(t *testing.T) {
	cases := map[string]string{
		"fsm m { initial state a {} initial state b {} }":                                   `already starts in "a"`,
		"fsm m { state a {} state a {} }":                                                   `the machine already has a state called "a"`,
		"fsm m { state p { state a {} } state q { state a {} } }":                           `the machine already has a state called "a"`,
		"fsm m { state a { on e goto nope } }":                                              `the machine has no state called "nope"`,
		"fsm m { state a { on e goto H } }":                                                 `state "a" has no children, so it has no history`,
		"fsm m { on e goto a state a {} }":                                                  `the machine itself has no behaviour`,
		"fsm m { after(5s) goto a state a {} }":                                             `put "after(5s)" inside a state`,
		"fsm m { goto a state a {} }":                                                       `put "goto a" inside a state`,
		"fsm m { state a {} choice a {} }":                                                  `the machine already has a state called "a"`,
		"fsm m { parallel state p { region a {} region b {} } fork a { goto a } }":          `the machine already has a state called "a"`,
		"fsm m { parallel state p { region r { initial state a {} initial state b {} } } }": `state "r" already starts in "a"`,
		"fsm m { state a { on e goto nope.H } }":                                            `the machine has no state called "nope"`,
		"fsm m { state a { on e goto b.H* } state b {} }":                                   `state "b" has no children, so it has no history`,
		"fsm m { choice c { goto H } }":                                                     `state "c" has no children, so it has no history`,
		"fsm m { parallel state p { region r { state s {} } } state a { on e goto p.H } }":  `parallel state "p" has no history of its own`,
		"fsm m { parallel state p { region r { state s {} } on e goto H* } }":               `parallel state "p" has no history of its own`,
		"fsm m { parallel state p { entry point e goto final region r {} } }":               `parallel state "p" holds only regions`,
		"fsm m { state a { on e goto local a } }":                                           `a local transition stays inside "a", but "a" is not inside it`,
		"fsm m { state a { on e goto local b } state b {} }":                                `a local transition stays inside "a", but "b" is not inside it`,
		"fsm m { state a { state c {} } state b { on e goto local a.H } }":                  `a local transition stays inside "b", but "a.H" is not inside it`,
		"fsm m { state a { choice c { goto local b } state b {} } }":                        `a local transition stays inside "c", but "b" is not inside it`,
		"fsm m { submachine a {} state a {} }":                                              `the machine already has a state called "a"`,
		"fsm m { initial submachine a {} initial state b {} }":                              `already starts in "a"`,
		"fsm m { state a { invariant [x] invariant [y] } }":                                 `state "a" already has the invariant [x]`,
		"fsm m { invariant [x] state a {} }":                                                `put "invariant [x]" inside a state`,
		"fsm m { submachine a { on e goto H } }":                                            `state "a" has no children, so it has no history`,
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

// A region starts in its initial child like any state, while a parallel state
// enters all its regions, and a scope holding only pseudostates has no state
// to start in.
func TestBuildNoInitialRegion(t *testing.T) {
	_, warnings, err := Parser{}.Parse([]byte("fsm m { initial parallel state p { region r { state a {} } } choice c { goto p } }"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if want := []string{`state "r" has no initial state`}; !slices.Equal(warnings, want) {
		t.Errorf("warnings = %v, want %v", warnings, want)
	}
}
