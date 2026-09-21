package format

import (
	"reflect"
	"testing"
)

func TestDetect(t *testing.T) {
	cases := map[string]string{"a.scxml": "scxml", "A.XML": "scxml", "m.json": "json", "x.puml": "", "noext": ""}
	for in, want := range cases {
		if got := Detect(in); got != want {
			t.Errorf("Detect(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestLookup(t *testing.T) {
	if _, err := ParserFor("scxml"); err != nil {
		t.Error(err)
	}
	if _, err := ParserFor("yaml"); err == nil {
		t.Error("expected error for unknown parser")
	}
	if _, err := EmitterFor("json"); err != nil {
		t.Error(err)
	}
	if _, err := EmitterFor("yaml"); err == nil {
		t.Error("expected error for unknown emitter")
	}
}

// Every format tpuml can read it must also be able to write, and vice versa;
// only template output is one-way.
func TestFormatsGoBothWays(t *testing.T) {
	if in, out := ParserNames(), EmitterNames(); !reflect.DeepEqual(in, out) {
		t.Errorf("parsers %v but emitters %v", in, out)
	}
	for ext, name := range extensions {
		if _, err := EmitterFor(name); err != nil {
			t.Errorf("%s: %v", ext, err)
		}
	}
}
