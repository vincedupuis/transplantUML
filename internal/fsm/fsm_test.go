package fsm

import (
	"fmt"
	"os"
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

// The generated parser accepts the grammar's own idioms: nested states,
// entry/exit actions, guards, action lists, and every kind of goto.
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
		"fsm m { on e goto s state s {} }",      // events must follow states
		"fsm m { state s { on e [] } }",         // empty guard
	} {
		if _, errs := parse(src); len(errs) == 0 {
			t.Errorf("%q: accepted, want a syntax error", src)
		}
	}
}
