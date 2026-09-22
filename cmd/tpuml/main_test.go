package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	example = "../../example/coffee-machine.scxml"
	uml     = "../../internal/scxml/testdata/uml.scxml"
)

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

// Every built-in output format can be read back in: converting a document
// through it must yield the same PlantUML as rendering the document directly.
func TestRoundTripThroughFiles(t *testing.T) {
	for _, input := range []string{example, uml} {
		fromSCXML, _, err := runCLI(t, "-i", input)
		if err != nil {
			t.Fatal(err)
		}
		for _, name := range []string{"json", "scxml"} {
			t.Run(filepath.Base(input)+"/"+name, func(t *testing.T) {
				dir := t.TempDir()
				midPath := filepath.Join(dir, "m."+name)
				pumlPath := filepath.Join(dir, "m.puml")

				if _, _, err := runCLI(t, "-i", input, "-F", name, "-o", midPath); err != nil {
					t.Fatal(err)
				}
				if _, _, err := runCLI(t, "-i", midPath, "-o", pumlPath); err != nil {
					t.Fatal(err)
				}
				got, err := os.ReadFile(pumlPath)
				if err != nil {
					t.Fatal(err)
				}
				if string(got) != fromSCXML {
					t.Errorf("PlantUML via %s differs from PlantUML from SCXML:\n%s", name, got)
				}
			})
		}
	}
}

// Warnings from the parser and from the output go to stderr, prefixed, and
// never fail the conversion.
func TestWarningsOnStderr(t *testing.T) {
	out, stderr, err := runCLI(t, "-i", uml)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(out, "@startuml\n") {
		t.Errorf("unexpected output:\n%s", out)
	}
	for _, want := range []string{
		"tpuml: warning: state \"failed\": <donedata> is not supported and was dropped\n",
		"tpuml: warning: state \"stop\": PlantUML has no terminate symbol",
	} {
		if !strings.Contains(stderr, want) {
			t.Errorf("stderr missing %q:\n%s", want, stderr)
		}
	}
	if _, stderr, err := runCLI(t, "-i", uml, "-F", "scxml"); err != nil || !strings.Contains(stderr, "SCXML has no deferred events") {
		t.Errorf("emitter warnings not reported: %v\n%s", err, stderr)
	}
	if _, stderr, err := runCLI(t, "-i", example); err != nil || stderr != "" {
		t.Errorf("a plain document must not warn: %v\n%s", err, stderr)
	}
}

func TestCustomTemplate(t *testing.T) {
	dir := t.TempDir()
	tmpl := filepath.Join(dir, "t.tmpl")
	os.WriteFile(tmpl, []byte(`{{ .Initial }}:{{ len .States }}`), 0o644)

	out, _, err := runCLI(t, "-i", example, "-t", tmpl)
	if err != nil {
		t.Fatal(err)
	}
	if out != "idle:6" {
		t.Errorf("template output = %q", out)
	}
}

func TestNoArgsPrintsUsage(t *testing.T) {
	out, _, err := runCLI(t)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Usage:", "-i, --input", "-F, --output-format", "-h, --help"} {
		if !strings.Contains(out, want) {
			t.Errorf("usage output missing %q:\n%s", want, out)
		}
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
		{[]string{"-i", example, "-t", "x", "-F", "json"}, "none of the others can be"},
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
