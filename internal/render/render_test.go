package render

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/vincedupuis/transplantUML/assets"
	"github.com/vincedupuis/transplantUML/internal/format"
	"github.com/vincedupuis/transplantUML/internal/model"
	"github.com/vincedupuis/transplantUML/internal/scxml"
)

// renderFile parses path with the parser its extension names and renders it
// with the built-in template.
func renderFile(t *testing.T, path string) string {
	t.Helper()
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	parser, err := format.ParserFor(format.Detect(path))
	if err != nil {
		t.Fatal(err)
	}
	sm, _, err := parser.Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	out, _, err := Render(sm, assets.PlantUML)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// goldens maps every input document that has a checked-in PlantUML rendering
// to that rendering: the test fixtures and the examples shipped in example/,
// whose .puml files are what users look at first.
func goldens(t *testing.T) map[string]string {
	t.Helper()
	out := map[string]string{
		"../scxml/testdata/edge.scxml": "testdata/edge.puml",
		"../scxml/testdata/uml.scxml":  "testdata/uml.puml",
	}
	for _, pattern := range []string{"../../example/*.scxml", "../../example/*.json", "../../example/*.fsm"} {
		paths, err := filepath.Glob(pattern)
		if err != nil {
			t.Fatal(err)
		}
		for _, p := range paths {
			out[p] = strings.TrimSuffix(p, filepath.Ext(p)) + ".puml"
		}
	}
	if len(out) < 4 {
		t.Fatalf("expected the examples to be found, got %v", out)
	}
	return out
}

// Golden-file tests for the built-in PlantUML template. Regenerate a golden
// file with: go run ./cmd/tpuml -i <input> -o <golden>.puml
func TestPlantUMLGolden(t *testing.T) {
	for input, golden := range goldens(t) {
		t.Run(filepath.Base(input), func(t *testing.T) {
			want, err := os.ReadFile(golden)
			if err != nil {
				t.Fatal(err)
			}
			if got := renderFile(t, input); got != string(want) {
				t.Errorf("output differs from %s\n--- got ---\n%s--- want ---\n%s", golden, got, want)
			}
		})
	}
}

// The built-in template must say what it cannot draw.
func TestPlantUMLWarnings(t *testing.T) {
	src, err := os.ReadFile("../scxml/testdata/uml.scxml")
	if err != nil {
		t.Fatal(err)
	}
	sm, _, err := scxml.Parser{}.Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	_, warnings, err := Render(sm, assets.PlantUML)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		`transition inner -> inner: PlantUML has no local transitions; drawn as an external one`,
		`transition inner (note): PlantUML cannot attach a note to an internal transition`,
		`state "stop": PlantUML has no terminate symbol; drawn as a final state`,
		`transition split -> right: PlantUML cannot draw arrows across the boundary of a parallel region other than the first`,
		`transition right -> sync: PlantUML cannot draw arrows across the boundary of a parallel region other than the first`,
	}
	if !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings =\n%s\nwant\n%s", strings.Join(warnings, "\n"), strings.Join(want, "\n"))
	}
}

// A compound region's own transitions have no PlantUML equivalent.
func TestPlantUMLRegionTransitions(t *testing.T) {
	sm := &model.StateMachine{
		States: []*model.State{
			{Name: "p", Kind: model.Parallel}, {Name: "r", Parent: "p", Kind: model.Normal},
			{Name: "r1", Parent: "r", Kind: model.Normal}, {Name: "x", Kind: model.Normal},
		},
		Transitions: []*model.Transition{{Source: "r", Targets: []string{"x"}, Event: "e"}},
	}
	out, warnings, err := Render(sm, assets.PlantUML)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(out, "r -->") {
		t.Errorf("region transition must not be drawn:\n%s", out)
	}
	if len(warnings) != 1 || !strings.Contains(warnings[0], `region "r"`) {
		t.Errorf("warnings = %q", warnings)
	}
}

