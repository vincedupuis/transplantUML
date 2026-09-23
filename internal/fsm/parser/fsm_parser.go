// Code generated from fsm.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // fsm
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type fsmParser struct {
	*antlr.BaseParser
}

var FsmParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fsmParserInit() {
	staticData := &FsmParserStaticData
	staticData.LiteralNames = []string{
		"", "'fsm'", "'{'", "'}'", "'state'", "'on'", "'entry'", "'exit'", "'/'",
		"','", "'['", "']'", "'or'", "'and'", "'not'", "'('", "')'", "'goto'",
		"'.'", "'final'", "'H'", "'initial'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "Initial", "Identifier", "Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"fsm", "state", "event", "actions", "identifiers", "guard", "expression",
		"or_expression", "and_expression", "not_expression", "single_expression",
		"goto",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 24, 123, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 30, 8, 0, 10, 0,
		12, 0, 33, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 3, 1, 39, 8, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 5, 1, 46, 8, 1, 10, 1, 12, 1, 49, 9, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 3, 2, 56, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 61, 8, 2, 1, 2, 1,
		2, 3, 2, 65, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 2, 3, 2, 73, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 81, 8, 4, 10, 4, 12, 4, 84,
		9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 95, 8,
		7, 10, 7, 12, 7, 98, 9, 7, 1, 8, 1, 8, 1, 8, 5, 8, 103, 8, 8, 10, 8, 12,
		8, 106, 9, 8, 1, 9, 3, 9, 109, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 118, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 0, 0, 12, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 2, 1, 0, 6, 7, 2, 0, 18, 20,
		22, 22, 126, 0, 24, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6,
		74, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10, 85, 1, 0, 0, 0, 12, 89, 1, 0, 0,
		0, 14, 91, 1, 0, 0, 0, 16, 99, 1, 0, 0, 0, 18, 108, 1, 0, 0, 0, 20, 117,
		1, 0, 0, 0, 22, 119, 1, 0, 0, 0, 24, 25, 5, 1, 0, 0, 25, 26, 5, 22, 0,
		0, 26, 31, 5, 2, 0, 0, 27, 30, 3, 2, 1, 0, 28, 30, 3, 4, 2, 0, 29, 27,
		1, 0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0,
		31, 32, 1, 0, 0, 0, 32, 34, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 5,
		3, 0, 0, 35, 36, 5, 0, 0, 1, 36, 1, 1, 0, 0, 0, 37, 39, 5, 21, 0, 0, 38,
		37, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 5, 4, 0,
		0, 41, 42, 5, 22, 0, 0, 42, 47, 5, 2, 0, 0, 43, 46, 3, 2, 1, 0, 44, 46,
		3, 4, 2, 0, 45, 43, 1, 0, 0, 0, 45, 44, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 1, 0, 0, 0, 49, 47, 1,
		0, 0, 0, 50, 51, 5, 3, 0, 0, 51, 3, 1, 0, 0, 0, 52, 53, 5, 5, 0, 0, 53,
		55, 7, 0, 0, 0, 54, 56, 3, 6, 3, 0, 55, 54, 1, 0, 0, 0, 55, 56, 1, 0, 0,
		0, 56, 73, 1, 0, 0, 0, 57, 58, 5, 5, 0, 0, 58, 60, 5, 22, 0, 0, 59, 61,
		3, 10, 5, 0, 60, 59, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 64, 3, 6, 3, 0, 63, 65, 3, 22, 11, 0, 64, 63, 1, 0, 0, 0, 64, 65, 1,
		0, 0, 0, 65, 73, 1, 0, 0, 0, 66, 67, 5, 5, 0, 0, 67, 69, 5, 22, 0, 0, 68,
		70, 3, 10, 5, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0,
		0, 0, 71, 73, 3, 22, 11, 0, 72, 52, 1, 0, 0, 0, 72, 57, 1, 0, 0, 0, 72,
		66, 1, 0, 0, 0, 73, 5, 1, 0, 0, 0, 74, 75, 5, 8, 0, 0, 75, 76, 3, 8, 4,
		0, 76, 7, 1, 0, 0, 0, 77, 82, 5, 22, 0, 0, 78, 79, 5, 9, 0, 0, 79, 81,
		5, 22, 0, 0, 80, 78, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0,
		82, 83, 1, 0, 0, 0, 83, 9, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 10,
		0, 0, 86, 87, 3, 12, 6, 0, 87, 88, 5, 11, 0, 0, 88, 11, 1, 0, 0, 0, 89,
		90, 3, 14, 7, 0, 90, 13, 1, 0, 0, 0, 91, 96, 3, 16, 8, 0, 92, 93, 5, 12,
		0, 0, 93, 95, 3, 16, 8, 0, 94, 92, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96,
		94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 15, 1, 0, 0, 0, 98, 96, 1, 0, 0,
		0, 99, 104, 3, 18, 9, 0, 100, 101, 5, 13, 0, 0, 101, 103, 3, 18, 9, 0,
		102, 100, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104,
		105, 1, 0, 0, 0, 105, 17, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 109, 5,
		14, 0, 0, 108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 1, 0, 0,
		0, 110, 111, 3, 20, 10, 0, 111, 19, 1, 0, 0, 0, 112, 118, 5, 22, 0, 0,
		113, 114, 5, 15, 0, 0, 114, 115, 3, 12, 6, 0, 115, 116, 5, 16, 0, 0, 116,
		118, 1, 0, 0, 0, 117, 112, 1, 0, 0, 0, 117, 113, 1, 0, 0, 0, 118, 21, 1,
		0, 0, 0, 119, 120, 5, 17, 0, 0, 120, 121, 7, 1, 0, 0, 121, 23, 1, 0, 0,
		0, 15, 29, 31, 38, 45, 47, 55, 60, 64, 69, 72, 82, 96, 104, 108, 117,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// fsmParserInit initializes any static state used to implement fsmParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewfsmParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FsmParserInit() {
	staticData := &FsmParserStaticData
	staticData.once.Do(fsmParserInit)
}

// NewfsmParser produces a new parser instance for the optional input antlr.TokenStream.
func NewfsmParser(input antlr.TokenStream) *fsmParser {
	FsmParserInit()
	this := new(fsmParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FsmParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "fsm.g4"

	return this
}

// fsmParser tokens.
const (
	fsmParserEOF        = antlr.TokenEOF
	fsmParserT__0       = 1
	fsmParserT__1       = 2
	fsmParserT__2       = 3
	fsmParserT__3       = 4
	fsmParserT__4       = 5
	fsmParserT__5       = 6
	fsmParserT__6       = 7
	fsmParserT__7       = 8
	fsmParserT__8       = 9
	fsmParserT__9       = 10
	fsmParserT__10      = 11
	fsmParserT__11      = 12
	fsmParserT__12      = 13
	fsmParserT__13      = 14
	fsmParserT__14      = 15
	fsmParserT__15      = 16
	fsmParserT__16      = 17
	fsmParserT__17      = 18
	fsmParserT__18      = 19
	fsmParserT__19      = 20
	fsmParserInitial    = 21
	fsmParserIdentifier = 22
	fsmParserComment    = 23
	fsmParserBlank      = 24
)

// fsmParser rules.
const (
	fsmParserRULE_fsm               = 0
	fsmParserRULE_state             = 1
	fsmParserRULE_event             = 2
	fsmParserRULE_actions           = 3
	fsmParserRULE_identifiers       = 4
	fsmParserRULE_guard             = 5
	fsmParserRULE_expression        = 6
	fsmParserRULE_or_expression     = 7
	fsmParserRULE_and_expression    = 8
	fsmParserRULE_not_expression    = 9
	fsmParserRULE_single_expression = 10
	fsmParserRULE_goto              = 11
)

// IFsmContext is an interface to support dynamic dispatch.
type IFsmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	EOF() antlr.TerminalNode
	AllState() []IStateContext
	State(i int) IStateContext
	AllEvent() []IEventContext
	Event(i int) IEventContext

	// IsFsmContext differentiates from other interfaces.
	IsFsmContext()
}

type FsmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFsmContext() *FsmContext {
	var p = new(FsmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_fsm
	return p
}

func InitEmptyFsmContext(p *FsmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_fsm
}

func (*FsmContext) IsFsmContext() {}

func NewFsmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FsmContext {
	var p = new(FsmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_fsm

	return p
}

func (s *FsmContext) GetParser() antlr.Parser { return s.parser }

func (s *FsmContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *FsmContext) EOF() antlr.TerminalNode {
	return s.GetToken(fsmParserEOF, 0)
}

func (s *FsmContext) AllState() []IStateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStateContext); ok {
			len++
		}
	}

	tst := make([]IStateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStateContext); ok {
			tst[i] = t.(IStateContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) State(i int) IStateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateContext)
}

