package scxml

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const schemaPath = "testdata/schema/scxml.xsd"

// requireXMLLint skips the calling test when xmllint is not installed. It ships
// with macOS and most Linux distributions, so the schema tests normally run.
func requireXMLLint(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("xmllint"); err != nil {
		t.Skip("xmllint not found in PATH; skipping W3C schema validation")
	}
}

// validateSCXML checks src against the vendored W3C SCXML schema. This is the
// only test that judges our documents by something other than ourselves: the
// golden files and round trips can only prove the parser and the emitter agree
// with each other, not that either agrees with the spec.
func validateSCXML(t *testing.T, what string, src []byte) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "doc.scxml")
	if err := os.WriteFile(path, src, 0o644); err != nil {
		t.Fatal(err)
	}
	// xmllint reports the offending line and element on stderr and exits non-zero.
	out, err := exec.Command("xmllint", "--noout", "--schema", schemaPath, path).CombinedOutput()
	if err != nil {
		t.Errorf("%s is not valid SCXML: %v\n%s\n--- document ---\n%s", what, err, out, src)
	}
}

// Every document we read and every document we write must validate against the
// W3C schema.
func TestSCXMLValidatesAgainstW3CSchema(t *testing.T) {
	requireXMLLint(t)

	examples, err := filepath.Glob("../../example/*.scxml")
	if err != nil || len(examples) == 0 {
		t.Fatalf("no examples found: %v", err)
	}
	for _, path := range append(examples, "testdata/edge.scxml", "testdata/uml.scxml") {
		t.Run(filepath.Base(path), func(t *testing.T) {
			src, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			validateSCXML(t, "fixture "+path, src)
			validateSCXML(t, "emitter output for "+path, emit(t, parseFile(t, path)))
		})
	}

	// Executable content is the emitter's most intricate path: <script> bodies,
	// raw XML re-inserted as elements, and a targetless internal transition.
	t.Run("executable-content", func(t *testing.T) {
		validateSCXML(t, "emitter output for the executable-content model", emit(t, execContentSM()))
	})
}
