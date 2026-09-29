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

	if got := sm.State("a").OnEntry; !reflect.DeepEqual(got, []string{"sayHi"}) {
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
		`state "work": the type "http://example.com/worker" of the do activity "runJob" is not kept`,
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
	if work.Stereotype != "worker" || work.Invariant != "counting and not (stalled or aborted)" || work.Note != "Runs the job." {
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
	if got := sm.State("outer"); got.Initial != "inner" || !slices.Equal(got.InitialActions, []string{"greet"}) {
		t.Errorf("outer initial = %q with %q, want inner with greet", got.Initial, got.InitialActions)
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
	if !reflect.DeepEqual(work.Do, []string{"runJob"}) {
		t.Errorf("work.Do = %q", work.Do)
	}
	// The timer's <send> and <cancel> became the time trigger, not actions.
	if !reflect.DeepEqual(work.OnEntry, []string{"startClock"}) || len(work.OnExit) != 0 {
		t.Errorf("work entry/exit = %q / %q", work.OnEntry, work.OnExit)
	}
	timed := sm.OutgoingTransitions("work")[0]
	if timed.After != "5s" || timed.Event != "" || timed.Note != "Timed out." || !reflect.DeepEqual(timed.Actions, []string{"countRetry"}) {
		t.Errorf("time trigger = %+v", *timed)
	}
	// A delayexpr is a delay the engine computes; it reaches After like a plain
	// delay does, and the emitter puts it back in delayexpr.
	if computed := sm.OutgoingTransitions("work")[1]; computed.After != "retryDelay" || computed.Event != "" {
		t.Errorf("computed time trigger = %+v", *computed)
	}

	if got := sm.State("sub").Submachine; got != "child" {
		t.Errorf("submachine = %q", got)
	}

	// type="internal" is a local transition with a target nested in the
	// source, and an internal one without a target.
	inner := sm.OutgoingTransitions("inner")
	if again := sm.OutgoingTransitions("outer")[1]; !inner[0].IsExternal() || !again.IsLocal() || !inner[1].IsInternal() {
		t.Errorf("transition kinds = %q %q %q", inner[0].Kind, again.Kind, inner[1].Kind)
	}
	if inner[1].Note != "Not drawn." {
		t.Errorf("internal transition note = %q", inner[1].Note)
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
		"has entry action":    {`<onentry><script>a</script></onentry><transition cond="a" target="x"/><transition cond="b" target="y"/>`, model.Normal},
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
			src := `<scxml xmlns="http://www.w3.org/2005/07/scxml" xmlns:tpuml="` + ExtNamespace + `" name="m" initial="s">` +
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

// An SCXML element the model has no place for would change what the machine
// does, and so would a note the parser cannot place. An element of another
// namespace, which engines ignore, is dropped with a warning.
func TestParserUnsupported(t *testing.T) {
	src := `<scxml xmlns="http://www.w3.org/2005/07/scxml" xmlns:tpuml="https://github.com/vincedupuis/transplantUML" xmlns:qt="http://www.qt.io/2015/02/scxml-ext">
	  <script>x</script>
	  <state id="s"><onentry/><foo/><qt:editorinfo/>
	    <tpuml:note about="initial">Not a thing a state lists.</tpuml:note>
	    <tpuml:note about="defer">Names no event.</tpuml:note>
	  </state>
	  <final id="f"><donedata><content>failed</content></donedata></final>
	</scxml>`
	_, warnings, err := Parser{}.Parse([]byte(src))
	want := []string{
		`state "s": a note about "initial" is not supported; it is about entry, exit, do, invariant or defer`,
		`state "s": a note about a deferred event names no event`,
		`<foo> in state "s" is not supported`,
		`<donedata> in state "f" is not supported`,
		`<script> at the top level is not supported`,
	}
	if err == nil || !reflect.DeepEqual(strings.Split(err.Error(), "\n"), want) {
		t.Errorf("errors =\n%v\nwant\n%s", err, strings.Join(want, "\n"))
	}
	wantWarnings := []string{`<qt:editorinfo> in state "s" belongs to "http://www.qt.io/2015/02/scxml-ext", which SCXML engines ignore; it was dropped`}
	if !reflect.DeepEqual([]string(warnings), wantWarnings) {
		t.Errorf("warnings = %q, want %q", warnings, wantWarnings)
	}
}

// A tpuml:kind="initial" state holds the machine's initial transition and
// nothing else, and the machine starts in it.
func TestMachineInitialErrors(t *testing.T) {
	for want, body := range map[string]string{
		"is the one the machine starts in": `initial="a"><state id="i" tpuml:kind="initial"><transition target="a"/></state><state id="a"/>`,
		"holds only one transition":        `initial="i"><state id="i" tpuml:kind="initial"><transition event="e" target="a"/></state><state id="a"/>`,
	} {
		src := `<scxml xmlns="http://www.w3.org/2005/07/scxml" xmlns:tpuml="https://github.com/vincedupuis/transplantUML" name="m" ` + body + `</scxml>`
		if _, _, err := (Parser{}).Parse([]byte(src)); err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("want an error containing %q, got %v", want, err)
		}
	}
}

// type="internal" with a target keeps the source active only when it is
// compound and every target lies inside it; otherwise SCXML runs it as an
// external transition.
func TestTypeInternalWithTarget(t *testing.T) {
	src := `<scxml xmlns="http://www.w3.org/2005/07/scxml" initial="c">
	  <state id="c">
	    <transition event="in" target="c1" type="internal"/>
	    <transition event="self" target="c" type="internal"/>
	    <state id="c1"><transition event="leaf" target="c1" type="internal"/></state>
	  </state>
	</scxml>`
	sm, _, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	for i, want := range []model.TransitionKind{model.Local, model.External, model.External} {
		if got := sm.Transitions[i].Kind; got != want {
			t.Errorf("%s: kind %q, want %q", sm.Transitions[i].Event, got, want)
		}
	}
}

// Invoking an SCXML document makes a submachine state, and invoking anything
// else is a do activity named by its src. Only the emitter's own type is kept
// without a warning.
func TestInvokeForms(t *testing.T) {
	src := `<scxml xmlns="http://www.w3.org/2005/07/scxml"><state id="s">
	  <invoke type="http://www.w3.org/TR/scxml/" src="a"/>
	  <invoke type="tpuml:do" src="spin"/>
	  <invoke type="x" src="blink"/>
	</state></scxml>`
	sm, warnings, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	s := sm.State("s")
	if s.Submachine != "a" {
		t.Errorf("submachine = %q", s.Submachine)
	}
	if want := []string{"spin", "blink"}; !reflect.DeepEqual(s.Do, want) {
		t.Errorf("Do = %q, want %q", s.Do, want)
	}
	if want := []string{`state "s": the type "x" of the do activity "blink" is not kept`}; !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings = %q, want %q", warnings, want)
	}
}

// An action is a <script> holding a name, and an internal transition may
// have one too.
func TestExecutableContentAndInternal(t *testing.T) {
	src := `<scxml initial="s">
	  <state id="s">
	    <onentry><script> start </script><script>count</script></onentry>
	    <onexit><script>stop</script></onexit>
	    <transition event="e" type="internal"><script>log</script></transition>
	    <transition event="nowhere"/>
	  </state>
	</scxml>`
	sm, _, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	s := sm.State("s")
	if got := s.OnEntry; !reflect.DeepEqual(got, []string{"start", "count"}) {
		t.Errorf("OnEntry = %q", got)
	}
	if got := s.OnExit; !reflect.DeepEqual(got, []string{"stop"}) {
		t.Errorf("OnExit = %q", got)
	}
	tr := sm.OutgoingTransitions("s")
	if !tr[0].IsInternal() || !reflect.DeepEqual(tr[0].Actions, []string{"log"}) {
		t.Errorf("internal transition = %+v", *tr[0])
	}
	if tr[1].IsInternal() || len(tr[1].Targets) != 0 {
		t.Errorf("targetless transition = %+v", *tr[1])
	}
}

// Executable content other than a <script>, and an <invoke> that is neither
// a submachine nor a named do activity, are errors of the parser.
func TestExpressionErrors(t *testing.T) {
	src, err := os.ReadFile("testdata/expressions.scxml")
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = Parser{}.Parse(src)
	if err == nil {
		t.Fatal("no error")
	}
	want := []string{
		`state "s": <log> in <initial> is not supported; an action is a <script> holding a name`,
		`state "s": a state runs one submachine, but it invokes "a" and "b"`,
		`state "s": an <invoke> with content or without src is not supported; a do activity is an <invoke> whose src is a name`,
		`state "s": <log> in <transition> is not supported; an action is a <script> holding a name`,
		`state "s": <assign> in <onentry> is not supported; an action is a <script> holding a name`,
		`state "s": <raise> in <onentry> is not supported; an action is a <script> holding a name`,
		`state "s": <send> in <onentry> is not supported; an action is a <script> holding a name`,
		`state "s": <if> in <onentry> is not supported; an action is a <script> holding a name`,
		`state "s": <cancel> in <onexit> is not supported; an action is a <script> holding a name`,
		`state "s": <script src> in <onexit> is not supported; an action is a <script> holding a name`,
	}
	if got := strings.Split(err.Error(), "\n"); !reflect.DeepEqual(got, want) {
		t.Errorf("errors =\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

// A <script>, a cond or a tpuml:invariant that holds more than names parses,
// and the model rejects it.
func TestExpressionsInvalid(t *testing.T) {
	src := `<scxml xmlns="http://www.w3.org/2005/07/scxml" xmlns:tpuml="` + ExtNamespace + `" name="m" initial="s">
	  <state id="s" tpuml:invariant="n >= 0">
	    <onentry><script>x = 1;</script></onentry>
	    <transition event="e" cond="n > 3" target="s"/>
	  </state>
	</scxml>`
	sm, _, err := Parser{}.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	err = sm.Validate()
	want := []string{
		`state "s": the entry action "x = 1;" is not a name`,
		`state "s": the invariant [n >= 0] is not made of names, not, and, or and parentheses`,
		`transition #0 (s): the guard [n > 3] is not made of names, not, and, or and parentheses`,
	}
	if err == nil || !reflect.DeepEqual(strings.Split(err.Error(), "\n"), want) {
		t.Errorf("errors = %v, want\n%s", err, strings.Join(want, "\n"))
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
		`<state id="s"><invoke src="m"/><transition event="done.invoke" target="s"/></state>`:                                "",
		`<state id="s"><invoke src="m"/><transition event="done.invoke.*" target="s"/></state>`:                              "",
		`<state id="s"><invoke id="i" src="m"/><transition event="done.invoke.j" target="s"/></state>`:                       "done.invoke.j",
		`<state id="s"><invoke src="m"/><invoke src="job.py" type="x"/><transition event="done.invoke" target="s"/></state>`: "done.invoke",
		`<state id="s"><invoke src="job.py" type="x"/><transition event="done.invoke" target="s"/></state>`:                  "done.invoke",
		`<state id="s"><state id="a"/><transition event="done.state.a" target="s"/></state>`:                                 "done.state.a",
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
