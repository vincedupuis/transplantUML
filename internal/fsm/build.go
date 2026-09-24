package fsm

import (
	"errors"
	"fmt"
	"slices"
	"strings"

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
		sm:     &model.StateMachine{Name: tree.Identifier().GetText(), Note: note(tree.Note())},
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
// submachine state, a region or a pseudostate. The builder needs the tree
// twice: once to declare every state, and again to resolve the goto targets,
// which may name a state declared further down the document.
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
	var (
		initial antlr.Token
		starts  []*parser.StartContext
	)
	resting := false // whether scope holds a state it could start in
	for _, child := range ctx.GetChildren() {
		var (
			name antlr.TerminalNode
			kind model.StateKind
			mark antlr.TerminalNode // the initial keyword, if any
		)
		switch c := child.(type) {
		case *parser.StartContext:
			starts = append(starts, c) // its goto may name a child declared further down
			continue
		case *parser.StateContext:
			name, kind, mark = c.Identifier(), model.Normal, c.Initial()
		case *parser.ParallelContext:
			name, kind, mark = c.Identifier(), model.Parallel, c.Initial()
		case *parser.SubmachineContext:
			name, kind, mark = c.Identifier(), model.Normal, c.Initial()
		case *parser.RegionContext:
			name, kind = c.Identifier(), model.Normal
		case *parser.ChoiceContext:
			name, kind, mark = c.Identifier(), model.Choice, c.Initial()
		case *parser.JunctionContext:
			name, kind, mark = c.Identifier(), model.Junction, c.Initial()
		case *parser.ForkContext:
			name, kind = c.Identifier(), model.Fork
		case *parser.JoinContext:
			name, kind = c.Identifier(), model.Join
		case *parser.PointContext:
			name, kind = c.Identifier(), model.EntryPoint
			if c.GetKind().GetText() == "exit" {
				kind = model.ExitPoint
			}
		case *parser.ReferenceContext:
			name, kind = c.Identifier(), model.EntryPoint
		case *parser.FinalContext:
			name, kind = c.Identifier(), model.Final
		case *parser.TerminateContext:
			name, kind = c.Identifier(), model.Terminate
		case *parser.HistoryContext:
			suffix, kind := "H", model.HistoryShallow
			if c.GetKind().GetText() == "H*" {
				suffix, kind = "H-deep", model.HistoryDeep
			}
			b.unnamed(scope, c, c.GetKind().GetText(), suffix, kind)
			continue
		default:
			continue // events, which walk reads
		}
		// A final or terminate state without a name is the one goto final or
		// goto terminate reaches.
		if name == nil {
			b.unnamed(scope, child.(antlr.ParserRuleContext), "an unnamed "+string(kind)+" state", string(kind), kind)
			continue
		}
		n := b.add(scope, name, kind, child.(antlr.ParserRuleContext))
		if n == nil {
			continue
		}
		// The state is named after the machine it refers to.
		if _, ok := child.(*parser.SubmachineContext); ok {
			n.state.Submachine = n.state.Name
		}
		n.state.Note = note(leading(n.ctx))
		// Every declaration may carry a stereotype after its name.
		if st := child.(stereotyped).Stereotype(); st != nil {
			n.state.Stereotype = st.Identifier().GetText()
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
	for _, c := range starts {
		if initial != nil {
			b.failf(c.Initial().GetSymbol(), "%s already starts in %q", describe(scope), initial.GetText())
			continue
		}
		b.start(scope, c)
		initial = c.Identifier().GetSymbol()
	}
	// Every region of a parallel state is entered, so it starts in none.
	if initial == nil && resting && (scope.state == nil || !scope.state.IsParallel()) {
		b.warnings.Addf("%s has no initial state", describe(scope))
	}
}

// start resolves an initial line: UML's initial pseudostate, whose one
// transition leads to the state scope starts in and may carry an effect.
func (b *builder) start(scope *node, c *parser.StartContext) {
	id := c.Identifier()
	target, ok := b.states[id.GetText()]
	if !ok || target.parent != scope {
		b.failf(id.GetSymbol(), "%s has no child called %q to start in", describe(scope), id.GetText())
		return
	}
	if scope.state != nil {
		scope.state.Initial, scope.state.InitialActions = target.state.Name, effects(c.Actions())
	} else {
		b.sm.Initial, b.sm.InitialActions = target.state.Name, effects(c.Actions())
	}
}

// unnamed declares the final, terminate or history state of scope that the
// language otherwise creates on a goto, so that it can carry a note and a
// stereotype. A second declaration would fold into the same state.
func (b *builder) unnamed(scope *node, ctx antlr.ParserRuleContext, what, suffix string, kind model.StateKind) {
	at := ctx.GetStart()
	if leading(ctx) != nil {
		at = ctx.GetChild(1).(antlr.TerminalNode).GetSymbol()
	}
	if _, ok := scope.synth[suffix]; ok {
		b.failf(at, "%s declares %s twice", describe(scope), what)
		return
	}
	b.synthesize(scope, suffix, kind)
	s := scope.synth[suffix]
	s.Note = note(leading(ctx))
	if st := ctx.(stereotyped).Stereotype(); st != nil {
		s.Stereotype = st.Identifier().GetText()
	}
}

// stereotyped is every declaration: each may carry a stereotype.
type stereotyped interface {
	Stereotype() parser.IStereotypeContext
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
			for _, hc := range c.AllHistory() {
				b.history(n, hc)
			}
			for _, ec := range c.AllEvent() {
				b.event(n, ec)
			}
		case *parser.RegionContext:
			for _, hc := range c.AllHistory() {
				b.history(n, hc)
			}
		case *parser.ParallelContext:
			for _, ec := range c.AllEvent() {
				b.event(n, ec)
			}
		case *parser.SubmachineContext:
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
			// A note before a line belongs to that line's transition; the
			// one before the declaration is the fork's own.
			var (
				actions parser.IActionsContext
				line    antlr.TerminalNode
			)
			for _, child := range c.GetChildren()[1:] {
				switch cc := child.(type) {
				case antlr.TerminalNode:
					if isNote(cc) {
						line = cc
					}
				case *parser.ActionsContext:
					actions = cc
				case *parser.GotoContext:
					if t, ok := b.transition(n, actions, cc); ok {
						t.Note = note(line)
						b.sm.Transitions = append(b.sm.Transitions, t)
					}
					actions, line = nil, nil
				}
			}
		case *parser.JoinContext:
			b.leave(n, c.Actions(), c.Goto_())
		case *parser.PointContext:
			if n.state.Kind == model.EntryPoint && n.parent.state != nil && n.parent.state.Submachine != "" {
				b.failf(c.GetKind(), "entry point %q leads into the machine %q refers to, so it takes no goto", n.state.Name, n.parent.state.Name)
				break
			}
			b.leave(n, c.Actions(), c.Goto_())
		}
		b.walk(n)
	}
}

