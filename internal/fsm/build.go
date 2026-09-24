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
	b.declare(b.root, tree)
	for _, ec := range tree.AllEvent() {
		b.failf(ec.GetStart(), "the machine itself has no behaviour, put %q inside a state", trigger(ec))
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

// node mirrors one declaration of the parse tree: a state, a parallel state, a
// region or a pseudostate. The builder needs the tree twice: once to declare
// every state, and again to resolve the goto targets, which may name a state
// declared further down the document.
type node struct {
	state  *model.State
	parent *node                   // nil for the root, whose state is nil too
	ctx    antlr.ParserRuleContext // the rule that declared the state
	order  []*node
	synth  map[string]*model.State // the final, terminate and history states goto asks for
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

// declare creates a state for every declaration directly inside ctx, in
// document order, and then for theirs, depth first.
func (b *builder) declare(scope *node, ctx antlr.ParserRuleContext) {
	var initial antlr.Token
	resting := false // whether scope holds a state it could start in
	for _, child := range ctx.GetChildren() {
		var (
			name antlr.TerminalNode
			kind model.StateKind
			mark antlr.TerminalNode // the initial keyword, if any
		)
		switch c := child.(type) {
		case *parser.StateContext:
			name, kind, mark = c.Identifier(), model.Normal, c.Initial()
		case *parser.ParallelContext:
			name, kind, mark = c.Identifier(), model.Parallel, c.Initial()
		case *parser.RegionContext:
			name, kind = c.Identifier(), model.Normal
		case *parser.ChoiceContext:
			name, kind = c.Identifier(), model.Choice
		case *parser.JunctionContext:
			name, kind = c.Identifier(), model.Junction
		case *parser.ForkContext:
			name, kind = c.Identifier(), model.Fork
		case *parser.JoinContext:
			name, kind = c.Identifier(), model.Join
		case *parser.PointContext:
			name, kind = c.Identifier(), model.EntryPoint
			if c.GetKind().GetText() == "exit" {
				kind = model.ExitPoint
			}
		default:
			continue // events, which walk reads
		}
		n := b.add(scope, name, kind, child.(antlr.ParserRuleContext))
		if n == nil {
			continue
		}
		resting = resting || kind == model.Normal || kind == model.Parallel
		if mark != nil {
			if initial != nil {
				b.failf(mark.GetSymbol(), "%s already starts in %q", describe(scope), initial.GetText())
			} else {
				initial = name.GetSymbol()
				if scope.state != nil {
					scope.state.Initial = n.state.Name
				} else {
					b.sm.Initial = n.state.Name
				}
			}
		}
		b.declare(n, n.ctx)
	}
	// Every region of a parallel state is entered, so it starts in none.
	if initial == nil && resting && (scope.state == nil || !scope.state.IsParallel()) {
		b.warnings.Addf("%s has no initial state", describe(scope))
	}
}

// add creates the state a declaration names, or reports that the name is
// taken and returns nil.
func (b *builder) add(scope *node, name antlr.TerminalNode, kind model.StateKind, ctx antlr.ParserRuleContext) *node {
	// One namespace for the whole machine: goto names a state outright, and
	// the model keys its states by name.
	if _, dup := b.states[name.GetText()]; dup {
		b.failf(name.GetSymbol(), "the machine already has a state called %q", name.GetText())
		return nil
	}
	st := &model.State{Name: name.GetText(), Kind: kind}
	if scope.state != nil {
		st.Parent = scope.state.Name
	}
	b.sm.States = append(b.sm.States, st)

	n := &node{state: st, parent: scope, ctx: ctx}
	b.states[st.Name] = n
	scope.order = append(scope.order, n)
	return n
}

// walk resolves the transitions and behaviours of every state under scope,
// once declare has seen the whole document.
func (b *builder) walk(scope *node) {
	for _, n := range scope.order {
		switch c := n.ctx.(type) {
		case *parser.StateContext:
			for _, ec := range c.AllEvent() {
				b.event(n, ec)
			}
		case *parser.ParallelContext:
			for _, ec := range c.AllEvent() {
				b.event(n, ec)
			}
		case *parser.ChoiceContext:
			b.branches(n, c.AllBranch())
		case *parser.JunctionContext:
			b.branches(n, c.AllBranch())
		case *parser.ForkContext:
			// A fork line is an optional actions clause followed by its goto,
			// so each actions clause belongs to the goto after it.
			var actions parser.IActionsContext
			for _, child := range c.GetChildren() {
				switch cc := child.(type) {
				case *parser.ActionsContext:
					actions = cc
				case *parser.GotoContext:
					b.leave(n, actions, cc)
					actions = nil
				}
			}
		case *parser.JoinContext:
			b.leave(n, c.Actions(), c.Goto_())
		case *parser.PointContext:
			b.leave(n, c.Actions(), c.Goto_())
		}
		b.walk(n)
	}
}

func (b *builder) event(n *node, ec parser.IEventContext) {
	// The behaviour and defer alternatives are the ones that label a name, and
	// "entry", "exit" and "do" are keywords, so no Identifier carries that text.
	if name := ec.GetName(); name != nil {
		switch {
		case ec.Defer() != nil:
			n.state.Defer = append(n.state.Defer, name.GetText())
		case name.GetText() == "entry":
			n.state.OnEntry = append(n.state.OnEntry, effects(ec.Actions())...)
		case name.GetText() == "exit":
			n.state.OnExit = append(n.state.OnExit, effects(ec.Actions())...)
		default:
			n.state.Do = append(n.state.Do, effects(ec.Actions())...)
		}
		return
	}
	t := &model.Transition{Source: n.state.Name, Actions: effects(ec.Actions())}
	// No trigger makes a completion transition, which leaves Event and After
	// empty.
	if tr := ec.Trigger(); tr != nil {
		if tr.GetName() != nil {
			t.Event = tr.GetName().GetText()
		} else {
			t.After = tr.GetDelay().GetText()
		}
	}
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

// branches adds the transitions leaving a choice or a junction. UML writes the
// catch-all branch as the guard "else", and so does the model.
func (b *builder) branches(n *node, bcs []parser.IBranchContext) {
	for _, bc := range bcs {
		t, ok := b.transition(n, bc.Actions(), bc.Goto_())
		if !ok {
			continue
		}
		if g := bc.Guard(); g != nil {
			t.Cond = text(g.Expression())
		} else if bc.GetStart().GetText() == "[" {
			t.Cond = "else"
		}
		b.sm.Transitions = append(b.sm.Transitions, t)
	}
}

// leave adds the eventless, unguarded transition leaving a fork, a join or a
// point.
func (b *builder) leave(n *node, ac parser.IActionsContext, gt parser.IGotoContext) {
	if t, ok := b.transition(n, ac, gt); ok {
		b.sm.Transitions = append(b.sm.Transitions, t)
	}
}

func (b *builder) transition(n *node, ac parser.IActionsContext, gt parser.IGotoContext) (*model.Transition, bool) {
	target, ok := b.target(n, gt)
	if !ok {
		return nil, false
	}
	return &model.Transition{Source: n.state.Name, Targets: []string{target}, Actions: effects(ac)}, true
}

// target names the state a goto leads to, creating the final, terminate or
// history state it asks for if this is the first time the document mentions
// it.
func (b *builder) target(n *node, g parser.IGotoContext) (string, bool) {
	last := g.GetStop().GetText()
	id := g.Identifier()
	if id != nil && last == id.GetText() {
		target, ok := b.states[last]
		if !ok {
			b.failf(id.GetSymbol(), "the machine has no state called %q", last)
			return "", false
		}
		return target.state.Name, true
	}
	switch last {
	case ".":
		return n.state.Name, true
	case "final":
		return b.ending(n, g, "final", model.Final)
	case "terminate":
		return b.ending(n, g, "terminate", model.Terminate)
	}
	// H or H*, of the named state or else of the declaring one.
	owner := n
	if id != nil {
		named, ok := b.states[id.GetText()]
		if !ok {
			b.failf(id.GetSymbol(), "the machine has no state called %q", id.GetText())
			return "", false
		}
		owner = named
	}
	switch {
	case owner.state.IsParallel():
		b.failf(g.GetStart(), "parallel state %q has no history of its own, only its regions do", owner.state.Name)
		return "", false
	case len(owner.order) == 0:
		b.failf(g.GetStart(), "%s has no children, so it has no history", describe(owner))
		return "", false
	case last == "H*":
		return b.synthesize(owner, "H-deep", model.HistoryDeep), true
	}
	return b.synthesize(owner, "H", model.HistoryShallow), true
}

// ending returns the final or terminate state that goto final or goto
// terminate leads to. It sits beside the transition's source, in the region
// the source belongs to, except for an exit point, which leads out of its
// state and so ends the region around that state.
func (b *builder) ending(n *node, g parser.IGotoContext, suffix string, kind model.StateKind) (string, bool) {
	scope := n.parent
	if n.state.Kind == model.ExitPoint && scope.parent != nil {
		scope = scope.parent
	}
	if scope.state != nil && scope.state.IsParallel() {
		b.failf(g.GetStart(), "parallel state %q holds only regions, so goto %s has no region to end", scope.state.Name, suffix)
		return "", false
	}
	return b.synthesize(scope, suffix, kind), true
}

// synthesize returns the final, terminate or history state of scope, which
// the language asks for by keyword instead of declaring. A name holds no dot,
// so the names made here are out of a document's reach and cannot collide
// with a declared state.
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

// trigger is what an event clause fires on, as the document writes it: an
// event name, "entry", "exit", "do", a whole "after(5s)", or the guard or
// goto of a completion transition.
func trigger(ec parser.IEventContext) string {
	if name := ec.GetName(); name != nil {
		return name.GetText()
	}
	if tr := ec.Trigger(); tr != nil {
		return text(tr)
	}
	return text(ec)
}

func effects(a parser.IActionsContext) []string {
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