func (s *FsmContext) AllEvent() []IEventContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEventContext); ok {
			len++
		}
	}

	tst := make([]IEventContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEventContext); ok {
			tst[i] = t.(IEventContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Event(i int) IEventContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventContext)
}

func (s *FsmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FsmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FsmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitFsm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Fsm() (localctx IFsmContext) {
	localctx = NewFsmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, fsmParserRULE_fsm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097200) != 0 {
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case fsmParserT__3, fsmParserInitial:
			{
				p.SetState(27)
				p.State()
			}

		case fsmParserT__4:
			{
				p.SetState(28)
				p.Event()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(fsmParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Match(fsmParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStateContext is an interface to support dynamic dispatch.
type IStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Initial() antlr.TerminalNode
	AllState() []IStateContext
	State(i int) IStateContext
	AllEvent() []IEventContext
	Event(i int) IEventContext

	// IsStateContext differentiates from other interfaces.
	IsStateContext()
}

type StateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateContext() *StateContext {
	var p = new(StateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_state
	return p
}

func InitEmptyStateContext(p *StateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_state
}

func (*StateContext) IsStateContext() {}

func NewStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateContext {
	var p = new(StateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_state

	return p
}

func (s *StateContext) GetParser() antlr.Parser { return s.parser }

func (s *StateContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *StateContext) Initial() antlr.TerminalNode {
	return s.GetToken(fsmParserInitial, 0)
}

func (s *StateContext) AllState() []IStateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStateContext); ok {
			len++
		}
	}

	tst := make([]IStateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStateContext); ok {
			tst[i] = t.(IStateContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) State(i int) IStateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateContext)
}

func (s *StateContext) AllEvent() []IEventContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEventContext); ok {
			len++
		}
	}

	tst := make([]IEventContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEventContext); ok {
			tst[i] = t.(IEventContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Event(i int) IEventContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventContext)
}

func (s *StateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitState(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) State() (localctx IStateContext) {
	localctx = NewStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, fsmParserRULE_state)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(37)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(40)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097200) != 0 {
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case fsmParserT__3, fsmParserInitial:
			{
				p.SetState(43)
				p.State()
			}

		case fsmParserT__4:
			{
				p.SetState(44)
				p.Event()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(fsmParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEventContext is an interface to support dynamic dispatch.
type IEventContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	Actions() IActionsContext
	Identifier() antlr.TerminalNode
	Guard() IGuardContext
	Goto_() IGotoContext

	// IsEventContext differentiates from other interfaces.
	IsEventContext()
}

type EventContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyEventContext() *EventContext {
	var p = new(EventContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_event
	return p
}

func InitEmptyEventContext(p *EventContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_event
}

func (*EventContext) IsEventContext() {}

func NewEventContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventContext {
	var p = new(EventContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_event

	return p
}

func (s *EventContext) GetParser() antlr.Parser { return s.parser }

func (s *EventContext) GetName() antlr.Token { return s.name }

func (s *EventContext) SetName(v antlr.Token) { s.name = v }

func (s *EventContext) Actions() IActionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
}

func (s *EventContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *EventContext) Guard() IGuardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardContext)
}

func (s *EventContext) Goto_() IGotoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGotoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGotoContext)
}

