package scxml

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/vincedupuis/transplantUML/internal/model"
)

func emit(t *testing.T, sm *model.StateMachine) []byte {
	t.Helper()
	out, _, err := Emitter{}.Emit(sm)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// SCXML -> model -> SCXML -> model must be lossless. The two documents are not
// byte-identical (executable content comes back as <script>), but the models are.
func TestEmitRoundTrip(t *testing.T) {
	for _, path := range []string{"../../example/coffee-machine.scxml", "testdata/edge.scxml", "testdata/uml-scxml.scxml"} {
		want := parseFile(t, path)
		out := emit(t, want)
		got, _, err := Parser{}.Parse(out)
		if err != nil {
			t.Fatalf("%s: re-parsing emitted SCXML: %v", path, err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: round trip differs\n%s", path, out)
		}
	}
}

// Regenerate with:
// go run ./cmd/fsm -i internal/scxml/testdata/<name>.scxml -F scxml -o internal/scxml/testdata/<name>.emitted.scxml
func TestEmitGolden(t *testing.T) {
	for _, name := range []string{"edge", "uml-scxml"} {
		want, err := os.ReadFile("testdata/" + name + ".emitted.scxml")
		if err != nil {
			t.Fatal(err)
		}
		if got := emit(t, parseFile(t, "testdata/"+name+".scxml")); string(got) != string(want) {
			t.Errorf("output differs from testdata/%s.emitted.scxml\n--- got ---\n%s--- want ---\n%s", name, got, want)
		}
	}
}

// Where SCXML runs the model through a stand-in that behaves the same, a
// warning says so. uml-scxml.scxml is uml.scxml without what SCXML cannot run
// (TestEmitErrors).
func TestEmitWarnings(t *testing.T) {
	_, warnings, err := Emitter{}.Emit(parseFile(t, "testdata/uml-scxml.scxml"))
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		`the machine's initial transition: <scxml> takes no <initial> element; its effect runs in the transient state "initial", which the machine leaves at once`,
		`state "stop": SCXML has no terminate; written as a <final> state at the top level, which ends the machine`,
	}
	if !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings =\n%s\nwant\n%s", strings.Join(warnings, "\n"), strings.Join(want, "\n"))
	}

	sm := &model.StateMachine{States: []*model.State{{Name: "s", Kind: model.Normal, Do: []string{"spin"}}}}
	out, warnings, err := Emitter{}.Emit(sm)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `<invoke type="tpuml:do" src="spin"/>`) {
		t.Errorf("do activity not written as an invoke of its name:\n%s", out)
	}
	if len(warnings) != 0 {
		t.Errorf("warnings = %q", warnings)
	}
}

// <scxml> takes no <initial> element, so the effect of the machine's initial
// transition goes in a transient state, named so as not to clash.
func TestEmitMachineInitialEffect(t *testing.T) {
	sm := &model.StateMachine{
		Name: "m", Initial: "a", InitialActions: []string{"boot", "log"},
		States:      []*model.State{{Name: "a", Kind: model.Normal}, {Name: "initial", Kind: model.Normal}},
		Transitions: []*model.Transition{},
	}
	out, warnings, err := Emitter{}.Emit(sm)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		` initial="initial_2">`,
		"<state id=\"initial_2\" tpuml:kind=\"initial\">\n    <transition target=\"a\">\n      <script>boot</script>\n      <script>log</script>\n",
	} {
		if !strings.Contains(string(out), want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
	if len(warnings) != 1 {
		t.Errorf("warnings = %q", warnings)
	}
	got, _, err := Parser{}.Parse(out)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, sm) {
		t.Errorf("round trip differs:\n%s", out)
	}
}

// A local transition is SCXML's type="internal" with a target, which keeps
// the source active when it is compound and every target lies inside it.
func TestEmitLocal(t *testing.T) {
	sm := &model.StateMachine{
		Initial: "c",
		States: []*model.State{
			{Name: "c", Kind: model.Normal, Initial: "c1", OnEntry: []string{"hello"}},
			{Name: "c1", Parent: "c", Kind: model.Normal},
			{Name: "c2", Parent: "c", Kind: model.Normal},
		},
		Transitions: []*model.Transition{{Source: "c", Targets: []string{"c2"}, Event: "e", Kind: model.Local}},
	}
	out := emit(t, sm)
	if !strings.Contains(string(out), `<transition event="e" target="c2" type="internal"/>`) {
		t.Errorf("local transition not written as type=internal:\n%s", out)
	}
	got, _, err := Parser{}.Parse(out)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, sm) {
		t.Errorf("round trip differs:\n%s", out)
	}
}

