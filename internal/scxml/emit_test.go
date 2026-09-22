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
	for _, path := range []string{"../../example/coffee-machine.scxml", "testdata/edge.scxml", "testdata/uml.scxml"} {
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
// go run ./cmd/tpuml -i internal/scxml/testdata/<name>.scxml -F scxml -o internal/scxml/testdata/<name>.emitted.scxml
func TestEmitGolden(t *testing.T) {
	for _, name := range []string{"edge", "uml"} {
		want, err := os.ReadFile("testdata/" + name + ".emitted.scxml")
		if err != nil {
			t.Fatal(err)
		}
		if got := emit(t, parseFile(t, "testdata/"+name+".scxml")); string(got) != string(want) {
			t.Errorf("output differs from testdata/%s.emitted.scxml\n--- got ---\n%s--- want ---\n%s", name, got, want)
		}
	}
}

// Every model feature SCXML can only approximate must be named in a warning.
func TestEmitWarnings(t *testing.T) {
	_, warnings, err := Emitter{}.Emit(parseFile(t, "testdata/uml.scxml"))
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		`state "sync": SCXML cannot join regions; the first region to reach it leaves the parallel state`,
		`state "work": SCXML has no deferred events; written as tpuml:defer, which engines ignore`,
		`transition inner -> inner: SCXML has no local transitions; written as an external one tagged tpuml:kind="local"`,
		`state "stop": SCXML has no terminate; written as a <final> state, which runs exit actions`,
	}
	if !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings =\n%s\nwant\n%s", strings.Join(warnings, "\n"), strings.Join(want, "\n"))
	}

	sm := &model.StateMachine{States: []*model.State{{Name: "s", Kind: model.Normal, Do: []string{"spin the wheel"}}}}
	out, warnings, err := Emitter{}.Emit(sm)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `<invoke type="tpuml:do">`) || !strings.Contains(string(out), "<content>spin the wheel</content>") {
		t.Errorf("free-text do activity not written as invoke content:\n%s", out)
	}
	if len(warnings) != 1 || !strings.Contains(warnings[0], "cannot run the do activity") {
		t.Errorf("warnings = %q", warnings)
	}
}

// The extension namespace is declared only when something uses it.
func TestEmitDeclaresExtensionOnlyWhenUsed(t *testing.T) {
	if out := emit(t, parseFile(t, "testdata/edge.scxml")); strings.Contains(string(out), "xmlns:tpuml") {
		t.Errorf("edge.scxml needs no extension, but got:\n%s", out)
	}
	if out := emit(t, parseFile(t, "testdata/uml.scxml")); !strings.Contains(string(out), `xmlns:tpuml="`+ExtNamespace+`"`) {
		t.Errorf("uml.scxml uses the extension, but got:\n%s", out)
	}
}

// execContentSM exercises every executable-content shape the emitter handles:
// a plain script body, an action that is already XML, exit actions, and a
// targetless internal transition. Shared with the schema test (schema_test.go).
func execContentSM() *model.StateMachine {
	return &model.StateMachine{
		Initial: "s",
		States: []*model.State{{
			Name:    "s",
			Kind:    model.Normal,
			OnEntry: []string{"x = 1;", `<if cond="n > 3"><log expr="'big'"/></if>`},
			OnExit:  []string{"log(bye)"},
		}},
		Transitions: []*model.Transition{
			{Source: "s", Event: "e", Cond: "n > 3", Kind: model.Internal, Actions: []string{"raise tick"}},
		},
	}
}

func TestEmitExecutableContent(t *testing.T) {
	sm := execContentSM()
	out := string(emit(t, sm))
	for _, want := range []string{
		"<onentry>", "<script>x = 1;</script>", `<if cond="n > 3">`, `<log expr="'big'"/>`,
		"<onexit>", `<transition event="e" cond="n > 3" type="internal">`,
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
