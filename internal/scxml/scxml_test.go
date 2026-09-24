package scxml

import (
	"os"
	"reflect"
	"slices"
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
	sm, _, err := Parser{}.Parse(src)
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

// uml.scxml holds every UML concept the model has, using the tpuml extension
// vocabulary where SCXML has nothing native.
func TestUMLConcepts(t *testing.T) {
	src, err := os.ReadFile("testdata/uml.scxml")
	if err != nil {
		t.Fatal(err)
	}
	sm, warnings, err := Parser{}.Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := sm.Validate(); err != nil {
		t.Fatalf("model does not validate: %v", err)
	}
	if want := []string{
		`state "outer": a note about "initial" is not supported and was dropped`,
		`state "outer": a note about a deferred event names no event and was dropped`,
		`state "failed": <donedata> is not supported and was dropped`,
	}; !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings = %q, want %q", warnings, want)
	}

	wantKinds := map[string]model.StateKind{
		"check": model.Choice, "split": model.Fork, "merge": model.Junction, "sync": model.Join,
		"in": model.EntryPoint, "out": model.ExitPoint, "stop": model.Terminate, "done": model.Final,
		"work": model.Normal, "both": model.Parallel, "sub": model.Normal, "inner": model.Normal,
	}
	for name, want := range wantKinds {
		if got := sm.State(name).Kind; got != want {
			t.Errorf("%s: kind %q, want %q", name, got, want)
		}
	}

	if sm.Note != "Every UML concept the model holds,\nin one document." {
		t.Errorf("machine note = %q", sm.Note)
	}
	wantVars := []model.Variable{{Name: "retries", Value: "0"}, {Name: "config", Value: "src(config.json)"}}
	if !reflect.DeepEqual(sm.Variables, wantVars) {
		t.Errorf("machine variables = %+v", sm.Variables)
	}

	work := sm.State("work")
	if work.Stereotype != "worker" || work.Invariant != "retries >= 0" || work.Note != "Runs the job." {
		t.Errorf("work annotations = %+v", *work)
	}
	if !reflect.DeepEqual(work.Defer, []string{"pause", "resume"}) {
		t.Errorf("work.Defer = %q", work.Defer)
	}
	// A note about something the state lists sits beside the state's own.
	if work.EntryNote != "Starts the clock." || work.DoNote != "Runs in a worker." || work.InvariantNote != "Never negative." ||
		!reflect.DeepEqual(work.DeferNotes, map[string]string{"pause": "Kept until the job ends."}) {
		t.Errorf("work behaviour notes = %+v", *work)
	}
	if got := sm.State("outer"); got.Initial != "inner" || !slices.Equal(got.InitialActions, []string{"log('hi')"}) {
		t.Errorf("outer initial = %q with %q, want inner with the log", got.Initial, got.InitialActions)
	}
	// done.invoke of the submachine's invoke and done.state of the state itself
	// are UML's completion event.
	for _, source := range []string{"sub", "outer"} {
		found := false
		for _, tr := range sm.OutgoingTransitions(source) {
			found = found || (tr.Event == "" && slices.Equal(tr.Targets, []string{"done"}))
		}
		if !found {
			t.Errorf("%s: no completion transition to done", source)
		}
	}
	if got := sm.State("outer").ExitNote; got != "Says goodbye." {
		t.Errorf("outer exit note = %q", got)
	}
	if !reflect.DeepEqual(work.Variables, []model.Variable{{Name: "progress", Value: "0"}}) {
		t.Errorf("work.Variables = %+v", work.Variables)
	}
	if !reflect.DeepEqual(work.Do, []string{"invoke(job.py, http://example.com/worker)"}) {
		t.Errorf("work.Do = %q", work.Do)
	}
	// The timer's <send> and <cancel> became the time trigger, not actions.
	if !reflect.DeepEqual(work.OnEntry, []string{"log('start')"}) || len(work.OnExit) != 0 {
		t.Errorf("work entry/exit = %q / %q", work.OnEntry, work.OnExit)
	}
	timed := sm.OutgoingTransitions("work")[0]
	if timed.After != "5s" || timed.Event != "" || timed.Note != "Timed out." || !reflect.DeepEqual(timed.Actions, []string{"retries = retries + 1"}) {
		t.Errorf("time trigger = %+v", *timed)
	}
	// A delayexpr is a delay the engine computes; it reaches After like a plain
	// delay does, and the emitter puts it back in delayexpr.
	if computed := sm.OutgoingTransitions("work")[1]; computed.After != "retryDelay" || computed.Event != "" {
		t.Errorf("computed time trigger = %+v", *computed)
	}

	if got := sm.State("sub").Submachine; got != "child.scxml" {
		t.Errorf("submachine = %q", got)
	}

	inner := sm.OutgoingTransitions("inner")
	if !inner[0].IsExternal() || !inner[1].IsLocal() || !inner[2].IsInternal() {
		t.Errorf("transition kinds = %q %q %q", inner[0].Kind, inner[1].Kind, inner[2].Kind)
	}
	if inner[2].Note != "Not drawn." {
		t.Errorf("internal transition note = %q", inner[2].Note)
	}
}