func (b *builder) event(n *node, ec parser.IEventContext) {
	about := note(ec.Note())
	// The model holds one condition per state, so a second would be lost.
	if inv := ec.Invariant(); inv != nil {
		if n.state.Invariant != "" {
			b.failf(inv.GetSymbol(), "state %q already has the invariant [%s], join the conditions with and", n.state.Name, n.state.Invariant)
			return
		}
		n.state.Invariant = text(ec.Guard().Expression())
		n.state.InvariantNote = about
		return
	}
	// The behaviour and defer alternatives are the ones that label a name, and
	// "entry", "exit" and "do" are keywords, so no Identifier carries that text.
	if name := ec.GetName(); name != nil {
		// UML gives a state one entry, one exit and one do behaviour, which
		// the clauses build up together, so their notes add up too.
		switch {
		case ec.Defer() != nil:
			n.state.Defer = append(n.state.Defer, name.GetText())
			if about != "" {
				if n.state.DeferNotes == nil {
					n.state.DeferNotes = map[string]string{}
				}
				n.state.DeferNotes[name.GetText()] = addLines(n.state.DeferNotes[name.GetText()], about)
			}
		case name.GetText() == "entry":
			n.state.OnEntry = append(n.state.OnEntry, effects(ec.Actions())...)
			n.state.EntryNote = addLines(n.state.EntryNote, about)
		case name.GetText() == "exit":
			n.state.OnExit = append(n.state.OnExit, effects(ec.Actions())...)
			n.state.ExitNote = addLines(n.state.ExitNote, about)
		default:
			n.state.Do = append(n.state.Do, effects(ec.Actions())...)
			n.state.DoNote = addLines(n.state.DoNote, about)
		}
		return
	}
	// A clause with no goto is UML's internal transition, the one a state
	// lists in its compartment: no state change, no exit or entry.
	t := &model.Transition{Source: n.state.Name, Actions: effects(ec.Actions()), Kind: model.Internal, Note: about}
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
		target, kind, ok := b.destination(n, gt)
		if !ok {
			return
		}
		t.Targets, t.Kind = []string{target}, kind
	}
	b.sm.Transitions = append(b.sm.Transitions, t)
}