// A completion transition waits for the done event of what the state runs: a
// submachine, a region's final state. A simple state completes at once.
func TestEmitCompletion(t *testing.T) {
	sm := &model.StateMachine{
		Initial: "c",
		States: []*model.State{
			{Name: "c", Kind: model.Normal, Initial: "c1"},
			{Name: "c1", Parent: "c", Kind: model.Normal},
			{Name: "s", Kind: model.Normal, Submachine: "help"},
			{Name: "a", Kind: model.Normal},
		},
		Transitions: []*model.Transition{
			{Source: "c", Targets: []string{"s"}},
			{Source: "s", Targets: []string{"a"}},
			{Source: "a", Targets: []string{"c"}},
		},
	}
	out, warnings, err := Emitter{}.Emit(sm)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`<transition event="done.state.c" target="s"/>`,
		`<invoke id="s.submachine" src="help"/>`,
		`<transition event="done.invoke.s.submachine" target="a"/>`,
		`<transition target="c"/>`,
	} {
		if !strings.Contains(string(out), want) {
			t.Errorf("missing %s in:\n%s", want, out)
		}
	}
	if len(warnings) != 0 {
		t.Errorf("warnings = %q", warnings)
	}
}

// The extension namespace is declared only when something uses it.
func TestEmitDeclaresExtensionOnlyWhenUsed(t *testing.T) {
	if out := emit(t, parseFile(t, "testdata/edge.scxml")); strings.Contains(string(out), "xmlns:tpuml") {
		t.Errorf("edge.scxml needs no extension, but got:\n%s", out)
	}
	if out := emit(t, parseFile(t, "testdata/uml-scxml.scxml")); !strings.Contains(string(out), `xmlns:tpuml="`+ExtNamespace+`"`) {
		t.Errorf("uml-scxml.scxml uses the extension, but got:\n%s", out)
	}
}

// execContentSM exercises every executable-content shape the emitter handles:
// entry and exit actions, a guard made of names, and a targetless internal
// transition. Shared with the schema test (schema_test.go).
func execContentSM() *model.StateMachine {
	return &model.StateMachine{
		Initial: "s",
		States: []*model.State{{
			Name:    "s",
			Kind:    model.Normal,
			OnEntry: []string{"start", "count"},
			OnExit:  []string{"stop"},
		}},
		Transitions: []*model.Transition{
			{Source: "s", Event: "e", Cond: "big and not (small or empty)", Kind: model.Internal, Actions: []string{"tick"}},
		},
	}
}