func (s *EventContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitEvent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Event() (localctx IEventContext) {
	localctx = NewEventContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, fsmParserRULE_event)
	var _la int

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(fsmParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*EventContext).name = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == fsmParserT__5 || _la == fsmParserT__6) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*EventContext).name = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__7 {
			{
				p.SetState(54)
				p.Actions()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(fsmParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*EventContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__9 {
			{
				p.SetState(59)
				p.Guard()
			}

		}
		{
			p.SetState(62)
			p.Actions()
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__16 {
			{
				p.SetState(63)
				p.Goto_()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Match(fsmParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*EventContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__9 {
			{
				p.SetState(68)
				p.Guard()
			}

		}
		{
			p.SetState(71)
			p.Goto_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActionsContext is an interface to support dynamic dispatch.
type IActionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifiers() IIdentifiersContext

	// IsActionsContext differentiates from other interfaces.
	IsActionsContext()
}

type ActionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionsContext() *ActionsContext {
	var p = new(ActionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_actions
	return p
}

func InitEmptyActionsContext(p *ActionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_actions
}

func (*ActionsContext) IsActionsContext() {}

func NewActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionsContext {
	var p = new(ActionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_actions

	return p
}

func (s *ActionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionsContext) Identifiers() IIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *ActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitActions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Actions() (localctx IActionsContext) {
	localctx = NewActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, fsmParserRULE_actions)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(fsmParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Identifiers()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifiersContext is an interface to support dynamic dispatch.
type IIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode

	// IsIdentifiersContext differentiates from other interfaces.
	IsIdentifiersContext()
}

type IdentifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifiersContext() *IdentifiersContext {
	var p = new(IdentifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_identifiers
	return p
}

func InitEmptyIdentifiersContext(p *IdentifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_identifiers
}

func (*IdentifiersContext) IsIdentifiersContext() {}

func NewIdentifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifiersContext {
	var p = new(IdentifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_identifiers

	return p
}

func (s *IdentifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifiersContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(fsmParserIdentifier)
}

func (s *IdentifiersContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, i)
}

func (s *IdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifiersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitIdentifiers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Identifiers() (localctx IIdentifiersContext) {
	localctx = NewIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, fsmParserRULE_identifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__8 {
		{
			p.SetState(78)
			p.Match(fsmParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardContext is an interface to support dynamic dispatch.
type IGuardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsGuardContext differentiates from other interfaces.
	IsGuardContext()
}

type GuardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardContext() *GuardContext {
	var p = new(GuardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_guard
	return p
}

func InitEmptyGuardContext(p *GuardContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_guard
}

func (*GuardContext) IsGuardContext() {}

func NewGuardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardContext {
	var p = new(GuardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_guard

	return p
}

func (s *GuardContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitGuard(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Guard() (localctx IGuardContext) {
	localctx = NewGuardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, fsmParserRULE_guard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(fsmParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Expression()
	}
	{
		p.SetState(87)
		p.Match(fsmParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Or_expression() IOr_expressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Or_expression() IOr_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOr_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOr_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, fsmParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Or_expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOr_expressionContext is an interface to support dynamic dispatch.
type IOr_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnd_expression() []IAnd_expressionContext
	And_expression(i int) IAnd_expressionContext

	// IsOr_expressionContext differentiates from other interfaces.
	IsOr_expressionContext()
}

type Or_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_expressionContext() *Or_expressionContext {
	var p = new(Or_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_or_expression
	return p
}

func InitEmptyOr_expressionContext(p *Or_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_or_expression
}

func (*Or_expressionContext) IsOr_expressionContext() {}

func NewOr_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_expressionContext {
	var p = new(Or_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_or_expression

	return p
}

func (s *Or_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_expressionContext) AllAnd_expression() []IAnd_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnd_expressionContext); ok {
			len++
		}
	}

	tst := make([]IAnd_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnd_expressionContext); ok {
			tst[i] = t.(IAnd_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Or_expressionContext) And_expression(i int) IAnd_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnd_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnd_expressionContext)
}

func (s *Or_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitOr_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Or_expression() (localctx IOr_expressionContext) {
	localctx = NewOr_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, fsmParserRULE_or_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.And_expression()
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__11 {
		{
			p.SetState(92)
			p.Match(fsmParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.And_expression()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnd_expressionContext is an interface to support dynamic dispatch.
type IAnd_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNot_expression() []INot_expressionContext
	Not_expression(i int) INot_expressionContext

	// IsAnd_expressionContext differentiates from other interfaces.
	IsAnd_expressionContext()
}

type And_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_expressionContext() *And_expressionContext {
	var p = new(And_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_and_expression
	return p
}

func InitEmptyAnd_expressionContext(p *And_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_and_expression
}

func (*And_expressionContext) IsAnd_expressionContext() {}

func NewAnd_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_expressionContext {
	var p = new(And_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_and_expression

	return p
}

func (s *And_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *And_expressionContext) AllNot_expression() []INot_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INot_expressionContext); ok {
			len++
		}
	}

	tst := make([]INot_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INot_expressionContext); ok {
			tst[i] = t.(INot_expressionContext)
			i++
		}
	}

	return tst
}

func (s *And_expressionContext) Not_expression(i int) INot_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INot_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INot_expressionContext)
}

func (s *And_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitAnd_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) And_expression() (localctx IAnd_expressionContext) {
	localctx = NewAnd_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, fsmParserRULE_and_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Not_expression()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__12 {
		{
			p.SetState(100)
			p.Match(fsmParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Not_expression()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INot_expressionContext is an interface to support dynamic dispatch.
type INot_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Single_expression() ISingle_expressionContext

	// IsNot_expressionContext differentiates from other interfaces.
	IsNot_expressionContext()
}

type Not_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_expressionContext() *Not_expressionContext {
	var p = new(Not_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_not_expression
	return p
}

func InitEmptyNot_expressionContext(p *Not_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_not_expression
}

func (*Not_expressionContext) IsNot_expressionContext() {}

func NewNot_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_expressionContext {
	var p = new(Not_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_not_expression

	return p
}

func (s *Not_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Not_expressionContext) Single_expression() ISingle_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingle_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingle_expressionContext)
}

func (s *Not_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitNot_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Not_expression() (localctx INot_expressionContext) {
	localctx = NewNot_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, fsmParserRULE_not_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__13 {
		{
			p.SetState(107)
			p.Match(fsmParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(110)
		p.Single_expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingle_expressionContext is an interface to support dynamic dispatch.
type ISingle_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Expression() IExpressionContext

	// IsSingle_expressionContext differentiates from other interfaces.
	IsSingle_expressionContext()
}

type Single_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_expressionContext() *Single_expressionContext {
	var p = new(Single_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_single_expression
	return p
}

func InitEmptySingle_expressionContext(p *Single_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_single_expression
}

func (*Single_expressionContext) IsSingle_expressionContext() {}

func NewSingle_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_expressionContext {
	var p = new(Single_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_single_expression

	return p
}

func (s *Single_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_expressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *Single_expressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Single_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitSingle_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Single_expression() (localctx ISingle_expressionContext) {
	localctx = NewSingle_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, fsmParserRULE_single_expression)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.Match(fsmParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Expression()
		}
		{
			p.SetState(115)
			p.Match(fsmParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGotoContext is an interface to support dynamic dispatch.
type IGotoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsGotoContext differentiates from other interfaces.
	IsGotoContext()
}

type GotoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotoContext() *GotoContext {
	var p = new(GotoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_goto
	return p
}

func InitEmptyGotoContext(p *GotoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_goto
}

func (*GotoContext) IsGotoContext() {}

func NewGotoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotoContext {
	var p = new(GotoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_goto

	return p
}

func (s *GotoContext) GetParser() antlr.Parser { return s.parser }

func (s *GotoContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *GotoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitGoto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Goto_() (localctx IGotoContext) {
	localctx = NewGotoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, fsmParserRULE_goto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(fsmParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6029312) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
