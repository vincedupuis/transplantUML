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
		"", "'fsm'", "'{'", "'}'", "'state'", "'entry'", "'exit'", "'do'", "'on'",
		"'/'", "'after'", "'('", "')'", "','", "'['", "']'", "'or'", "'and'",
		"'not'", "'goto'", "'.'", "'final'", "'H'", "'initial'", "'defer'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "Initial", "Defer", "Duration", "Identifier",
		"Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "Initial", "Defer", "Duration",
		"Identifier", "Comment", "Blank",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 187, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 4, 24, 148, 8, 24, 11, 24, 12, 24, 149,
		1, 24, 1, 24, 4, 24, 154, 8, 24, 11, 24, 12, 24, 155, 3, 24, 158, 8, 24,
		1, 24, 1, 24, 1, 24, 3, 24, 163, 8, 24, 1, 25, 1, 25, 5, 25, 167, 8, 25,
		10, 25, 12, 25, 170, 9, 25, 1, 26, 1, 26, 5, 26, 174, 8, 26, 10, 26, 12,
		26, 177, 9, 26, 1, 26, 1, 26, 1, 27, 4, 27, 182, 8, 27, 11, 27, 12, 27,
		183, 1, 27, 1, 27, 0, 0, 28, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 1, 0, 5, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10,
		13, 13, 32, 32, 193, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 1, 57, 1, 0, 0, 0, 3, 61, 1, 0,
		0, 0, 5, 63, 1, 0, 0, 0, 7, 65, 1, 0, 0, 0, 9, 71, 1, 0, 0, 0, 11, 77,
		1, 0, 0, 0, 13, 82, 1, 0, 0, 0, 15, 85, 1, 0, 0, 0, 17, 88, 1, 0, 0, 0,
		19, 90, 1, 0, 0, 0, 21, 96, 1, 0, 0, 0, 23, 98, 1, 0, 0, 0, 25, 100, 1,
		0, 0, 0, 27, 102, 1, 0, 0, 0, 29, 104, 1, 0, 0, 0, 31, 106, 1, 0, 0, 0,
		33, 109, 1, 0, 0, 0, 35, 113, 1, 0, 0, 0, 37, 117, 1, 0, 0, 0, 39, 122,
		1, 0, 0, 0, 41, 124, 1, 0, 0, 0, 43, 130, 1, 0, 0, 0, 45, 132, 1, 0, 0,
		0, 47, 140, 1, 0, 0, 0, 49, 147, 1, 0, 0, 0, 51, 164, 1, 0, 0, 0, 53, 171,
		1, 0, 0, 0, 55, 181, 1, 0, 0, 0, 57, 58, 5, 102, 0, 0, 58, 59, 5, 115,
		0, 0, 59, 60, 5, 109, 0, 0, 60, 2, 1, 0, 0, 0, 61, 62, 5, 123, 0, 0, 62,
		4, 1, 0, 0, 0, 63, 64, 5, 125, 0, 0, 64, 6, 1, 0, 0, 0, 65, 66, 5, 115,
		0, 0, 66, 67, 5, 116, 0, 0, 67, 68, 5, 97, 0, 0, 68, 69, 5, 116, 0, 0,
		69, 70, 5, 101, 0, 0, 70, 8, 1, 0, 0, 0, 71, 72, 5, 101, 0, 0, 72, 73,
		5, 110, 0, 0, 73, 74, 5, 116, 0, 0, 74, 75, 5, 114, 0, 0, 75, 76, 5, 121,
		0, 0, 76, 10, 1, 0, 0, 0, 77, 78, 5, 101, 0, 0, 78, 79, 5, 120, 0, 0, 79,
		80, 5, 105, 0, 0, 80, 81, 5, 116, 0, 0, 81, 12, 1, 0, 0, 0, 82, 83, 5,
		100, 0, 0, 83, 84, 5, 111, 0, 0, 84, 14, 1, 0, 0, 0, 85, 86, 5, 111, 0,
		0, 86, 87, 5, 110, 0, 0, 87, 16, 1, 0, 0, 0, 88, 89, 5, 47, 0, 0, 89, 18,
		1, 0, 0, 0, 90, 91, 5, 97, 0, 0, 91, 92, 5, 102, 0, 0, 92, 93, 5, 116,
		0, 0, 93, 94, 5, 101, 0, 0, 94, 95, 5, 114, 0, 0, 95, 20, 1, 0, 0, 0, 96,
		97, 5, 40, 0, 0, 97, 22, 1, 0, 0, 0, 98, 99, 5, 41, 0, 0, 99, 24, 1, 0,
		0, 0, 100, 101, 5, 44, 0, 0, 101, 26, 1, 0, 0, 0, 102, 103, 5, 91, 0, 0,
		103, 28, 1, 0, 0, 0, 104, 105, 5, 93, 0, 0, 105, 30, 1, 0, 0, 0, 106, 107,
		5, 111, 0, 0, 107, 108, 5, 114, 0, 0, 108, 32, 1, 0, 0, 0, 109, 110, 5,
		97, 0, 0, 110, 111, 5, 110, 0, 0, 111, 112, 5, 100, 0, 0, 112, 34, 1, 0,
		0, 0, 113, 114, 5, 110, 0, 0, 114, 115, 5, 111, 0, 0, 115, 116, 5, 116,
		0, 0, 116, 36, 1, 0, 0, 0, 117, 118, 5, 103, 0, 0, 118, 119, 5, 111, 0,
		0, 119, 120, 5, 116, 0, 0, 120, 121, 5, 111, 0, 0, 121, 38, 1, 0, 0, 0,
		122, 123, 5, 46, 0, 0, 123, 40, 1, 0, 0, 0, 124, 125, 5, 102, 0, 0, 125,
		126, 5, 105, 0, 0, 126, 127, 5, 110, 0, 0, 127, 128, 5, 97, 0, 0, 128,
		129, 5, 108, 0, 0, 129, 42, 1, 0, 0, 0, 130, 131, 5, 72, 0, 0, 131, 44,
		1, 0, 0, 0, 132, 133, 5, 105, 0, 0, 133, 134, 5, 110, 0, 0, 134, 135, 5,
		105, 0, 0, 135, 136, 5, 116, 0, 0, 136, 137, 5, 105, 0, 0, 137, 138, 5,
		97, 0, 0, 138, 139, 5, 108, 0, 0, 139, 46, 1, 0, 0, 0, 140, 141, 5, 100,
		0, 0, 141, 142, 5, 101, 0, 0, 142, 143, 5, 102, 0, 0, 143, 144, 5, 101,
		0, 0, 144, 145, 5, 114, 0, 0, 145, 48, 1, 0, 0, 0, 146, 148, 7, 0, 0, 0,
		147, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 157, 1, 0, 0, 0, 151, 153, 5, 46, 0, 0, 152, 154,
		7, 0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0,
		0, 0, 155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157, 151, 1, 0, 0, 0,
		157, 158, 1, 0, 0, 0, 158, 162, 1, 0, 0, 0, 159, 160, 5, 109, 0, 0, 160,
		163, 5, 115, 0, 0, 161, 163, 5, 115, 0, 0, 162, 159, 1, 0, 0, 0, 162, 161,
		1, 0, 0, 0, 163, 50, 1, 0, 0, 0, 164, 168, 7, 1, 0, 0, 165, 167, 7, 2,
		0, 0, 166, 165, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0,
		168, 169, 1, 0, 0, 0, 169, 52, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 175,
		5, 35, 0, 0, 172, 174, 8, 3, 0, 0, 173, 172, 1, 0, 0, 0, 174, 177, 1, 0,
		0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 1, 0, 0, 0,
		177, 175, 1, 0, 0, 0, 178, 179, 6, 26, 0, 0, 179, 54, 1, 0, 0, 0, 180,
		182, 7, 4, 0, 0, 181, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 181,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 6, 27,
		1, 0, 186, 56, 1, 0, 0, 0, 8, 0, 149, 155, 157, 162, 168, 175, 183, 2,
		6, 0, 0, 0, 1, 0,
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
	fsmLexerT__20      = 21
	fsmLexerT__21      = 22
	fsmLexerInitial    = 23
	fsmLexerDefer      = 24
	fsmLexerDuration   = 25
	fsmLexerIdentifier = 26
	fsmLexerComment    = 27
	fsmLexerBlank      = 28
)
