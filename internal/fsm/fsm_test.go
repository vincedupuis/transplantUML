package fsm

import (
	"os"
	"path/filepath"
	"testing"
)

// The generated parser accepts the grammar's own idioms: nested states
// interleaved with the events that concern them, entry/exit actions, guards,
// action lists, and every kind of goto.
func TestGrammarAccepts(t *testing.T) {
	src, err := os.ReadFile("testdata/coffee.fsm")
	if err != nil {
		t.Fatal(err)
	}
	tree, err := parse(string(src))
	if err != nil {
		t.Fatalf("syntax errors: %v", err)
	}
	if got := tree.Identifier().GetText(); got != "coffee" {
		t.Errorf("machine name = %q, want coffee", got)
	}
	if got := len(tree.AllState()); got != 2 {
		t.Errorf("top-level states = %d, want 2", got)
	}
}

func TestGrammarRejects(t *testing.T) {
	for _, src := range []string{
		"fsm {}",                                 // missing name
		"fsm m { state s { on entry goto s } }",  // entry cannot leave the state
		"fsm m { state s { on e [] } }",          // empty guard
		"fsm m { state s { on e goto } }",        // goto without a target
		"fsm m { state s { on e } }",             // neither effect nor target
		"fsm m { state s { on e [g] } }",         // a guard alone is not a transition
		"fsm m { state s { on e goto a/b } }",    // a target is a name, not a path
		"fsm m { state s { after(5s) } }",        // neither effect nor target
		"fsm m { state s { after 5s goto s } }",  // a delay sits in parentheses
		"fsm m { state s { after(5) goto s } }",  // a delay carries its unit
		"fsm m { state s { after(5m) goto s } }", // and the unit is ms or s
		"fsm m { state s { after() goto s } }",   // a delay is not optional
		"fsm m { state s { on after(5s) goto s } }",
		"fsm m { initial }", // initial modifies a state
		"fsm m { state s { initial on e / a } }",
		"fsm m { state initial {} }", // initial is a keyword, not a name
	} {
		if _, err := parse(src); err == nil {
			t.Errorf("%q: accepted, want a syntax error", src)
		}
	}
}

// Two alternatives cover every combination of guard, actions and goto that
// carries an effect or a target; TestGrammarRejects covers the rest.
func TestEventForms(t *testing.T) {
	for _, body := range []string{
		"on e / a",
		"on e / a, b",
		"on e goto t",
		"on e [g] / a",
		"on e [g] goto t",
		"on e / a goto t",
		"on e [g] / a goto t",
		"on entry",
		"on entry / a",
		"on exit / a, b",
		"after(5s) / a",
		"after(250ms) goto t",
		"after(1.5s) [g] / a goto t",
		"after(retryDelay) goto t",
	} {
		if _, err := parse("fsm m { state s { " + body + " } state t {} }"); err != nil {
			t.Errorf("%q: %v", body, err)
		}
	}
}

// Every goto target form lexes: a state name, or one of the three keywords
// standing for a state the document never declares.
func TestGotoTargets(t *testing.T) {
	for _, target := range []string{
		".",
		"final",
		"H",
		"t",
	} {
		if _, err := parse("fsm m { state s { on e goto " + target + " } }"); err != nil {
			t.Errorf("goto %s: %v", target, err)
		}
	}
}

// The documents in example/ are what the README points a reader at, so they
// have to parse.
func TestExamplesParse(t *testing.T) {
	paths, err := filepath.Glob("../../example/*.fsm")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 {
		t.Fatal("no examples found")
	}
	for _, path := range paths {
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := parse(string(src)); err != nil {
			t.Errorf("%s: %v", path, err)
		}
	}
}

// "initial" marks the child a compound state starts in, and the machine's own
// starting state at the top level. It is a named lexer rule declared ahead of
// Identifier, so a state cannot be called "initial" and the visitor gets an
// Initial() accessor.
func TestInitialState(t *testing.T) {
	tree, err := parse("fsm m { initial state a { initial state b {} state c {} } state d {} }")
	if err != nil {
		t.Fatalf("syntax errors: %v", err)
	}
	top := tree.AllState()
	if top[0].Initial() == nil {
		t.Error("a: not marked initial")
	}
	if top[1].Initial() != nil {
		t.Error("d: marked initial")
	}
	inner := top[0].AllState()
	if inner[0].Initial() == nil {
		t.Error("b: not marked initial")
	}
	if inner[1].Initial() != nil {
		t.Error("c: marked initial")
	}
}

// A name is a letter or an underscore followed by letters, digits and
// underscores. The punctuation other formats allow in an id, dots and hyphens,
// is not part of a name here, so "goto ." stays a keyword and no document can
// declare the synthesized "<scope>.final" or "<state>.H".
func TestIdentifierCharset(t *testing.T) {
	src := "fsm vending_machine { state _idle1 { on coin_in [hasChange] / open_1 goto payNow } state payNow {} }"
	if _, err := parse(src); err != nil {
		t.Errorf("charset: %v", err)
	}
	for _, src := range []string{
		"fsm m { state .a {} }",
		"fsm m { state -a {} }",
		"fsm m { state 1a {} }",
		"fsm vending.machine { state a {} }",
		"fsm m { state pay.now {} }",
		"fsm m { state pay-now {} }",
		"fsm m { state a { on coin.in / g goto a } }",
		"fsm m { state a { on e [has-change] / g goto a } }",
		"fsm m { state a { on e / g goto pay.now } state pay.now {} }",
		"fsm m { state a { after(retry.delay) goto a } }",
	} {
		if _, err := parse(src); err == nil {
			t.Errorf("%q: accepted, want a syntax error", src)
		}
	}
}