func TestEmitExecutableContent(t *testing.T) {
	sm := execContentSM()
	out := string(emit(t, sm))
	for _, want := range []string{
		"<onentry>\n      <script>start</script>\n      <script>count</script>\n    </onentry>",
		"<onexit>\n      <script>stop</script>\n    </onexit>",
		`<transition event="e" cond="big and not (small or empty)" type="internal">`,
	} {
		if !strings.Contains(out, want) {
			t.Errorf("emitted SCXML is missing %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "target=") {
		t.Errorf("targetless transition must not get a target attribute:\n%s", out)
	}

	got, _, err := Parser{}.Parse([]byte(out))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, sm) {
		t.Errorf("round trip differs:\ngot  %+v\nwant %+v\n%s", *got.States[0], *sm.States[0], out)
	}
}

// An SCXML engine takes the first enabled transition in document order, so a
// fork's transitions are written as one, and an else branch comes last with
// no guard.
func TestEmitForkAndElse(t *testing.T) {
	sm := &model.StateMachine{
		Initial: "c",
		States: []*model.State{
			{Name: "c", Kind: model.Choice},
			{Name: "f", Kind: model.Fork},
			{Name: "p", Kind: model.Parallel},
			{Name: "r1", Parent: "p", Kind: model.Normal},
			{Name: "r2", Parent: "p", Kind: model.Normal},
		},
		Transitions: []*model.Transition{
			{Source: "c", Cond: "else", Targets: []string{"p"}},
			{Source: "c", Cond: "ready", Targets: []string{"f"}},
			{Source: "f", Targets: []string{"r1"}, Actions: []string{"a"}},
			{Source: "f", Targets: []string{"r2"}, Actions: []string{"b"}},
		},
	}
	out := string(emit(t, sm))
	for _, want := range []string{
		`<state id="c" tpuml:kind="choice">
    <transition cond="ready" target="f"/>
    <transition target="p"/>
  </state>`,
		`<state id="f" tpuml:kind="fork">
    <transition target="r1 r2">
      <script>a</script>
      <script>b</script>
    </transition>
  </state>`,
	} {
		if !strings.Contains(out, want) {
			t.Errorf("emitted SCXML is missing\n%s\n%s", want, out)
		}
	}
}

// What SCXML cannot run the way the model means it is an error, and Emit
// fails with all of them.
func TestEmitBehaviourErrors(t *testing.T) {
	cases := map[string]struct {
		sm   *model.StateMachine
		want []string
	}{
		"uml.scxml": {parseFile(t, "testdata/uml.scxml"), []string{
			`state "sync": SCXML cannot join regions; the first region to reach it would leave the parallel state`,
			`state "work": SCXML has no deferred events`,
		}},
		// An invoked SCXML machine starts in its own initial state and
		// reports only when it is done.
		"references": {&model.StateMachine{
			Name: "m", Initial: "a",
			States: []*model.State{
				{Name: "a", Kind: model.Normal},
				{Name: "s", Kind: model.Normal, Submachine: "help"},
				{Name: "in", Parent: "s", Kind: model.EntryPoint},
				{Name: "out", Parent: "s", Kind: model.ExitPoint},
			},
			Transitions: []*model.Transition{
				{Source: "a", Targets: []string{"in"}, Event: "e"},
				{Source: "out", Targets: []string{"a"}},
			},
		}, []string{
			`state "in": SCXML cannot enter an invoked machine through its entry point`,
			`state "out": SCXML cannot leave an invoked machine through its exit point`,
		}},
		// A simple state completes at once, even while its do activity runs; a <final> inside a state or
		// with exit actions ends no machine the way a terminate does; a
		// local transition to its own source leaves it.
		"stand-ins": {&model.StateMachine{
			Name: "m", Initial: "c",
			States: []*model.State{
				{Name: "c", Kind: model.Normal, Initial: "d"},
				{Name: "d", Parent: "c", Kind: model.Normal, Do: []string{"spin"}},
				{Name: "t", Parent: "c", Kind: model.Terminate},
			},
			Transitions: []*model.Transition{
				{Source: "d", Targets: []string{"t"}},
				{Source: "c", Targets: []string{"c"}, Event: "e", Kind: model.Local},
			},
		}, []string{
			`transition c -> c: SCXML keeps a transition inside its source only when the source is compound and every target lies inside it`,
			`state "d": SCXML would take its completion transition without waiting for the do activity to end`,
			`state "t": SCXML has no terminate, and a <final> state that is not at the top level, or that has exit actions, does not end the machine the same way`,
		}},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if err := c.sm.Validate(); err != nil {
				t.Fatal(err)
			}
			out, _, err := Emitter{}.Emit(c.sm)
			if err == nil {
				t.Fatalf("emitted without an error:\n%s", out)
			}
			if got := strings.Split(err.Error(), "\n"); !reflect.DeepEqual(got, c.want) {
				t.Errorf("errors =\n%s\nwant\n%s", err, strings.Join(c.want, "\n"))
			}
		})
	}
}

func TestEmitErrors(t *testing.T) {
	cases := map[string]*model.StateMachine{
		"not reachable": {
			States:      []*model.State{{Name: "a", Parent: "gone", Kind: model.Normal}},
			Transitions: []*model.Transition{{Source: "a"}},
		},
		"unknown kind": {States: []*model.State{{Name: "a", Kind: "sparkly"}}},
	}
	for want, sm := range cases {
		if _, _, err := (Emitter{}).Emit(sm); err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("want error containing %q, got %v", want, err)
		}
	}
}