// The choice/fork idiom must not swallow states that merely look transient.
func TestConnectorIdiom(t *testing.T) {
	cases := map[string]struct {
		body string
		want model.StateKind
	}{
		"guarded branches":    {`<transition cond="a" target="x"/><transition cond="b" target="y"/>`, model.Choice},
		"branches with else":  {`<transition cond="a" target="x"/><transition target="y"/>`, model.Choice},
		"multi-target":        {`<transition target="x y"/>`, model.Fork},
		"single pass-through": {`<transition target="x"/>`, model.Normal},
		"unguarded branches":  {`<transition target="x"/><transition target="y"/>`, model.Normal},
		"has an event":        {`<transition event="e" cond="a" target="x"/><transition cond="b" target="y"/>`, model.Normal},
		"has entry action":    {`<onentry><log/></onentry><transition cond="a" target="x"/><transition cond="b" target="y"/>`, model.Normal},
		"targetless":          {`<transition cond="a"/><transition cond="b" target="y"/>`, model.Normal},
		"opted out":           {`<transition tpuml:kind="x" cond="a" target="x"/><transition cond="b" target="y"/>`, model.Choice},
		"forced normal":       {`<transition cond="a" target="x"/><transition cond="b" target="y"/>`, model.Normal},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			attr := ""
			if name == "forced normal" {
				attr = ` tpuml:kind="normal"`
			}
			src := `<scxml xmlns="http://www.w3.org/2005/07/scxml" xmlns:tpuml="` + ExtNamespace + `" initial="s">` +
				`<state id="s"` + attr + `>` + c.body + `</state><state id="x"/><state id="y"/></scxml>`
			sm, _, err := Parser{}.Parse([]byte(src))
			if err != nil {
				t.Fatal(err)
			}
			if got := sm.State("s").Kind; got != c.want {
				t.Errorf("kind = %q, want %q", got, c.want)
			}
		})
	}
}

// The extension attributes are recognised by namespace, whatever the prefix.
func TestExtensionPrefix(t *testing.T) {
	src := `<scxml xmlns="http://www.w3.org/2005/07/scxml" xmlns:x="` + ExtNamespace + `" xmlns:o="urn:other">` +
		`<state id="s" x:kind="junction" o:kind="ignored"><x:note>hi</x:note><transition target="s"/></state></scxml>`
	sm, _, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	if s := sm.State("s"); s.Kind != model.Junction || s.Note != "hi" {
		t.Errorf("state = %+v", *s)
	}
}

func TestParserWarnings(t *testing.T) {
	src := `<scxml xmlns="http://www.w3.org/2005/07/scxml"><script>x()</script><state id="s"><onentry/><foo/></state></scxml>`
	_, warnings, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	want := []string{`state "s": <foo> is not supported and was dropped`, `<script> at the top level is not supported and was dropped`}
	if !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings = %q, want %q", warnings, want)
	}
}

func TestInvokeForms(t *testing.T) {
	src := `<scxml xmlns="http://www.w3.org/2005/07/scxml"><state id="s">
	  <invoke src="a.scxml"/>
	  <invoke type="http://www.w3.org/TR/scxml/" src="b.scxml"/>
	  <invoke type="x" srcexpr="'c'"/>
	  <invoke src="d.scxml"><param name="p" expr="1"/></invoke>
	</state></scxml>`
	sm, _, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	s := sm.State("s")
	if s.Submachine != "a.scxml" {
		t.Errorf("submachine = %q", s.Submachine)
	}
	want := []string{"invoke(b.scxml)", "invoke('c', x)", `<invoke src="d.scxml"><param name="p" expr="1"/></invoke>`}
	if !reflect.DeepEqual(s.Do, want) {
		t.Errorf("Do = %q, want %q", s.Do, want)
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
	sm, _, err := Parser{}.Parse([]byte(src))
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
	if !tr[0].IsInternal() || !reflect.DeepEqual(tr[0].Actions, []string{"log(evt: _event.name)"}) {
		t.Errorf("internal transition = %+v", *tr[0])
	}
	if tr[1].IsInternal() || len(tr[1].Targets) != 0 {
		t.Errorf("targetless transition = %+v", *tr[1])
	}
}

func TestErrors(t *testing.T) {
	if _, _, err := (Parser{}).Parse([]byte("<scxml>")); err == nil || !strings.Contains(err.Error(), "parsing XML") {
		t.Errorf("malformed XML: %v", err)
	}
	if _, _, err := (Parser{}).Parse([]byte("<root/>")); err == nil || !strings.Contains(err.Error(), "<scxml>") {
		t.Errorf("missing root: %v", err)
	}
}

// A bare done.invoke matches every invoke of its state, so it is the
// submachine's completion only when nothing else is invoked there.
func TestCompletionEvent(t *testing.T) {
	for src, want := range map[string]string{
		`<state id="s"><invoke src="m.scxml"/><transition event="done.invoke" target="s"/></state>`:                                "",
		`<state id="s"><invoke src="m.scxml"/><transition event="done.invoke.*" target="s"/></state>`:                              "",
		`<state id="s"><invoke id="i" src="m.scxml"/><transition event="done.invoke.j" target="s"/></state>`:                       "done.invoke.j",
		`<state id="s"><invoke src="m.scxml"/><invoke src="job.py" type="x"/><transition event="done.invoke" target="s"/></state>`: "done.invoke",
		`<state id="s"><invoke src="job.py" type="x"/><transition event="done.invoke" target="s"/></state>`:                        "done.invoke",
		`<state id="s"><state id="a"/><transition event="done.state.a" target="s"/></state>`:                                       "done.state.a",
	} {
		sm, _, err := Parser{}.Parse([]byte(`<scxml xmlns="http://www.w3.org/2005/07/scxml" version="1.0">` + src + `</scxml>`))
		if err != nil {
			t.Fatal(err)
		}
		if got := sm.OutgoingTransitions("s")[0].Event; got != want {
			t.Errorf("%s: event = %q, want %q", src, got, want)
		}
	}
}
