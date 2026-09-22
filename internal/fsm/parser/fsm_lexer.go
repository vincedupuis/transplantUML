// Code generated from fsm.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type fsmLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FsmLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fsmlexerLexerInit() {
	staticData := &FsmLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'fsm'", "'{'", "'}'", "'state'", "'on'", "'entry'", "'exit'", "'/'",
		"','", "'['", "']'", "'or'", "'and'", "'not'", "'('", "')'", "'goto'",
		"'.'", "'final'", "'H'", "'initial'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "Initial", "Identifier", "Prefix", "Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "Initial", "Identifier", "Prefix", "Comment",
		"Blank",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 159, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 128, 8, 21, 10,
		21, 12, 21, 131, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 4, 22, 138,
		8, 22, 11, 22, 12, 22, 139, 3, 22, 142, 8, 22, 1, 23, 1, 23, 5, 23, 146,
		8, 23, 10, 23, 12, 23, 149, 9, 23, 1, 23, 1, 23, 1, 24, 4, 24, 154, 8,
		24, 11, 24, 12, 24, 155, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 1, 0, 4, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97,
		122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 163, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 55, 1, 0, 0, 0, 5, 57, 1, 0,
		0, 0, 7, 59, 1, 0, 0, 0, 9, 65, 1, 0, 0, 0, 11, 68, 1, 0, 0, 0, 13, 74,
		1, 0, 0, 0, 15, 79, 1, 0, 0, 0, 17, 81, 1, 0, 0, 0, 19, 83, 1, 0, 0, 0,
		21, 85, 1, 0, 0, 0, 23, 87, 1, 0, 0, 0, 25, 90, 1, 0, 0, 0, 27, 94, 1,
		0, 0, 0, 29, 98, 1, 0, 0, 0, 31, 100, 1, 0, 0, 0, 33, 102, 1, 0, 0, 0,
		35, 107, 1, 0, 0, 0, 37, 109, 1, 0, 0, 0, 39, 115, 1, 0, 0, 0, 41, 117,
		1, 0, 0, 0, 43, 125, 1, 0, 0, 0, 45, 141, 1, 0, 0, 0, 47, 143, 1, 0, 0,
		0, 49, 153, 1, 0, 0, 0, 51, 52, 5, 102, 0, 0, 52, 53, 5, 115, 0, 0, 53,
		54, 5, 109, 0, 0, 54, 2, 1, 0, 0, 0, 55, 56, 5, 123, 0, 0, 56, 4, 1, 0,
		0, 0, 57, 58, 5, 125, 0, 0, 58, 6, 1, 0, 0, 0, 59, 60, 5, 115, 0, 0, 60,
		61, 5, 116, 0, 0, 61, 62, 5, 97, 0, 0, 62, 63, 5, 116, 0, 0, 63, 64, 5,
		101, 0, 0, 64, 8, 1, 0, 0, 0, 65, 66, 5, 111, 0, 0, 66, 67, 5, 110, 0,
		0, 67, 10, 1, 0, 0, 0, 68, 69, 5, 101, 0, 0, 69, 70, 5, 110, 0, 0, 70,
		71, 5, 116, 0, 0, 71, 72, 5, 114, 0, 0, 72, 73, 5, 121, 0, 0, 73, 12, 1,
		0, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76, 5, 120, 0, 0, 76, 77, 5, 105, 0,
		0, 77, 78, 5, 116, 0, 0, 78, 14, 1, 0, 0, 0, 79, 80, 5, 47, 0, 0, 80, 16,
		1, 0, 0, 0, 81, 82, 5, 44, 0, 0, 82, 18, 1, 0, 0, 0, 83, 84, 5, 91, 0,
		0, 84, 20, 1, 0, 0, 0, 85, 86, 5, 93, 0, 0, 86, 22, 1, 0, 0, 0, 87, 88,
		5, 111, 0, 0, 88, 89, 5, 114, 0, 0, 89, 24, 1, 0, 0, 0, 90, 91, 5, 97,
		0, 0, 91, 92, 5, 110, 0, 0, 92, 93, 5, 100, 0, 0, 93, 26, 1, 0, 0, 0, 94,
		95, 5, 110, 0, 0, 95, 96, 5, 111, 0, 0, 96, 97, 5, 116, 0, 0, 97, 28, 1,
		0, 0, 0, 98, 99, 5, 40, 0, 0, 99, 30, 1, 0, 0, 0, 100, 101, 5, 41, 0, 0,
		101, 32, 1, 0, 0, 0, 102, 103, 5, 103, 0, 0, 103, 104, 5, 111, 0, 0, 104,
		105, 5, 116, 0, 0, 105, 106, 5, 111, 0, 0, 106, 34, 1, 0, 0, 0, 107, 108,
		5, 46, 0, 0, 108, 36, 1, 0, 0, 0, 109, 110, 5, 102, 0, 0, 110, 111, 5,
		105, 0, 0, 111, 112, 5, 110, 0, 0, 112, 113, 5, 97, 0, 0, 113, 114, 5,
		108, 0, 0, 114, 38, 1, 0, 0, 0, 115, 116, 5, 72, 0, 0, 116, 40, 1, 0, 0,
		0, 117, 118, 5, 105, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5, 105, 0,
		0, 120, 121, 5, 116, 0, 0, 121, 122, 5, 105, 0, 0, 122, 123, 5, 97, 0,
		0, 123, 124, 5, 108, 0, 0, 124, 42, 1, 0, 0, 0, 125, 129, 7, 0, 0, 0, 126,
		128, 7, 1, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 44, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 132, 133, 5, 46, 0, 0, 133, 142, 5, 47, 0, 0, 134, 135, 5, 46, 0,
		0, 135, 136, 5, 46, 0, 0, 136, 138, 5, 47, 0, 0, 137, 134, 1, 0, 0, 0,
		138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		142, 1, 0, 0, 0, 141, 132, 1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 142, 46, 1,
		0, 0, 0, 143, 147, 5, 35, 0, 0, 144, 146, 8, 2, 0, 0, 145, 144, 1, 0, 0,
		0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 151, 6, 23, 0, 0, 151, 48,
		1, 0, 0, 0, 152, 154, 7, 3, 0, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0,
		157, 158, 6, 24, 1, 0, 158, 50, 1, 0, 0, 0, 6, 0, 129, 139, 141, 147, 155,
		2, 6, 0, 0, 0, 1, 0,
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

