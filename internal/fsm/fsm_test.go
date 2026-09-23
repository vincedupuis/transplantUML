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
		"fsm {}",                                // missing name
		"fsm m { state s { on entry goto s } }", // entry cannot leave the state
		"fsm m { state s { on e [] } }",         // empty guard
		"fsm m { state s { on e goto } }",       // goto without a target
		"fsm m { state s { on e } }",            // neither effect nor target
		"fsm m { state s { on e [g] } }",        // a guard alone is not a transition
		"fsm m { state s { on e goto a/b } }",   // a target is a name, not a path
		"fsm m { initial }",                     // initial modifies a state
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

// Identifier carries every name the language can write, so the punctuation a
// source document may already use — underscore, dot, hyphen — has to lex in
// each of those positions. A name still begins with a letter or an underscore,
// which is what keeps "goto ." a keyword rather than a name.
func TestIdentifierCharset(t *testing.T) {
	src := "fsm vending.machine { state _idle-1 { on coin.in [has-change] / open_1 goto pay.now } state pay.now {} }"
	if _, err := parse(src); err != nil {
		t.Errorf("charset: %v", err)
	}
	for _, src := range []string{
		"fsm m { state .a {} }",
		"fsm m { state -a {} }",
		"fsm m { state 1a {} }",
	} {
		if _, err := parse(src); err == nil {
			t.Errorf("%q: accepted, want a syntax error", src)
		}
	}
}
