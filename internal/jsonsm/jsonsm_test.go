package jsonsm

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/vincedupuis/transplantUML/internal/model"
	"github.com/vincedupuis/transplantUML/internal/scxml"
)

// SCXML -> model -> JSON -> model must be lossless.
func TestRoundTrip(t *testing.T) {
	for _, path := range []string{"../../example/coffee-machine.scxml", "../scxml/testdata/edge.scxml"} {
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		want, err := scxml.Parser{}.Parse(src)
		if err != nil {
			t.Fatal(err)
		}
		out, err := Emitter{}.Emit(want)
		if err != nil {
			t.Fatal(err)
		}
		got, err := Parser{}.Parse(out)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: round trip differs\n%s", path, out)
		}
	}
}

func TestParseDefaults(t *testing.T) {
	sm, err := Parser{}.Parse([]byte(`{"initial":"a","states":[{"name":"a"}]}`))
	if err != nil {
		t.Fatal(err)
	}
	if sm.States[0].Kind != model.Normal || sm.Transitions == nil {
		t.Errorf("defaults not applied: %+v", sm)
	}
	if err := sm.Validate(); err != nil {
		t.Error(err)
	}
}

func TestParseRejectsUnknownFields(t *testing.T) {
	_, err := Parser{}.Parse([]byte(`{"states":[{"name":"a","colour":"red"}]}`))
	if err == nil || !strings.Contains(err.Error(), "colour") {
		t.Errorf("want unknown field error, got %v", err)
	}
}
