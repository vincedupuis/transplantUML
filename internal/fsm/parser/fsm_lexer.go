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
		"", "", "", "", "Initial", "Identifier", "Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "Initial", "Identifier", "Comment", "Blank",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 24, 146, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 126, 8, 21, 10, 21, 12, 21, 129,
		9, 21, 1, 22, 1, 22, 5, 22, 133, 8, 22, 10, 22, 12, 22, 136, 9, 22, 1,
		22, 1, 22, 1, 23, 4, 23, 141, 8, 23, 11, 23, 12, 23, 142, 1, 23, 1, 23,
		0, 0, 24, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 1, 0, 4, 2, 0, 65, 90, 97,
		122, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10,
		13, 13, 32, 32, 148, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 1, 49, 1, 0, 0, 0, 3, 53, 1, 0, 0,
		0, 5, 55, 1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 63, 1, 0, 0, 0, 11, 66, 1,
		0, 0, 0, 13, 72, 1, 0, 0, 0, 15, 77, 1, 0, 0, 0, 17, 79, 1, 0, 0, 0, 19,
		81, 1, 0, 0, 0, 21, 83, 1, 0, 0, 0, 23, 85, 1, 0, 0, 0, 25, 88, 1, 0, 0,
		0, 27, 92, 1, 0, 0, 0, 29, 96, 1, 0, 0, 0, 31, 98, 1, 0, 0, 0, 33, 100,
		1, 0, 0, 0, 35, 105, 1, 0, 0, 0, 37, 107, 1, 0, 0, 0, 39, 113, 1, 0, 0,
		0, 41, 115, 1, 0, 0, 0, 43, 123, 1, 0, 0, 0, 45, 130, 1, 0, 0, 0, 47, 140,
		1, 0, 0, 0, 49, 50, 5, 102, 0, 0, 50, 51, 5, 115, 0, 0, 51, 52, 5, 109,
		0, 0, 52, 2, 1, 0, 0, 0, 53, 54, 5, 123, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56,
		5, 125, 0, 0, 56, 6, 1, 0, 0, 0, 57, 58, 5, 115, 0, 0, 58, 59, 5, 116,
		0, 0, 59, 60, 5, 97, 0, 0, 60, 61, 5, 116, 0, 0, 61, 62, 5, 101, 0, 0,
		62, 8, 1, 0, 0, 0, 63, 64, 5, 111, 0, 0, 64, 65, 5, 110, 0, 0, 65, 10,
		1, 0, 0, 0, 66, 67, 5, 101, 0, 0, 67, 68, 5, 110, 0, 0, 68, 69, 5, 116,
		0, 0, 69, 70, 5, 114, 0, 0, 70, 71, 5, 121, 0, 0, 71, 12, 1, 0, 0, 0, 72,
		73, 5, 101, 0, 0, 73, 74, 5, 120, 0, 0, 74, 75, 5, 105, 0, 0, 75, 76, 5,
		116, 0, 0, 76, 14, 1, 0, 0, 0, 77, 78, 5, 47, 0, 0, 78, 16, 1, 0, 0, 0,
		79, 80, 5, 44, 0, 0, 80, 18, 1, 0, 0, 0, 81, 82, 5, 91, 0, 0, 82, 20, 1,
		0, 0, 0, 83, 84, 5, 93, 0, 0, 84, 22, 1, 0, 0, 0, 85, 86, 5, 111, 0, 0,
		86, 87, 5, 114, 0, 0, 87, 24, 1, 0, 0, 0, 88, 89, 5, 97, 0, 0, 89, 90,
		5, 110, 0, 0, 90, 91, 5, 100, 0, 0, 91, 26, 1, 0, 0, 0, 92, 93, 5, 110,
		0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 116, 0, 0, 95, 28, 1, 0, 0, 0, 96,
		97, 5, 40, 0, 0, 97, 30, 1, 0, 0, 0, 98, 99, 5, 41, 0, 0, 99, 32, 1, 0,
		0, 0, 100, 101, 5, 103, 0, 0, 101, 102, 5, 111, 0, 0, 102, 103, 5, 116,
		0, 0, 103, 104, 5, 111, 0, 0, 104, 34, 1, 0, 0, 0, 105, 106, 5, 46, 0,
		0, 106, 36, 1, 0, 0, 0, 107, 108, 5, 102, 0, 0, 108, 109, 5, 105, 0, 0,
		109, 110, 5, 110, 0, 0, 110, 111, 5, 97, 0, 0, 111, 112, 5, 108, 0, 0,
		112, 38, 1, 0, 0, 0, 113, 114, 5, 72, 0, 0, 114, 40, 1, 0, 0, 0, 115, 116,
		5, 105, 0, 0, 116, 117, 5, 110, 0, 0, 117, 118, 5, 105, 0, 0, 118, 119,
		5, 116, 0, 0, 119, 120, 5, 105, 0, 0, 120, 121, 5, 97, 0, 0, 121, 122,
		5, 108, 0, 0, 122, 42, 1, 0, 0, 0, 123, 127, 7, 0, 0, 0, 124, 126, 7, 1,
		0, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0,
		127, 128, 1, 0, 0, 0, 128, 44, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 134,
		5, 35, 0, 0, 131, 133, 8, 2, 0, 0, 132, 131, 1, 0, 0, 0, 133, 136, 1, 0,
		0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 1, 0, 0, 0,
		136, 134, 1, 0, 0, 0, 137, 138, 6, 22, 0, 0, 138, 46, 1, 0, 0, 0, 139,
		141, 7, 3, 0, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 140,
		1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 6, 23,
		1, 0, 145, 48, 1, 0, 0, 0, 4, 0, 127, 134, 142, 2, 6, 0, 0, 0, 1, 0,
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
	fsmLexerComment    = 23
	fsmLexerBlank      = 24
)
