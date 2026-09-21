package scxml

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/vincedupuis/transplantUML/internal/model"
)

func parseFile(t *testing.T, path string) *model.StateMachine {
	t.Helper()
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	sm, err := Parser{}.Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := sm.Validate(); err != nil {
		t.Fatalf("model does not validate: %v", err)
	}
	return sm
}

func TestCoffeeMachine(t *testing.T) {
	sm := parseFile(t, "../../example/coffee-machine.scxml")
	if sm.Initial != "idle" {
		t.Errorf("Initial = %q", sm.Initial)
	}
	if len(sm.States) != 6 {
		t.Errorf("want 6 states, got %d", len(sm.States))
	}
	if len(sm.Transitions) != 14 {
		t.Errorf("want 14 transitions, got %d", len(sm.Transitions))
	}
	for _, s := range sm.States {
		if s.Parent != "" || s.Kind != model.Normal || s.Initial != "" {
			t.Errorf("unexpected state %+v", *s)
		}
	}
	want := &model.Transition{Source: "idle", Targets: []string{"ready"}, Event: "power_on"}
	if got := sm.OutgoingTransitions("idle"); len(got) != 1 || !reflect.DeepEqual(got[0], want) {
		t.Errorf("idle transitions = %+v", got)
	}
}

func TestEdgeCases(t *testing.T) {
	sm := parseFile(t, "testdata/edge.scxml")

	if sm.Name != "edge" || sm.Initial != "a" {
		t.Errorf("name/initial = %q/%q (initial must come from the <initial> element)", sm.Name, sm.Initial)
	}

	kinds := map[string]model.StateKind{}
	parents := map[string]string{}
	for _, s := range sm.States {
		kinds[s.Name] = s.Kind
		parents[s.Name] = s.Parent
	}
	wantKinds := map[string]model.StateKind{
		"a": model.Normal, "p": model.Parallel, "r1": model.Normal, "r1a": model.Normal,
		"r2": model.Normal, "r2a": model.Normal, "b": model.Normal, "h": model.HistoryDeep,
		"b1": model.Normal, "f": model.Final,
	}
	if !reflect.DeepEqual(kinds, wantKinds) {
		t.Errorf("kinds = %v", kinds)
	}
	if parents["r1"] != "p" || parents["r1a"] != "r1" || parents["h"] != "b" || parents["f"] != "b" {
		t.Errorf("parents = %v", parents)
	}

	if got := sm.State("a").OnEntry; !reflect.DeepEqual(got, []string{"log('hi')"}) {
		t.Errorf("a.OnEntry = %q", got)
	}
	if got := sm.State("b").Initial; got != "b1" {
		t.Errorf("b.Initial = %q (must default to the first child state)", got)
	}
	if got := sm.State("r1").Initial; got != "r1a" {
		t.Errorf("r1.Initial = %q", got)
	}
	if got := sm.State("p").Initial; got != "" {
		t.Errorf("parallel states have no initial, got %q", got)
	}

	// The <initial> element's transition must not become an ordinary transition.
	for _, tr := range sm.Transitions {
		if tr.Source == "" {
			t.Errorf("transition with empty source: %+v", *tr)
		}
	}
	multi := sm.OutgoingTransitions("a")[1]
	if !reflect.DeepEqual(multi.Targets, []string{"a", "b"}) {
		t.Errorf("multi-target = %v", multi.Targets)
	}
	hist := sm.OutgoingTransitions("h")
	if len(hist) != 1 || hist[0].Targets[0] != "b1" {
		t.Errorf("history default transition = %+v", hist)
	}
	done := sm.OutgoingTransitions("p")
	if len(done) != 1 || done[0].Targets[0] != "h" || done[0].Event != "done" {
		t.Errorf("parallel transition = %+v", done)
	}
}

func TestExecutableContentAndInternal(t *testing.T) {
	src := `<scxml initial="s">
	  <state id="s">
	    <onentry>
	      <script> x = 1; </script>
	      <assign location="n" expr="n + 1"/>
	      <raise event="tick"/>
	      <send event="ping"/>
	      <cancel sendid="ping"/>
	      <if cond="n > 3"><log expr="'big'"/></if>
	    </onentry>
	    <onexit><script src="cleanup.js"/></onexit>
	    <transition event="e" type="internal"><log label="evt" expr="_event.name"/></transition>
	    <transition event="nowhere"/>
	  </state>
	</scxml>`
	sm, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	s := sm.State("s")
	want := []string{"x = 1;", "n = n + 1", "raise tick", "send ping", "cancel ping"}
	if got := s.OnEntry[:5]; !reflect.DeepEqual(got, want) {
		t.Errorf("OnEntry = %q", got)
	}
	if got := s.OnEntry[5]; !strings.HasPrefix(got, `<if cond=`) || !strings.Contains(got, `<log expr=`) {
		t.Errorf("unknown content should be kept as XML, got %q", got)
	}
	if got := s.OnExit; !reflect.DeepEqual(got, []string{"script(cleanup.js)"}) {
		t.Errorf("OnExit = %q", got)
	}
	tr := sm.OutgoingTransitions("s")
	if !tr[0].Internal || !reflect.DeepEqual(tr[0].Actions, []string{"log(evt: _event.name)"}) {
		t.Errorf("internal transition = %+v", *tr[0])
	}
	if tr[1].Internal || len(tr[1].Targets) != 0 {
		t.Errorf("targetless transition = %+v", *tr[1])
	}
}

func TestErrors(t *testing.T) {
	if _, err := (Parser{}).Parse([]byte("<scxml>")); err == nil || !strings.Contains(err.Error(), "parsing XML") {
		t.Errorf("malformed XML: %v", err)
	}
	if _, err := (Parser{}).Parse([]byte("<root/>")); err == nil || !strings.Contains(err.Error(), "<scxml>") {
		t.Errorf("missing root: %v", err)
	}
}
