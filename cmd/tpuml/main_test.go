package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const example = "../../example/coffee-machine.scxml"

func runCLI(t *testing.T, args ...string) (stdout, stderr string, err error) {
	t.Helper()
	var out, errBuf bytes.Buffer
	err = run(args, &out, &errBuf)
	return out.String(), errBuf.String(), err
}

func TestDefaultTemplateToStdout(t *testing.T) {
	out, _, err := runCLI(t, "-i", example)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(out, "@startuml\n[*] --> idle\n") || !strings.HasSuffix(out, "@enduml\n") {
		t.Errorf("unexpected output:\n%s", out)
	}
}

func TestJSONRoundTripThroughFiles(t *testing.T) {
	dir := t.TempDir()
	jsonPath := filepath.Join(dir, "m.json")
	pumlPath := filepath.Join(dir, "m.puml")

	if _, _, err := runCLI(t, "-i", example, "-F", "json", "-o", jsonPath); err != nil {
		t.Fatal(err)
	}
	if _, _, err := runCLI(t, "-i", jsonPath, "-o", pumlPath); err != nil {
		t.Fatal(err)
	}
	fromJSON, _ := os.ReadFile(pumlPath)
	fromSCXML, _, _ := runCLI(t, "-i", example)
	if string(fromJSON) != fromSCXML {
		t.Errorf("PlantUML from JSON differs from PlantUML from SCXML")
	}
}

func TestCustomTemplateAndExport(t *testing.T) {
	dir := t.TempDir()
	tmpl := filepath.Join(dir, "t.tmpl")
	os.WriteFile(tmpl, []byte(`{{ .Initial }}:{{ len .States }}`), 0o644)
	export := filepath.Join(dir, "e.json")

	out, _, err := runCLI(t, "-i", example, "-t", tmpl, "-e", export)
	if err != nil {
		t.Fatal(err)
	}
	if out != "idle:6" {
		t.Errorf("template output = %q", out)
	}
	if _, err := os.Stat(export); err != nil {
		t.Errorf("export not written: %v", err)
	}
}

func TestErrors(t *testing.T) {
	dir := t.TempDir()
	bad := filepath.Join(dir, "bad.scxml")
	os.WriteFile(bad, []byte(`<scxml initial="nope"><state id="a"><transition target="zzz"/></state></scxml>`), 0o644)
	noext := filepath.Join(dir, "noext")
	os.WriteFile(noext, nil, 0o644)

	cases := []struct {
		args []string
		want string
	}{
		{[]string{}, "-i is required"},
		{[]string{"-i", example, "-t", "x", "-F", "json"}, "mutually exclusive"},
		{[]string{"-i", noext}, "cannot infer the format"},
		{[]string{"-i", noext, "-f", "yaml"}, `unknown input format "yaml"`},
		{[]string{"-i", example, "-F", "yaml"}, `unknown output format "yaml"`},
		{[]string{"-i", filepath.Join(dir, "missing.scxml")}, "reading input"},
		{[]string{"-i", bad}, `unknown target "zzz"`},
		{[]string{"-i", bad}, `unknown initial state "nope"`},
	}
	for _, c := range cases {
		_, _, err := runCLI(t, c.args...)
		if err == nil || !strings.Contains(err.Error(), c.want) {
			t.Errorf("args %v: want error containing %q, got %v", c.args, c.want, err)
		}
	}
}
