package fsm

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/vincedupuis/transplantUML/internal/fsm/parser"
	"github.com/vincedupuis/transplantUML/internal/model"
)

// Parser reads the fsm language into the model.
type Parser struct{}

func (Parser) Parse(src []byte) (*model.StateMachine, model.Warnings, error) {
	tree, err := parse(string(src))
	if err != nil {
		return nil, nil, err
	}
	b := &builder{
		sm:     &model.StateMachine{Name: tree.Identifier().GetText()},
		root:   &node{},
		states: map[string]*node{},
	}
	b.declare(b.root, tree.AllState())
	for _, ec := range tree.AllEvent() {
		b.failf(ec.GetStart(), "the machine itself has no behaviour, put %q inside a state", ec.GetName().GetText())
	}
	b.walk(b.root)
	if err := errors.Join(b.errs...); err != nil {
		return nil, nil, err
	}
	return b.sm, b.warnings, nil
}

// syntaxErrors collects what the generated lexer and parser reject, instead of
// printing them to stderr as ANTLR's default listener does.
type syntaxErrors struct {
	*antlr.DefaultErrorListener
	errs []error
}

func (l *syntaxErrors) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	l.errs = append(l.errs, fmt.Errorf("%d:%d: %s", line, column, msg))
}

func parse(src string) (parser.IFsmContext, error) {
	errs := &syntaxErrors{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lexer := parser.NewfsmLexer(antlr.NewInputStream(src))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)
	p := parser.NewfsmParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)
	tree := p.Fsm()
	return tree, errors.Join(errs.errs...)
}

// node mirrors one state of the parse tree. The builder needs the tree twice:
// once to declare every state, and again to resolve the goto targets, which may
// name a state declared further down the document.
type node struct {
	state  *model.State
	parent *node // nil for the root, whose state is nil too
	ctx    parser.IStateContext
	order  []*node
	synth  map[string]*model.State // the final and history states goto asks for
}

type builder struct {
	sm       *model.StateMachine
	root     *node
	states   map[string]*node // every declared state, by name
	warnings model.Warnings
	errs     []error
}

func (b *builder) failf(tok antlr.Token, format string, args ...any) {
	b.errs = append(b.errs, fmt.Errorf("%d:%d: %s", tok.GetLine(), tok.GetColumn(), fmt.Sprintf(format, args...)))
}

// declare creates a state for every child of scope, depth first.
func (b *builder) declare(scope *node, states []parser.IStateContext) {
	var initial antlr.Token
	for _, sc := range states {
		name := sc.Identifier().GetText()
		// One namespace for the whole machine: goto names a state outright,
		// and the model keys its states by name.
		if _, dup := b.states[name]; dup {
			b.failf(sc.Identifier().GetSymbol(), "the machine already has a state called %q", name)
			continue
		}
		st := &model.State{Name: name, Kind: model.Normal}
		if scope.state != nil {
			st.Parent = scope.state.Name
		}
		b.sm.States = append(b.sm.States, st)

		n := &node{state: st, parent: scope, ctx: sc}
		b.states[name] = n
		scope.order = append(scope.order, n)

		if tok := sc.Initial(); tok != nil {
			if initial != nil {
				b.failf(tok.GetSymbol(), "%s already starts in %q", describe(scope), initial.GetText())
			} else {
				initial = sc.Identifier().GetSymbol()
				if scope.state != nil {
					scope.state.Initial = name
				} else {
					b.sm.Initial = name
				}
			}
		}
		b.declare(n, sc.AllState())
	}
	if initial == nil && len(scope.order) > 0 {
		b.warnings.Addf("%s has no initial state", describe(scope))
	}
}

// walk resolves the transitions and behaviours of every state under scope,
// once declare has seen the whole document.
func (b *builder) walk(scope *node) {
	for _, n := range scope.order {
		for _, ec := range n.ctx.AllEvent() {
			b.event(n, ec)
		}
		b.walk(n)
	}
}

func (b *builder) event(n *node, ec parser.IEventContext) {
	// "entry" and "exit" are keywords, so no Identifier can carry that text.
	switch name := ec.GetName().GetText(); name {
	case "entry":
		n.state.OnEntry = append(n.state.OnEntry, effects(ec)...)
		return
	case "exit":
		n.state.OnExit = append(n.state.OnExit, effects(ec)...)
		return
	}
	t := &model.Transition{Source: n.state.Name, Event: ec.GetName().GetText(), Actions: effects(ec)}
	if g := ec.Guard(); g != nil {
		t.Cond = text(g.Expression())
	}
	if gt := ec.Goto_(); gt != nil {
		target, ok := b.target(n, gt)
		if !ok {
			return
		}
		t.Targets = []string{target}
	}
	b.sm.Transitions = append(b.sm.Transitions, t)
}

// target names the state a goto leads to, creating the final or history state
// it asks for if this is the first time the document mentions it.
func (b *builder) target(n *node, g parser.IGotoContext) (string, bool) {
	if id := g.Identifier(); id != nil {
		target, ok := b.states[id.GetText()]
		if !ok {
			b.failf(id.GetSymbol(), "the machine has no state called %q", id.GetText())
			return "", false
		}
		return target.state.Name, true
	}
	switch g.GetChild(1).(antlr.TerminalNode).GetText() {
	case ".":
		return n.state.Name, true
	case "final":
		// A final state completes the region holding the transition's source,
		// so it is the source's sibling, not its child.
		return b.synthesize(n.parent, "final", model.Final), true
	default: // H
		if len(n.order) == 0 {
			b.failf(g.GetStart(), "state %q has no children, so it has no history", n.state.Name)
			return "", false
		}
		return b.synthesize(n, "H", model.HistoryShallow), true
	}
}

// synthesize returns the final or history state of scope, which the language
// asks for by keyword instead of declaring. The "." in the name cannot clash
// with a declared one, since Identifier has no punctuation.
func (b *builder) synthesize(scope *node, suffix string, kind model.StateKind) string {
	if s, ok := scope.synth[suffix]; ok {
		return s.Name
	}
	s := &model.State{Name: suffix, Kind: kind}
	if scope.state != nil {
		s.Name = scope.state.Name + "." + suffix
		s.Parent = scope.state.Name
	}
	if scope.synth == nil {
		scope.synth = map[string]*model.State{}
	}
	scope.synth[suffix] = s
	b.sm.States = append(b.sm.States, s)
	return s.Name
}

func effects(ec parser.IEventContext) []string {
	a := ec.Actions()
	if a == nil {
		return nil
	}
	ids := a.Identifiers().AllIdentifier()
	out := make([]string, 0, len(ids))
	for _, id := range ids {
		out = append(out, id.GetText())
	}
	return out
}

func describe(n *node) string {
	if n.state == nil {
		return "the machine"
	}
	return fmt.Sprintf("state %q", n.state.Name)
}

// text returns a rule's source text. GetText concatenates the tokens with no
// separator, which would turn "not a and b" into "notaandb".
func text(ctx antlr.ParserRuleContext) string {
	start, stop := ctx.GetStart(), ctx.GetStop()
	return start.GetInputStream().GetText(start.GetStart(), stop.GetStop())
}