// fsmLexerInit initializes any static state used to implement fsmLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewfsmLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FsmLexerInit() {
	staticData := &FsmLexerLexerStaticData
	staticData.once.Do(fsmlexerLexerInit)
}

// NewfsmLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewfsmLexer(input antlr.CharStream) *fsmLexer {
	FsmLexerInit()
	l := new(fsmLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FsmLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "fsm.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// fsmLexer tokens.
const (
	fsmLexerT__0       = 1
	fsmLexerT__1       = 2
	fsmLexerT__2       = 3
	fsmLexerT__3       = 4
	fsmLexerT__4       = 5
	fsmLexerT__5       = 6
	fsmLexerT__6       = 7
	fsmLexerT__7       = 8
	fsmLexerT__8       = 9
	fsmLexerT__9       = 10
	fsmLexerT__10      = 11
	fsmLexerT__11      = 12
	fsmLexerT__12      = 13
	fsmLexerT__13      = 14
	fsmLexerT__14      = 15
	fsmLexerT__15      = 16
	fsmLexerT__16      = 17
	fsmLexerT__17      = 18
	fsmLexerT__18      = 19
	fsmLexerT__19      = 20
	fsmLexerInitial    = 21
	fsmLexerIdentifier = 22
	fsmLexerPrefix     = 23
	fsmLexerComment    = 24
	fsmLexerBlank      = 25
)
