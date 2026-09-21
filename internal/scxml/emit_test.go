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
	out, err := Emitter{}.Emit(sm)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// SCXML -> model -> SCXML -> model must be lossless. The two documents are not
// byte-identical (executable content comes back as <script>), but the models are.
func TestEmitRoundTrip(t *testing.T) {
	for _, path := range []string{"../../example/coffee-machine.scxml", "testdata/edge.scxml"} {
		want := parseFile(t, path)
		out := emit(t, want)
		got, err := Parser{}.Parse(out)
		if err != nil {
			t.Fatalf("%s: re-parsing emitted SCXML: %v", path, err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: round trip differs\n%s", path, out)
		}
	}
}

// Regenerate with:
// go run ./cmd/tpuml -i internal/scxml/testdata/edge.scxml -F scxml -o internal/scxml/testdata/edge.emitted.scxml
func TestEmitGolden(t *testing.T) {
	want, err := os.ReadFile("testdata/edge.emitted.scxml")
	if err != nil {
		t.Fatal(err)
	}
	if got := emit(t, parseFile(t, "testdata/edge.scxml")); string(got) != string(want) {
		t.Errorf("output differs from testdata/edge.emitted.scxml\n--- got ---\n%s--- want ---\n%s", got, want)
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
			{Source: "s", Event: "e", Cond: "n > 3", Internal: true, Actions: []string{"raise tick"}},
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

	got, err := Parser{}.Parse([]byte(out))
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
		if _, err := (Emitter{}).Emit(sm); err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("want error containing %q, got %v", want, err)
		}
	}
}
