package fsm

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/vincedupuis/transplantUML/internal/fsm/parser"
)

// syntaxErrors collects what the generated lexer and parser reject, instead of
// printing it to stderr as ANTLR's default listener does.
type syntaxErrors struct {
	*antlr.DefaultErrorListener
	errs []string
}

func (l *syntaxErrors) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	l.errs = append(l.errs, fmt.Sprintf("%d:%d: %s", line, column, msg))
}

func parse(src string) (parser.IFsmContext, []string) {
	errs := &syntaxErrors{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lexer := parser.NewfsmLexer(antlr.NewInputStream(src))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)
	p := parser.NewfsmParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)
	return p.Fsm(), errs.errs
}

// The generated parser accepts the grammar's own idioms: nested states
// interleaved with the events that concern them, entry/exit actions, guards,
// action lists, and every kind of goto.
func TestGrammarAccepts(t *testing.T) {
	src, err := os.ReadFile("testdata/coffee.fsm")
	if err != nil {
		t.Fatal(err)
	}
	tree, errs := parse(string(src))
	if len(errs) > 0 {
		t.Fatalf("syntax errors: %v", errs)
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
		"fsm m { state s { on e goto a/ } }",    // path ending in a separator
		"fsm m { initial }",                     // initial modifies a state
		"fsm m { state s { initial on e / a } }",
	} {
		if _, errs := parse(src); len(errs) == 0 {
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
		if _, errs := parse("fsm m { state s { " + body + " } state t {} }"); len(errs) > 0 {
			t.Errorf("%q: %v", body, errs)
		}
	}
}

// Every goto target form lexes: '/' opens an absolute path even though the
// same token opens an action list, and Prefix carries the relative forms.
func TestGotoTargets(t *testing.T) {
	for _, target := range []string{
		".",
		"final",
		"H",
		"t",
		"a/b/c",
		"/a/b",
		"./b",
		"../b",
		"../../a/b",
	} {
		if _, errs := parse("fsm m { state s { on e goto " + target + " } }"); len(errs) > 0 {
			t.Errorf("goto %s: %v", target, errs)
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
		if _, errs := parse(string(src)); len(errs) > 0 {
			t.Errorf("%s: %v", path, errs)
		}
	}
}

// "initial" marks the child a compound state starts in, and the machine's own
// starting state at the top level. The label gives the visitor the token.
func TestInitialState(t *testing.T) {
	tree, errs := parse("fsm m { initial state a { initial state b {} state c {} } state d {} }")
	if len(errs) > 0 {
		t.Fatalf("syntax errors: %v", errs)
	}
	top := tree.AllState()
	if top[0].GetInit() == nil {
		t.Error("a: not marked initial")
	}
	if top[1].GetInit() != nil {
		t.Error("d: marked initial")
	}
	inner := top[0].AllState()
	if inner[0].GetInit() == nil {
		t.Error("b: not marked initial")
	}
	if inner[1].GetInit() != nil {
		t.Error("c: marked initial")
	}
}