// history adds the default transition an H or H* line may give the history
// state declare made for it, the one taken while the state has no history
// yet.
func (b *builder) history(n *node, hc parser.IHistoryContext) {
	kind := hc.GetKind()
	if len(n.order) == 0 {
		b.failf(kind, "%s has no children, so it has no history", describe(n))
		return
	}
	if hc.Goto_() == nil {
		return
	}
	suffix := "H"
	if kind.GetText() == "H*" {
		suffix = "H-deep"
	}
	// The history sits inside n, so goto final ends n's own region.
	b.leave(&node{state: n.synth[suffix], parent: n}, hc.Actions(), hc.Goto_())
}

// branches adds the transitions leaving a choice or a junction. UML writes the
// catch-all branch as the guard "else", and so does the model.
func (b *builder) branches(n *node, bcs []parser.IBranchContext) {
	for _, bc := range bcs {
		t, ok := b.transition(n, bc.Actions(), bc.Goto_())
		if !ok {
			continue
		}
		t.Note = note(bc.Note())
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
	target, kind, ok := b.destination(n, gt)
	if !ok {
		return nil, false
	}
	return &model.Transition{Source: n.state.Name, Targets: []string{target}, Actions: effects(ac), Kind: kind}, true
}

// destination resolves a goto to its target and the transition's kind:
// external, UML's default, or local when the goto says so. UML allows a local
// transition only into its own source, which it then never leaves.
func (b *builder) destination(n *node, g parser.IGotoContext) (string, model.TransitionKind, bool) {
	target, ok := b.target(n, g)
	if !ok || g.Local() == nil {
		return target, "", ok
	}
	if !slices.Contains(b.sm.Ancestors(target), n.state.Name) {
		b.failf(g.Local().GetSymbol(), "a local transition stays inside %q, but %q is not inside it", n.state.Name, target)
		return "", "", false
	}
	return target, model.Local, true
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
	case owner.state.Submachine != "":
		b.failf(g.GetStart(), "submachine state %q has no history of its own, the machine it refers to may", owner.state.Name)
		return "", false
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
// event name, "entry", "exit", "do", a whole "after(5s)", the guard or goto
// of a completion transition, or a whole invariant clause, but never the note
// before it.
func trigger(ec parser.IEventContext) string {
	if name := ec.GetName(); name != nil {
		return name.GetText()
	}
	if tr := ec.Trigger(); tr != nil {
		return text(tr)
	}
	if nt := ec.Note(); nt != nil {
		return strings.TrimSpace(strings.TrimPrefix(text(ec), nt.GetText()))
	}
	return text(ec)
}

// leading returns the note written before a declaration, or nil.
func leading(ctx antlr.ParserRuleContext) antlr.TerminalNode {
	if t, ok := ctx.GetChild(0).(antlr.TerminalNode); ok && isNote(t) {
		return t
	}
	return nil
}

// isNote tells a note from the other tokens, none of which starts with a bar.
func isNote(t antlr.TerminalNode) bool {
	return strings.HasPrefix(t.GetText(), "|")
}

// note returns the text between a note's bars. A backslash escapes the
// character after it, and each line loses the indentation that lines it up
// with the document, as do the blank lines around the text.
func note(t antlr.TerminalNode) string {
	if t == nil {
		return ""
	}
	raw := t.GetText()
	raw = raw[1 : len(raw)-1]
	var text strings.Builder
	for i := 0; i < len(raw); i++ {
		if raw[i] == '\\' && i+1 < len(raw) {
			i++
		}
		text.WriteByte(raw[i])
	}
	lines := strings.Split(text.String(), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Trim(strings.Join(lines, "\n"), "\n")
}

// addLines appends more lines to a note, if there are any.
func addLines(note, more string) string {
	if note == "" || more == "" {
		return note + more
	}
	return note + "\n" + more
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
