package format

import "testing"

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
	if _, err := EmitterFor("scxml"); err == nil {
		t.Error("expected error for unknown emitter")
	}
}
