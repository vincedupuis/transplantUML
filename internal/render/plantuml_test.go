package render

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// plantUMLArgs returns the argv that runs PlantUML's syntax checker, or skips
// the calling test when PlantUML is not available. A `plantuml` on PATH wins;
// otherwise PLANTUML_JAR points at a jar (see `make plantuml`).
func plantUMLArgs(t *testing.T) []string {
	t.Helper()
	if path, err := exec.LookPath("plantuml"); err == nil {
		return []string{path, "-syntax"}
	}
	if jar := os.Getenv("PLANTUML_JAR"); jar != "" {
		java, err := exec.LookPath("java")
		if err != nil {
			t.Skipf("PLANTUML_JAR is set but java is not in PATH: %v", err)
		}
		return []string{java, "-jar", jar, "-syntax"}
	}
	t.Skip("PlantUML not found; install `plantuml` or run `make plantuml` to set up PLANTUML_JAR")
	return nil
}

// The golden files only prove the template's output has not changed. This runs
// it through PlantUML's own parser, which is the only thing that can say the
// output is a diagram PlantUML will actually accept. The input is rendered here
// rather than read from testdata, so a stale golden cannot mask a break.
//
// -syntax parses without rendering, so Graphviz is not needed. It prints the
// diagram type and entity count on success, and "ERROR" plus the offending line
// on failure.
func TestPlantUMLSyntax(t *testing.T) {
	args := plantUMLArgs(t)
	for input := range goldens(t) {
		t.Run(filepath.Base(input), func(t *testing.T) {
			puml := renderFile(t, input)
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdin = strings.NewReader(puml)
			out, err := cmd.CombinedOutput()
			if err != nil || strings.HasPrefix(string(out), "ERROR") {
				t.Errorf("PlantUML rejected the rendered diagram: %v\n%s\n--- diagram ---\n%s", err, out, puml)
			}
		})
	}
}
