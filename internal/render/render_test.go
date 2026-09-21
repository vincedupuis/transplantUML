package render

import (
	"os"
	"strings"
	"testing"

	"github.com/vincedupuis/transplantUML/assets"
	"github.com/vincedupuis/transplantUML/internal/model"
	"github.com/vincedupuis/transplantUML/internal/scxml"
)

func renderSCXML(t *testing.T, path string) string {
	t.Helper()
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	sm, err := scxml.Parser{}.Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	out, err := Render(sm, assets.PlantUML)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// Golden-file tests for the built-in PlantUML template. Regenerate a golden
// file with: go run ./cmd/tpuml -i <input> -o internal/render/testdata/<name>.puml
func TestPlantUMLGolden(t *testing.T) {
	cases := map[string]string{
		"coffee-machine": "../../example/coffee-machine.scxml",
		"edge":           "../scxml/testdata/edge.scxml",
	}
	for name, path := range cases {
		t.Run(name, func(t *testing.T) {
			want, err := os.ReadFile("testdata/" + name + ".puml")
			if err != nil {
				t.Fatal(err)
			}
			if got := renderSCXML(t, path); got != string(want) {
				t.Errorf("output differs from testdata/%s.puml\n--- got ---\n%s--- want ---\n%s", name, got, want)
			}
		})
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
		`{{ define "x" }}{{ .Name }}{{ end }}{{ include "x" (State "b") | upper }}`
	got, err := Render(sm, tmpl)
	if err != nil {
		t.Fatal(err)
	}
	if want := "a|b|normal|go,[ok],/x();y()|B"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestErrors(t *testing.T) {
	sm := &model.StateMachine{}
	if _, err := Render(sm, "{{ end }}"); err == nil || !strings.Contains(err.Error(), "parsing template") {
		t.Errorf("parse error: %v", err)
	}
	if _, err := Render(sm, `{{ include "missing" . }}`); err == nil || !strings.Contains(err.Error(), "executing template") {
		t.Errorf("exec error: %v", err)
	}
}