// What PlantUML cannot say on its own: names it would misread, behaviours and
// notes on states it draws as symbols, a region's own details, a second
// stereotype, line breaks in text.
func TestPlantUMLApproximations(t *testing.T) {
	sm := &model.StateMachine{
		Name: "m", Initial: "a.b",
		States: []*model.State{
			{Name: "a.b", Kind: model.Normal, OnEntry: []string{"one\ntwo"}, Stereotype: "st"},
			{Name: "my-comp", Kind: model.Normal, Initial: "in"},
			{Name: "h", Parent: "my-comp", Kind: model.HistoryShallow, Note: "remembers"},
			{Name: "in", Parent: "my-comp", Kind: model.Normal},
			{Name: "c", Kind: model.Choice, Stereotype: "mine"},
			{Name: "done", Kind: model.Final, OnEntry: []string{"bye()"}, Note: "over"},
			{Name: "p", Kind: model.Parallel},
			{Name: "r", Parent: "p", Kind: model.Normal, Note: "region"},
			{Name: "r1", Parent: "r", Kind: model.Normal},
			{Name: "q", Parent: "p", Kind: model.Parallel},
			{Name: "q1", Parent: "q", Kind: model.Normal},
			{Name: "q2", Parent: "q", Kind: model.Normal},
		},
		Transitions: []*model.Transition{
			{Source: "a.b", Targets: []string{"my-comp"}, Event: "go"},
			{Source: "h", Targets: []string{"in"}, Note: "default"},
			{Source: "a.b", Targets: []string{"r"}, Event: "into"},
			{Source: "in", Targets: []string{"c"}, Event: "end"},
			{Source: "c", Targets: []string{"done"}},
		},
	}
	if err := sm.Validate(); err != nil {
		t.Fatal(err)
	}
	out, warnings, err := Render(sm, assets.PlantUML)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"\ntitle m\n",
		"\n[*] --> a_b\n",
		"\nstate \"«st»\\na.b\" as a_b <<st>>\n",
		"\na_b : entry / one\\ntwo\n",
		"\nstate \"my-comp\" as my_comp {\n",
		"\n    state h <<history>>\n    note right of h : remembers\n",
		"\n    h --> in\n    note on link : default\n",
		"\nstate c <<choice>>\n",
		"\nstate done <<end>>\nnote right of done\n    entry / bye()\n    over\nend note\n",
		"\nstate p {\n    state r1\n    --\n    state q {\n        state q1\n        --\n        state q2\n    }\n}\n",
		"\na_b --> my_comp: go\n",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output lacks %q:\n%s", want, out)
		}
	}
	wantWarnings := []string{
		`state "c": PlantUML allows one stereotype per state; «mine» is not drawn`,
		`region "r": PlantUML regions are anonymous; its behaviours, stereotype and note are not drawn`,
		`region "r": PlantUML regions are anonymous; the transition a.b -> r into it is not drawn`,
	}
	if !reflect.DeepEqual([]string(warnings), wantWarnings) {
		t.Errorf("warnings =\n%s\nwant\n%s", strings.Join(warnings, "\n"), strings.Join(wantWarnings, "\n"))
	}
}

func TestHelpers(t *testing.T) {
	sm := &model.StateMachine{
		Initial: "a",
		States:  []*model.State{{Name: "a", Kind: model.Normal}, {Name: "b", Parent: "a", Kind: model.Normal}},
		Transitions: []*model.Transition{
			{Source: "a", Targets: []string{"b"}, Event: "go", Cond: "ok", Actions: []string{"x()", "y()"}},
		},
	}
	tmpl := `{{ InitialOf "" }}|{{ (index (Children "a") 0).Name }}|{{ (State "a").Kind }}|` +
		`{{ range OutgoingTransitions "a" }}{{ joinNonEmpty "," .Event (surround "[" .Cond "]") (prefix "/" (join ";" .Actions)) "" }}{{ end }}|` +
		`{{ define "x" }}{{ .Name }}{{ end }}{{ include "x" (State "b") | upper }}|` +
		`{{ len States }}{{ len Transitions }}|{{ join "," (Ancestors "b") }}|{{ CommonAncestor "b" "a" }}|{{ ScopeOf (index Transitions 0) }}` +
		`{{ warn "no %s here" "cheese" }}`
	got, warnings, err := Render(sm, tmpl)
	if err != nil {
		t.Fatal(err)
	}
	if want := "a|b|normal|go,[ok],/x();y()|B|21|a||"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if !reflect.DeepEqual([]string(warnings), []string{"no cheese here"}) {
		t.Errorf("warnings = %q", warnings)
	}
}

func TestErrors(t *testing.T) {
	sm := &model.StateMachine{}
	if _, _, err := Render(sm, "{{ end }}"); err == nil || !strings.Contains(err.Error(), "parsing template") {
		t.Errorf("parse error: %v", err)
	}
	if _, _, err := Render(sm, `{{ include "missing" . }}`); err == nil || !strings.Contains(err.Error(), "executing template") {
		t.Errorf("exec error: %v", err)
	}
}
