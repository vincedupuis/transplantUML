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
		"", "'{'", "'}'", "'<<'", "'>>'", "'['", "']'", "'/'", "'('", "')'",
		"','", "'.'", "'fsm'", "'state'", "'parallel'", "'submachine'", "'region'",
		"'final'", "'terminate'", "'choice'", "'junction'", "'fork'", "'join'",
		"'entry'", "'exit'", "'point'", "'do'", "'on'", "'after'", "'else'",
		"'goto'", "'local'", "'initial'", "'defer'", "'invariant'", "'and'",
		"'or'", "'not'", "'H'", "'H*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "FSM", "STATE", "PARALLEL",
		"SUBMACHINE", "REGION", "FINAL", "TERMINATE", "CHOICE", "JUNCTION",
		"FORK", "JOIN", "ENTRY", "EXIT", "POINT", "DO", "ON", "AFTER", "ELSE",
		"GOTO", "LOCAL", "INITIAL", "DEFER", "INVARIANT", "AND", "OR", "NOT",
		"H", "H_DEEP", "Duration", "Identifier", "Note", "Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"fsm", "state", "parallel", "submachine", "final", "terminate", "region",
		"choice", "junction", "fork", "join", "point", "reference", "history",
		"start", "stereotype", "branch", "event", "trigger", "actions", "identifiers",
		"guard", "expression", "or_expression", "and_expression", "not_expression",
		"single_expression", "goto",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 455, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 3, 0, 58, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 75, 8,
		0, 10, 0, 12, 0, 78, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 3, 1, 84, 8, 1, 1, 1,
		3, 1, 87, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 92, 8, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 108,
		8, 1, 10, 1, 12, 1, 111, 9, 1, 1, 1, 1, 1, 1, 2, 3, 2, 116, 8, 2, 1, 2,
		3, 2, 119, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 125, 8, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 5, 2, 131, 8, 2, 10, 2, 12, 2, 134, 9, 2, 1, 2, 1, 2, 1, 3, 3,
		3, 139, 8, 3, 1, 3, 3, 3, 142, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 147, 8, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 153, 8, 3, 10, 3, 12, 3, 156, 9, 3, 1, 3,
		1, 3, 1, 4, 3, 4, 161, 8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 166, 8, 4, 1, 4, 3,
		4, 169, 8, 4, 1, 5, 3, 5, 172, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 177, 8, 5,
		1, 5, 3, 5, 180, 8, 5, 1, 6, 3, 6, 183, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 188,
		8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 5, 6, 202, 8, 6, 10, 6, 12, 6, 205, 9, 6, 1, 6, 1, 6, 1, 7, 3, 7,
		210, 8, 7, 1, 7, 3, 7, 213, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 218, 8, 7, 1,
		7, 1, 7, 5, 7, 222, 8, 7, 10, 7, 12, 7, 225, 9, 7, 1, 7, 1, 7, 3, 7, 229,
		8, 7, 1, 8, 3, 8, 232, 8, 8, 1, 8, 3, 8, 235, 8, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 240, 8, 8, 1, 8, 1, 8, 5, 8, 244, 8, 8, 10, 8, 12, 8, 247, 9, 8, 1,
		8, 1, 8, 3, 8, 251, 8, 8, 1, 9, 3, 9, 254, 8, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		259, 8, 9, 1, 9, 1, 9, 3, 9, 263, 8, 9, 1, 9, 3, 9, 266, 8, 9, 1, 9, 5,
		9, 269, 8, 9, 10, 9, 12, 9, 272, 9, 9, 1, 9, 1, 9, 1, 10, 3, 10, 277, 8,
		10, 1, 10, 1, 10, 1, 10, 3, 10, 282, 8, 10, 1, 10, 3, 10, 285, 8, 10, 1,
		10, 1, 10, 1, 11, 3, 11, 290, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11,
		296, 8, 11, 1, 11, 3, 11, 299, 8, 11, 1, 11, 1, 11, 1, 12, 3, 12, 304,
		8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 310, 8, 12, 1, 13, 3, 13, 313,
		8, 13, 1, 13, 1, 13, 3, 13, 317, 8, 13, 1, 13, 3, 13, 320, 8, 13, 1, 13,
		3, 13, 323, 8, 13, 1, 14, 1, 14, 3, 14, 327, 8, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 3, 16, 337, 8, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 343, 8, 16, 1, 16, 3, 16, 346, 8, 16, 1, 16, 1, 16, 1,
		17, 3, 17, 351, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 3, 17, 363, 8, 17, 1, 17, 1, 17, 3, 17, 367, 8, 17,
		1, 17, 1, 17, 3, 17, 371, 8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 376, 8, 17,
		1, 17, 3, 17, 379, 8, 17, 1, 17, 3, 17, 382, 8, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 3, 18, 390, 8, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 5, 20, 398, 8, 20, 10, 20, 12, 20, 401, 9, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 412, 8, 23, 10,
		23, 12, 23, 415, 9, 23, 1, 24, 1, 24, 1, 24, 5, 24, 420, 8, 24, 10, 24,
		12, 24, 423, 9, 24, 1, 25, 3, 25, 426, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 3, 26, 435, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3,
		27, 441, 8, 27, 1, 27, 1, 27, 1, 27, 3, 27, 446, 8, 27, 1, 27, 1, 27, 3,
		27, 450, 8, 27, 1, 27, 3, 27, 453, 8, 27, 1, 27, 0, 0, 28, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 0, 5, 1, 0, 23, 24, 1, 0, 38, 39, 2, 0, 23, 24,
		26, 26, 1, 0, 40, 41, 1, 0, 17, 18, 540, 0, 57, 1, 0, 0, 0, 2, 83, 1, 0,
		0, 0, 4, 115, 1, 0, 0, 0, 6, 138, 1, 0, 0, 0, 8, 160, 1, 0, 0, 0, 10, 171,
		1, 0, 0, 0, 12, 182, 1, 0, 0, 0, 14, 209, 1, 0, 0, 0, 16, 231, 1, 0, 0,
		0, 18, 253, 1, 0, 0, 0, 20, 276, 1, 0, 0, 0, 22, 289, 1, 0, 0, 0, 24, 303,
		1, 0, 0, 0, 26, 312, 1, 0, 0, 0, 28, 324, 1, 0, 0, 0, 30, 331, 1, 0, 0,
		0, 32, 336, 1, 0, 0, 0, 34, 350, 1, 0, 0, 0, 36, 389, 1, 0, 0, 0, 38, 391,
		1, 0, 0, 0, 40, 394, 1, 0, 0, 0, 42, 402, 1, 0, 0, 0, 44, 406, 1, 0, 0,
		0, 46, 408, 1, 0, 0, 0, 48, 416, 1, 0, 0, 0, 50, 425, 1, 0, 0, 0, 52, 434,
		1, 0, 0, 0, 54, 452, 1, 0, 0, 0, 56, 58, 5, 42, 0, 0, 57, 56, 1, 0, 0,
		0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 5, 12, 0, 0, 60, 61,
		5, 41, 0, 0, 61, 76, 5, 1, 0, 0, 62, 75, 3, 28, 14, 0, 63, 75, 3, 2, 1,
		0, 64, 75, 3, 4, 2, 0, 65, 75, 3, 6, 3, 0, 66, 75, 3, 8, 4, 0, 67, 75,
		3, 10, 5, 0, 68, 75, 3, 14, 7, 0, 69, 75, 3, 16, 8, 0, 70, 75, 3, 18, 9,
		0, 71, 75, 3, 20, 10, 0, 72, 75, 3, 22, 11, 0, 73, 75, 3, 34, 17, 0, 74,
		62, 1, 0, 0, 0, 74, 63, 1, 0, 0, 0, 74, 64, 1, 0, 0, 0, 74, 65, 1, 0, 0,
		0, 74, 66, 1, 0, 0, 0, 74, 67, 1, 0, 0, 0, 74, 68, 1, 0, 0, 0, 74, 69,
		1, 0, 0, 0, 74, 70, 1, 0, 0, 0, 74, 71, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0,
		74, 73, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1,
		0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 5, 2, 0, 0, 80,
		81, 5, 0, 0, 1, 81, 1, 1, 0, 0, 0, 82, 84, 5, 42, 0, 0, 83, 82, 1, 0, 0,
		0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 87, 5, 32, 0, 0, 86, 85,
		1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 5, 13, 0, 0,
		89, 91, 5, 41, 0, 0, 90, 92, 3, 30, 15, 0, 91, 90, 1, 0, 0, 0, 91, 92,
		1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 109, 5, 1, 0, 0, 94, 108, 3, 28, 14,
		0, 95, 108, 3, 2, 1, 0, 96, 108, 3, 4, 2, 0, 97, 108, 3, 6, 3, 0, 98, 108,
		3, 8, 4, 0, 99, 108, 3, 10, 5, 0, 100, 108, 3, 14, 7, 0, 101, 108, 3, 16,
		8, 0, 102, 108, 3, 18, 9, 0, 103, 108, 3, 20, 10, 0, 104, 108, 3, 22, 11,
		0, 105, 108, 3, 26, 13, 0, 106, 108, 3, 34, 17, 0, 107, 94, 1, 0, 0, 0,
		107, 95, 1, 0, 0, 0, 107, 96, 1, 0, 0, 0, 107, 97, 1, 0, 0, 0, 107, 98,
		1, 0, 0, 0, 107, 99, 1, 0, 0, 0, 107, 100, 1, 0, 0, 0, 107, 101, 1, 0,
		0, 0, 107, 102, 1, 0, 0, 0, 107, 103, 1, 0, 0, 0, 107, 104, 1, 0, 0, 0,
		107, 105, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109,
		107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 109,
		1, 0, 0, 0, 112, 113, 5, 2, 0, 0, 113, 3, 1, 0, 0, 0, 114, 116, 5, 42,
		0, 0, 115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0, 0, 0,
		117, 119, 5, 32, 0, 0, 118, 117, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 121, 5, 14, 0, 0, 121, 122, 5, 13, 0, 0, 122, 124,
		5, 41, 0, 0, 123, 125, 3, 30, 15, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1,
		0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 132, 5, 1, 0, 0, 127, 131, 3, 12, 6,
		0, 128, 131, 3, 22, 11, 0, 129, 131, 3, 34, 17, 0, 130, 127, 1, 0, 0, 0,
		130, 128, 1, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132,
		130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 132,
		1, 0, 0, 0, 135, 136, 5, 2, 0, 0, 136, 5, 1, 0, 0, 0, 137, 139, 5, 42,
		0, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0,
		140, 142, 5, 32, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		143, 1, 0, 0, 0, 143, 144, 5, 15, 0, 0, 144, 146, 5, 41, 0, 0, 145, 147,
		3, 30, 15, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1,
		0, 0, 0, 148, 154, 5, 1, 0, 0, 149, 153, 3, 22, 11, 0, 150, 153, 3, 24,
		12, 0, 151, 153, 3, 34, 17, 0, 152, 149, 1, 0, 0, 0, 152, 150, 1, 0, 0,
		0, 152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158,
		5, 2, 0, 0, 158, 7, 1, 0, 0, 0, 159, 161, 5, 42, 0, 0, 160, 159, 1, 0,
		0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 5, 17, 0, 0,
		163, 165, 5, 13, 0, 0, 164, 166, 5, 41, 0, 0, 165, 164, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167, 169, 3, 30, 15, 0, 168, 167,
		1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 9, 1, 0, 0, 0, 170, 172, 5, 42,
		0, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0,
		173, 174, 5, 18, 0, 0, 174, 176, 5, 13, 0, 0, 175, 177, 5, 41, 0, 0, 176,
		175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178, 180,
		3, 30, 15, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 11, 1, 0,
		0, 0, 181, 183, 5, 42, 0, 0, 182, 181, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 185, 5, 16, 0, 0, 185, 187, 5, 41, 0, 0, 186,
		188, 3, 30, 15, 0, 187, 186, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189,
		1, 0, 0, 0, 189, 203, 5, 1, 0, 0, 190, 202, 3, 28, 14, 0, 191, 202, 3,
		2, 1, 0, 192, 202, 3, 4, 2, 0, 193, 202, 3, 6, 3, 0, 194, 202, 3, 8, 4,
		0, 195, 202, 3, 10, 5, 0, 196, 202, 3, 14, 7, 0, 197, 202, 3, 16, 8, 0,
		198, 202, 3, 18, 9, 0, 199, 202, 3, 20, 10, 0, 200, 202, 3, 26, 13, 0,
		201, 190, 1, 0, 0, 0, 201, 191, 1, 0, 0, 0, 201, 192, 1, 0, 0, 0, 201,
		193, 1, 0, 0, 0, 201, 194, 1, 0, 0, 0, 201, 195, 1, 0, 0, 0, 201, 196,
		1, 0, 0, 0, 201, 197, 1, 0, 0, 0, 201, 198, 1, 0, 0, 0, 201, 199, 1, 0,
		0, 0, 201, 200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0,
		203, 204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206,
		207, 5, 2, 0, 0, 207, 13, 1, 0, 0, 0, 208, 210, 5, 42, 0, 0, 209, 208,
		1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212, 1, 0, 0, 0, 211, 213, 5, 32,
		0, 0, 212, 211, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0,
		214, 215, 5, 19, 0, 0, 215, 217, 5, 41, 0, 0, 216, 218, 3, 30, 15, 0, 217,
		216, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 228, 1, 0, 0, 0, 219, 223,
		5, 1, 0, 0, 220, 222, 3, 32, 16, 0, 221, 220, 1, 0, 0, 0, 222, 225, 1,
		0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 226, 1, 0, 0,
		0, 225, 223, 1, 0, 0, 0, 226, 229, 5, 2, 0, 0, 227, 229, 3, 32, 16, 0,
		228, 219, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229, 15, 1, 0, 0, 0, 230, 232,
		5, 42, 0, 0, 231, 230, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0,
		0, 0, 233, 235, 5, 32, 0, 0, 234, 233, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0,
		235, 236, 1, 0, 0, 0, 236, 237, 5, 20, 0, 0, 237, 239, 5, 41, 0, 0, 238,
		240, 3, 30, 15, 0, 239, 238, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 250,
		1, 0, 0, 0, 241, 245, 5, 1, 0, 0, 242, 244, 3, 32, 16, 0, 243, 242, 1,
		0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0,
		0, 246, 248, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 251, 5, 2, 0, 0, 249,
		251, 3, 32, 16, 0, 250, 241, 1, 0, 0, 0, 250, 249, 1, 0, 0, 0, 251, 17,
		1, 0, 0, 0, 252, 254, 5, 42, 0, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0,
		0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 5, 21, 0, 0, 256, 258, 5, 41, 0,
		0, 257, 259, 3, 30, 15, 0, 258, 257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0,
		259, 260, 1, 0, 0, 0, 260, 270, 5, 1, 0, 0, 261, 263, 5, 42, 0, 0, 262,
		261, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0, 0, 0, 264, 266,
		3, 38, 19, 0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1,
		0, 0, 0, 267, 269, 3, 54, 27, 0, 268, 262, 1, 0, 0, 0, 269, 272, 1, 0,
		0, 0, 270, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 273, 1, 0, 0, 0,
		272, 270, 1, 0, 0, 0, 273, 274, 5, 2, 0, 0, 274, 19, 1, 0, 0, 0, 275, 277,
		5, 42, 0, 0, 276, 275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 1, 0,
		0, 0, 278, 279, 5, 22, 0, 0, 279, 281, 5, 41, 0, 0, 280, 282, 3, 30, 15,
		0, 281, 280, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 284, 1, 0, 0, 0, 283,
		285, 3, 38, 19, 0, 284, 283, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286,
		1, 0, 0, 0, 286, 287, 3, 54, 27, 0, 287, 21, 1, 0, 0, 0, 288, 290, 5, 42,
		0, 0, 289, 288, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0,
		291, 292, 7, 0, 0, 0, 292, 293, 5, 25, 0, 0, 293, 295, 5, 41, 0, 0, 294,
		296, 3, 30, 15, 0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298,
		1, 0, 0, 0, 297, 299, 3, 38, 19, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1,
		0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 301, 3, 54, 27, 0, 301, 23, 1, 0, 0,
		0, 302, 304, 5, 42, 0, 0, 303, 302, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304,
		305, 1, 0, 0, 0, 305, 306, 5, 23, 0, 0, 306, 307, 5, 25, 0, 0, 307, 309,
		5, 41, 0, 0, 308, 310, 3, 30, 15, 0, 309, 308, 1, 0, 0, 0, 309, 310, 1,
		0, 0, 0, 310, 25, 1, 0, 0, 0, 311, 313, 5, 42, 0, 0, 312, 311, 1, 0, 0,
		0, 312, 313, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316, 7, 1, 0, 0, 315,
		317, 3, 30, 15, 0, 316, 315, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 322,
		1, 0, 0, 0, 318, 320, 3, 38, 19, 0, 319, 318, 1, 0, 0, 0, 319, 320, 1,
		0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 323, 3, 54, 27, 0, 322, 319, 1, 0,
		0, 0, 322, 323, 1, 0, 0, 0, 323, 27, 1, 0, 0, 0, 324, 326, 5, 32, 0, 0,
		325, 327, 3, 38, 19, 0, 326, 325, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327,
		328, 1, 0, 0, 0, 328, 329, 5, 30, 0, 0, 329, 330, 5, 41, 0, 0, 330, 29,
		1, 0, 0, 0, 331, 332, 5, 3, 0, 0, 332, 333, 5, 41, 0, 0, 333, 334, 5, 4,
		0, 0, 334, 31, 1, 0, 0, 0, 335, 337, 5, 42, 0, 0, 336, 335, 1, 0, 0, 0,
		336, 337, 1, 0, 0, 0, 337, 342, 1, 0, 0, 0, 338, 339, 5, 5, 0, 0, 339,
		340, 5, 29, 0, 0, 340, 343, 5, 6, 0, 0, 341, 343, 3, 42, 21, 0, 342, 338,
		1, 0, 0, 0, 342, 341, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 345, 1, 0,
		0, 0, 344, 346, 3, 38, 19, 0, 345, 344, 1, 0, 0, 0, 345, 346, 1, 0, 0,
		0, 346, 347, 1, 0, 0, 0, 347, 348, 3, 54, 27, 0, 348, 33, 1, 0, 0, 0, 349,
		351, 5, 42, 0, 0, 350, 349, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 381,
		1, 0, 0, 0, 352, 353, 7, 2, 0, 0, 353, 382, 3, 38, 19, 0, 354, 355, 5,
		27, 0, 0, 355, 356, 5, 41, 0, 0, 356, 357, 5, 7, 0, 0, 357, 382, 5, 33,
		0, 0, 358, 359, 5, 34, 0, 0, 359, 382, 3, 42, 21, 0, 360, 362, 3, 36, 18,
		0, 361, 363, 3, 42, 21, 0, 362, 361, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0,
		363, 364, 1, 0, 0, 0, 364, 366, 3, 38, 19, 0, 365, 367, 3, 54, 27, 0, 366,
		365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 382, 1, 0, 0, 0, 368, 370,
		3, 36, 18, 0, 369, 371, 3, 42, 21, 0, 370, 369, 1, 0, 0, 0, 370, 371, 1,
		0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 373, 3, 54, 27, 0, 373, 382, 1, 0,
		0, 0, 374, 376, 3, 42, 21, 0, 375, 374, 1, 0, 0, 0, 375, 376, 1, 0, 0,
		0, 376, 378, 1, 0, 0, 0, 377, 379, 3, 38, 19, 0, 378, 377, 1, 0, 0, 0,
		378, 379, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 382, 3, 54, 27, 0, 381,
		352, 1, 0, 0, 0, 381, 354, 1, 0, 0, 0, 381, 358, 1, 0, 0, 0, 381, 360,
		1, 0, 0, 0, 381, 368, 1, 0, 0, 0, 381, 375, 1, 0, 0, 0, 382, 35, 1, 0,
		0, 0, 383, 384, 5, 27, 0, 0, 384, 390, 5, 41, 0, 0, 385, 386, 5, 28, 0,
		0, 386, 387, 5, 8, 0, 0, 387, 388, 7, 3, 0, 0, 388, 390, 5, 9, 0, 0, 389,
		383, 1, 0, 0, 0, 389, 385, 1, 0, 0, 0, 390, 37, 1, 0, 0, 0, 391, 392, 5,
		7, 0, 0, 392, 393, 3, 40, 20, 0, 393, 39, 1, 0, 0, 0, 394, 399, 5, 41,
		0, 0, 395, 396, 5, 10, 0, 0, 396, 398, 5, 41, 0, 0, 397, 395, 1, 0, 0,
		0, 398, 401, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400,
		41, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 402, 403, 5, 5, 0, 0, 403, 404, 3,
		44, 22, 0, 404, 405, 5, 6, 0, 0, 405, 43, 1, 0, 0, 0, 406, 407, 3, 46,
		23, 0, 407, 45, 1, 0, 0, 0, 408, 413, 3, 48, 24, 0, 409, 410, 5, 36, 0,
		0, 410, 412, 3, 48, 24, 0, 411, 409, 1, 0, 0, 0, 412, 415, 1, 0, 0, 0,
		413, 411, 1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 47, 1, 0, 0, 0, 415, 413,
		1, 0, 0, 0, 416, 421, 3, 50, 25, 0, 417, 418, 5, 35, 0, 0, 418, 420, 3,
		50, 25, 0, 419, 417, 1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421, 419, 1, 0,
		0, 0, 421, 422, 1, 0, 0, 0, 422, 49, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0,
		424, 426, 5, 37, 0, 0, 425, 424, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426,
		427, 1, 0, 0, 0, 427, 428, 3, 52, 26, 0, 428, 51, 1, 0, 0, 0, 429, 435,
		5, 41, 0, 0, 430, 431, 5, 8, 0, 0, 431, 432, 3, 44, 22, 0, 432, 433, 5,
		9, 0, 0, 433, 435, 1, 0, 0, 0, 434, 429, 1, 0, 0, 0, 434, 430, 1, 0, 0,
		0, 435, 53, 1, 0, 0, 0, 436, 437, 5, 30, 0, 0, 437, 453, 7, 4, 0, 0, 438,
		440, 5, 30, 0, 0, 439, 441, 5, 31, 0, 0, 440, 439, 1, 0, 0, 0, 440, 441,
		1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 453, 5, 41, 0, 0, 443, 445, 5, 30,
		0, 0, 444, 446, 5, 31, 0, 0, 445, 444, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0,
		446, 449, 1, 0, 0, 0, 447, 448, 5, 41, 0, 0, 448, 450, 5, 11, 0, 0, 449,
		447, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 453,
		7, 1, 0, 0, 452, 436, 1, 0, 0, 0, 452, 438, 1, 0, 0, 0, 452, 443, 1, 0,
		0, 0, 453, 55, 1, 0, 0, 0, 76, 57, 74, 76, 83, 86, 91, 107, 109, 115, 118,
		124, 130, 132, 138, 141, 146, 152, 154, 160, 165, 168, 171, 176, 179, 182,
		187, 201, 203, 209, 212, 217, 223, 228, 231, 234, 239, 245, 250, 253, 258,
		262, 265, 270, 276, 281, 284, 289, 295, 298, 303, 309, 312, 316, 319, 322,
		326, 336, 342, 345, 350, 362, 366, 370, 375, 378, 381, 389, 399, 413, 421,
		425, 434, 440, 445, 449, 452,
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
	fsmParserFSM        = 12
	fsmParserSTATE      = 13
	fsmParserPARALLEL   = 14
	fsmParserSUBMACHINE = 15
	fsmParserREGION     = 16
	fsmParserFINAL      = 17
	fsmParserTERMINATE  = 18
	fsmParserCHOICE     = 19
	fsmParserJUNCTION   = 20
	fsmParserFORK       = 21
	fsmParserJOIN       = 22
	fsmParserENTRY      = 23
	fsmParserEXIT       = 24
	fsmParserPOINT      = 25
	fsmParserDO         = 26
	fsmParserON         = 27
	fsmParserAFTER      = 28
	fsmParserELSE       = 29
	fsmParserGOTO       = 30
	fsmParserLOCAL      = 31
	fsmParserINITIAL    = 32
	fsmParserDEFER      = 33
	fsmParserINVARIANT  = 34
	fsmParserAND        = 35
	fsmParserOR         = 36
	fsmParserNOT        = 37
	fsmParserH          = 38
	fsmParserH_DEEP     = 39
	fsmParserDuration   = 40
	fsmParserIdentifier = 41
	fsmParserNote       = 42
	fsmParserComment    = 43
	fsmParserBlank      = 44
)

// fsmParser rules.
const (
	fsmParserRULE_fsm               = 0
	fsmParserRULE_state             = 1
	fsmParserRULE_parallel          = 2
	fsmParserRULE_submachine        = 3
	fsmParserRULE_final             = 4
	fsmParserRULE_terminate         = 5
	fsmParserRULE_region            = 6
	fsmParserRULE_choice            = 7
	fsmParserRULE_junction          = 8
	fsmParserRULE_fork              = 9
	fsmParserRULE_join              = 10
	fsmParserRULE_point             = 11
	fsmParserRULE_reference         = 12
	fsmParserRULE_history           = 13
	fsmParserRULE_start             = 14
	fsmParserRULE_stereotype        = 15
	fsmParserRULE_branch            = 16
	fsmParserRULE_event             = 17
	fsmParserRULE_trigger           = 18
	fsmParserRULE_actions           = 19
	fsmParserRULE_identifiers       = 20
	fsmParserRULE_guard             = 21
	fsmParserRULE_expression        = 22
	fsmParserRULE_or_expression     = 23
	fsmParserRULE_and_expression    = 24
	fsmParserRULE_not_expression    = 25
	fsmParserRULE_single_expression = 26
	fsmParserRULE_goto              = 27
)

// IFsmContext is an interface to support dynamic dispatch.
type IFsmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FSM() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	EOF() antlr.TerminalNode
	Note() antlr.TerminalNode
	AllStart_() []IStartContext
	Start_(i int) IStartContext
	AllState() []IStateContext
	State(i int) IStateContext
	AllParallel() []IParallelContext
	Parallel(i int) IParallelContext
	AllSubmachine() []ISubmachineContext
	Submachine(i int) ISubmachineContext
	AllFinal() []IFinalContext
	Final(i int) IFinalContext
	AllTerminate() []ITerminateContext
	Terminate(i int) ITerminateContext
	AllChoice() []IChoiceContext
	Choice(i int) IChoiceContext
	AllJunction() []IJunctionContext
	Junction(i int) IJunctionContext
	AllFork() []IForkContext
	Fork(i int) IForkContext
	AllJoin() []IJoinContext
	Join(i int) IJoinContext
	AllPoint() []IPointContext
	Point(i int) IPointContext
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

func (s *FsmContext) FSM() antlr.TerminalNode {
	return s.GetToken(fsmParserFSM, 0)
}

func (s *FsmContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *FsmContext) EOF() antlr.TerminalNode {
	return s.GetToken(fsmParserEOF, 0)
}

func (s *FsmContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *FsmContext) AllStart_() []IStartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStartContext); ok {
			len++
		}
	}

	tst := make([]IStartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStartContext); ok {
			tst[i] = t.(IStartContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Start_(i int) IStartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartContext); ok {
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

	return t.(IStartContext)
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

func (s *FsmContext) AllParallel() []IParallelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParallelContext); ok {
			len++
		}
	}

	tst := make([]IParallelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParallelContext); ok {
			tst[i] = t.(IParallelContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Parallel(i int) IParallelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParallelContext); ok {
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

	return t.(IParallelContext)
}

func (s *FsmContext) AllSubmachine() []ISubmachineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubmachineContext); ok {
			len++
		}
	}

	tst := make([]ISubmachineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubmachineContext); ok {
			tst[i] = t.(ISubmachineContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Submachine(i int) ISubmachineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubmachineContext); ok {
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

	return t.(ISubmachineContext)
}

func (s *FsmContext) AllFinal() []IFinalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFinalContext); ok {
			len++
		}
	}

	tst := make([]IFinalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFinalContext); ok {
			tst[i] = t.(IFinalContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Final(i int) IFinalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFinalContext); ok {
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

	return t.(IFinalContext)
}

func (s *FsmContext) AllTerminate() []ITerminateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminateContext); ok {
			len++
		}
	}

	tst := make([]ITerminateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminateContext); ok {
			tst[i] = t.(ITerminateContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Terminate(i int) ITerminateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminateContext); ok {
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

	return t.(ITerminateContext)
}

func (s *FsmContext) AllChoice() []IChoiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChoiceContext); ok {
			len++
		}
	}

	tst := make([]IChoiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChoiceContext); ok {
			tst[i] = t.(IChoiceContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Choice(i int) IChoiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChoiceContext); ok {
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

	return t.(IChoiceContext)
}

func (s *FsmContext) AllJunction() []IJunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJunctionContext); ok {
			len++
		}
	}

	tst := make([]IJunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJunctionContext); ok {
			tst[i] = t.(IJunctionContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Junction(i int) IJunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJunctionContext); ok {
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

	return t.(IJunctionContext)
}

func (s *FsmContext) AllFork() []IForkContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForkContext); ok {
			len++
		}
	}

	tst := make([]IForkContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForkContext); ok {
			tst[i] = t.(IForkContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Fork(i int) IForkContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForkContext); ok {
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

	return t.(IForkContext)
}

func (s *FsmContext) AllJoin() []IJoinContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinContext); ok {
			len++
		}
	}

	tst := make([]IJoinContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinContext); ok {
			tst[i] = t.(IJoinContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Join(i int) IJoinContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinContext); ok {
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

	return t.(IJoinContext)
}

func (s *FsmContext) AllPoint() []IPointContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPointContext); ok {
			len++
		}
	}

	tst := make([]IPointContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPointContext); ok {
			tst[i] = t.(IPointContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Point(i int) IPointContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointContext); ok {
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

	return t.(IPointContext)
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(56)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(59)
		p.Match(fsmParserFSM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4421098332320) != 0 {
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(62)
				p.Start_()
			}

		case 2:
			{
				p.SetState(63)
				p.State()
			}

		case 3:
			{
				p.SetState(64)
				p.Parallel()
			}

		case 4:
			{
				p.SetState(65)
				p.Submachine()
			}

		case 5:
			{
				p.SetState(66)
				p.Final()
			}

		case 6:
			{
				p.SetState(67)
				p.Terminate()
			}

		case 7:
			{
				p.SetState(68)
				p.Choice()
			}

		case 8:
			{
				p.SetState(69)
				p.Junction()
			}

		case 9:
			{
				p.SetState(70)
				p.Fork()
			}

		case 10:
			{
				p.SetState(71)
				p.Join()
			}

		case 11:
			{
				p.SetState(72)
				p.Point()
			}

		case 12:
			{
				p.SetState(73)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(79)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
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
	STATE() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	INITIAL() antlr.TerminalNode
	Stereotype() IStereotypeContext
	AllStart_() []IStartContext
	Start_(i int) IStartContext
	AllState() []IStateContext
	State(i int) IStateContext
	AllParallel() []IParallelContext
	Parallel(i int) IParallelContext
	AllSubmachine() []ISubmachineContext
	Submachine(i int) ISubmachineContext
	AllFinal() []IFinalContext
	Final(i int) IFinalContext
	AllTerminate() []ITerminateContext
	Terminate(i int) ITerminateContext
	AllChoice() []IChoiceContext
	Choice(i int) IChoiceContext
	AllJunction() []IJunctionContext
	Junction(i int) IJunctionContext
	AllFork() []IForkContext
	Fork(i int) IForkContext
	AllJoin() []IJoinContext
	Join(i int) IJoinContext
	AllPoint() []IPointContext
	Point(i int) IPointContext
	AllHistory() []IHistoryContext
	History(i int) IHistoryContext
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

func (s *StateContext) STATE() antlr.TerminalNode {
	return s.GetToken(fsmParserSTATE, 0)
}

func (s *StateContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *StateContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *StateContext) INITIAL() antlr.TerminalNode {
	return s.GetToken(fsmParserINITIAL, 0)
}

func (s *StateContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *StateContext) AllStart_() []IStartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStartContext); ok {
			len++
		}
	}

	tst := make([]IStartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStartContext); ok {
			tst[i] = t.(IStartContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Start_(i int) IStartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartContext); ok {
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

	return t.(IStartContext)
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

func (s *StateContext) AllParallel() []IParallelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParallelContext); ok {
			len++
		}
	}

	tst := make([]IParallelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParallelContext); ok {
			tst[i] = t.(IParallelContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Parallel(i int) IParallelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParallelContext); ok {
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

	return t.(IParallelContext)
}

func (s *StateContext) AllSubmachine() []ISubmachineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubmachineContext); ok {
			len++
		}
	}

	tst := make([]ISubmachineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubmachineContext); ok {
			tst[i] = t.(ISubmachineContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Submachine(i int) ISubmachineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubmachineContext); ok {
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

	return t.(ISubmachineContext)
}

func (s *StateContext) AllFinal() []IFinalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFinalContext); ok {
			len++
		}
	}

	tst := make([]IFinalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFinalContext); ok {
			tst[i] = t.(IFinalContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Final(i int) IFinalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFinalContext); ok {
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

	return t.(IFinalContext)
}

func (s *StateContext) AllTerminate() []ITerminateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminateContext); ok {
			len++
		}
	}

	tst := make([]ITerminateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminateContext); ok {
			tst[i] = t.(ITerminateContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Terminate(i int) ITerminateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminateContext); ok {
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

	return t.(ITerminateContext)
}

func (s *StateContext) AllChoice() []IChoiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChoiceContext); ok {
			len++
		}
	}

	tst := make([]IChoiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChoiceContext); ok {
			tst[i] = t.(IChoiceContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Choice(i int) IChoiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChoiceContext); ok {
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

	return t.(IChoiceContext)
}

func (s *StateContext) AllJunction() []IJunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJunctionContext); ok {
			len++
		}
	}

	tst := make([]IJunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJunctionContext); ok {
			tst[i] = t.(IJunctionContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Junction(i int) IJunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJunctionContext); ok {
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

	return t.(IJunctionContext)
}

func (s *StateContext) AllFork() []IForkContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForkContext); ok {
			len++
		}
	}

	tst := make([]IForkContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForkContext); ok {
			tst[i] = t.(IForkContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Fork(i int) IForkContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForkContext); ok {
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

	return t.(IForkContext)
}

func (s *StateContext) AllJoin() []IJoinContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinContext); ok {
			len++
		}
	}

	tst := make([]IJoinContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinContext); ok {
			tst[i] = t.(IJoinContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Join(i int) IJoinContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinContext); ok {
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

	return t.(IJoinContext)
}

func (s *StateContext) AllPoint() []IPointContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPointContext); ok {
			len++
		}
	}

	tst := make([]IPointContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPointContext); ok {
			tst[i] = t.(IPointContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) Point(i int) IPointContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointContext); ok {
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

	return t.(IPointContext)
}

func (s *StateContext) AllHistory() []IHistoryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHistoryContext); ok {
			len++
		}
	}

	tst := make([]IHistoryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHistoryContext); ok {
			tst[i] = t.(IHistoryContext)
			i++
		}
	}

	return tst
}

func (s *StateContext) History(i int) IHistoryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHistoryContext); ok {
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

	return t.(IHistoryContext)
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(82)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(85)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(88)
		p.Match(fsmParserSTATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(90)
			p.Stereotype()
		}

	}
	{
		p.SetState(93)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5245732053152) != 0 {
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(94)
				p.Start_()
			}

		case 2:
			{
				p.SetState(95)
				p.State()
			}

		case 3:
			{
				p.SetState(96)
				p.Parallel()
			}

		case 4:
			{
				p.SetState(97)
				p.Submachine()
			}

		case 5:
			{
				p.SetState(98)
				p.Final()
			}

		case 6:
			{
				p.SetState(99)
				p.Terminate()
			}

		case 7:
			{
				p.SetState(100)
				p.Choice()
			}

		case 8:
			{
				p.SetState(101)
				p.Junction()
			}

		case 9:
			{
				p.SetState(102)
				p.Fork()
			}

		case 10:
			{
				p.SetState(103)
				p.Join()
			}

		case 11:
			{
				p.SetState(104)
				p.Point()
			}

		case 12:
			{
				p.SetState(105)
				p.History()
			}

		case 13:
			{
				p.SetState(106)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
		p.Match(fsmParserT__1)
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

// IParallelContext is an interface to support dynamic dispatch.
type IParallelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARALLEL() antlr.TerminalNode
	STATE() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	INITIAL() antlr.TerminalNode
	Stereotype() IStereotypeContext
	AllRegion() []IRegionContext
	Region(i int) IRegionContext
	AllPoint() []IPointContext
	Point(i int) IPointContext
	AllEvent() []IEventContext
	Event(i int) IEventContext

	// IsParallelContext differentiates from other interfaces.
	IsParallelContext()
}

type ParallelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParallelContext() *ParallelContext {
	var p = new(ParallelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_parallel
	return p
}

func InitEmptyParallelContext(p *ParallelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_parallel
}

func (*ParallelContext) IsParallelContext() {}

func NewParallelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParallelContext {
	var p = new(ParallelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_parallel

	return p
}

func (s *ParallelContext) GetParser() antlr.Parser { return s.parser }

func (s *ParallelContext) PARALLEL() antlr.TerminalNode {
	return s.GetToken(fsmParserPARALLEL, 0)
}

func (s *ParallelContext) STATE() antlr.TerminalNode {
	return s.GetToken(fsmParserSTATE, 0)
}

func (s *ParallelContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *ParallelContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *ParallelContext) INITIAL() antlr.TerminalNode {
	return s.GetToken(fsmParserINITIAL, 0)
}

func (s *ParallelContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *ParallelContext) AllRegion() []IRegionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegionContext); ok {
			len++
		}
	}

	tst := make([]IRegionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegionContext); ok {
			tst[i] = t.(IRegionContext)
			i++
		}
	}

	return tst
}

func (s *ParallelContext) Region(i int) IRegionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegionContext); ok {
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

	return t.(IRegionContext)
}

func (s *ParallelContext) AllPoint() []IPointContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPointContext); ok {
			len++
		}
	}

	tst := make([]IPointContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPointContext); ok {
			tst[i] = t.(IPointContext)
			i++
		}
	}

	return tst
}

func (s *ParallelContext) Point(i int) IPointContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointContext); ok {
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

	return t.(IPointContext)
}

func (s *ParallelContext) AllEvent() []IEventContext {
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

func (s *ParallelContext) Event(i int) IEventContext {
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

func (s *ParallelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParallelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParallelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitParallel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Parallel() (localctx IParallelContext) {
	localctx = NewParallelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, fsmParserRULE_parallel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(114)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(117)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(120)
		p.Match(fsmParserPARALLEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(fsmParserSTATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(123)
			p.Stereotype()
		}

	}
	{
		p.SetState(126)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4416795115680) != 0 {
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(127)
				p.Region()
			}

		case 2:
			{
				p.SetState(128)
				p.Point()
			}

		case 3:
			{
				p.SetState(129)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
		p.Match(fsmParserT__1)
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

// ISubmachineContext is an interface to support dynamic dispatch.
type ISubmachineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUBMACHINE() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	INITIAL() antlr.TerminalNode
	Stereotype() IStereotypeContext
	AllPoint() []IPointContext
	Point(i int) IPointContext
	AllReference() []IReferenceContext
	Reference(i int) IReferenceContext
	AllEvent() []IEventContext
	Event(i int) IEventContext

	// IsSubmachineContext differentiates from other interfaces.
	IsSubmachineContext()
}

type SubmachineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubmachineContext() *SubmachineContext {
	var p = new(SubmachineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_submachine
	return p
}

func InitEmptySubmachineContext(p *SubmachineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_submachine
}

func (*SubmachineContext) IsSubmachineContext() {}

func NewSubmachineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubmachineContext {
	var p = new(SubmachineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_submachine

	return p
}

func (s *SubmachineContext) GetParser() antlr.Parser { return s.parser }

func (s *SubmachineContext) SUBMACHINE() antlr.TerminalNode {
	return s.GetToken(fsmParserSUBMACHINE, 0)
}

func (s *SubmachineContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *SubmachineContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *SubmachineContext) INITIAL() antlr.TerminalNode {
	return s.GetToken(fsmParserINITIAL, 0)
}

func (s *SubmachineContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *SubmachineContext) AllPoint() []IPointContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPointContext); ok {
			len++
		}
	}

	tst := make([]IPointContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPointContext); ok {
			tst[i] = t.(IPointContext)
			i++
		}
	}

	return tst
}

func (s *SubmachineContext) Point(i int) IPointContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointContext); ok {
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

	return t.(IPointContext)
}

func (s *SubmachineContext) AllReference() []IReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReferenceContext); ok {
			len++
		}
	}

	tst := make([]IReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReferenceContext); ok {
			tst[i] = t.(IReferenceContext)
			i++
		}
	}

	return tst
}

func (s *SubmachineContext) Reference(i int) IReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
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

	return t.(IReferenceContext)
}

func (s *SubmachineContext) AllEvent() []IEventContext {
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

func (s *SubmachineContext) Event(i int) IEventContext {
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

func (s *SubmachineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubmachineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubmachineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitSubmachine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Submachine() (localctx ISubmachineContext) {
	localctx = NewSubmachineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, fsmParserRULE_submachine)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(137)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(140)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(143)
		p.Match(fsmParserSUBMACHINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(145)
			p.Stereotype()
		}

	}
	{
		p.SetState(148)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4416795050144) != 0 {
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(149)
				p.Point()
			}

		case 2:
			{
				p.SetState(150)
				p.Reference()
			}

		case 3:
			{
				p.SetState(151)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(fsmParserT__1)
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

// IFinalContext is an interface to support dynamic dispatch.
type IFinalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FINAL() antlr.TerminalNode
	STATE() antlr.TerminalNode
	Note() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Stereotype() IStereotypeContext

	// IsFinalContext differentiates from other interfaces.
	IsFinalContext()
}

type FinalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFinalContext() *FinalContext {
	var p = new(FinalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_final
	return p
}

func InitEmptyFinalContext(p *FinalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_final
}

func (*FinalContext) IsFinalContext() {}

func NewFinalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FinalContext {
	var p = new(FinalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_final

	return p
}

func (s *FinalContext) GetParser() antlr.Parser { return s.parser }

func (s *FinalContext) FINAL() antlr.TerminalNode {
	return s.GetToken(fsmParserFINAL, 0)
}

func (s *FinalContext) STATE() antlr.TerminalNode {
	return s.GetToken(fsmParserSTATE, 0)
}

func (s *FinalContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *FinalContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *FinalContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *FinalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FinalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FinalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitFinal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Final() (localctx IFinalContext) {
	localctx = NewFinalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, fsmParserRULE_final)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(159)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(162)
		p.Match(fsmParserFINAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(fsmParserSTATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserIdentifier {
		{
			p.SetState(164)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(167)
			p.Stereotype()
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

// ITerminateContext is an interface to support dynamic dispatch.
type ITerminateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TERMINATE() antlr.TerminalNode
	STATE() antlr.TerminalNode
	Note() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Stereotype() IStereotypeContext

	// IsTerminateContext differentiates from other interfaces.
	IsTerminateContext()
}

type TerminateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminateContext() *TerminateContext {
	var p = new(TerminateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_terminate
	return p
}

func InitEmptyTerminateContext(p *TerminateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_terminate
}

func (*TerminateContext) IsTerminateContext() {}

func NewTerminateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminateContext {
	var p = new(TerminateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_terminate

	return p
}

func (s *TerminateContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminateContext) TERMINATE() antlr.TerminalNode {
	return s.GetToken(fsmParserTERMINATE, 0)
}

func (s *TerminateContext) STATE() antlr.TerminalNode {
	return s.GetToken(fsmParserSTATE, 0)
}

func (s *TerminateContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *TerminateContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *TerminateContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *TerminateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitTerminate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Terminate() (localctx ITerminateContext) {
	localctx = NewTerminateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, fsmParserRULE_terminate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(170)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(173)
		p.Match(fsmParserTERMINATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(fsmParserSTATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserIdentifier {
		{
			p.SetState(175)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(178)
			p.Stereotype()
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

// IRegionContext is an interface to support dynamic dispatch.
type IRegionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGION() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	Stereotype() IStereotypeContext
	AllStart_() []IStartContext
	Start_(i int) IStartContext
	AllState() []IStateContext
	State(i int) IStateContext
	AllParallel() []IParallelContext
	Parallel(i int) IParallelContext
	AllSubmachine() []ISubmachineContext
	Submachine(i int) ISubmachineContext
	AllFinal() []IFinalContext
	Final(i int) IFinalContext
	AllTerminate() []ITerminateContext
	Terminate(i int) ITerminateContext
	AllChoice() []IChoiceContext
	Choice(i int) IChoiceContext
	AllJunction() []IJunctionContext
	Junction(i int) IJunctionContext
	AllFork() []IForkContext
	Fork(i int) IForkContext
	AllJoin() []IJoinContext
	Join(i int) IJoinContext
	AllHistory() []IHistoryContext
	History(i int) IHistoryContext

	// IsRegionContext differentiates from other interfaces.
	IsRegionContext()
}

type RegionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegionContext() *RegionContext {
	var p = new(RegionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_region
	return p
}

func InitEmptyRegionContext(p *RegionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_region
}

func (*RegionContext) IsRegionContext() {}

func NewRegionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegionContext {
	var p = new(RegionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_region

	return p
}

func (s *RegionContext) GetParser() antlr.Parser { return s.parser }

func (s *RegionContext) REGION() antlr.TerminalNode {
	return s.GetToken(fsmParserREGION, 0)
}

func (s *RegionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *RegionContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *RegionContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *RegionContext) AllStart_() []IStartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStartContext); ok {
			len++
		}
	}

	tst := make([]IStartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStartContext); ok {
			tst[i] = t.(IStartContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Start_(i int) IStartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartContext); ok {
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

	return t.(IStartContext)
}

func (s *RegionContext) AllState() []IStateContext {
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

func (s *RegionContext) State(i int) IStateContext {
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

func (s *RegionContext) AllParallel() []IParallelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParallelContext); ok {
			len++
		}
	}

	tst := make([]IParallelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParallelContext); ok {
			tst[i] = t.(IParallelContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Parallel(i int) IParallelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParallelContext); ok {
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

	return t.(IParallelContext)
}

func (s *RegionContext) AllSubmachine() []ISubmachineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubmachineContext); ok {
			len++
		}
	}

	tst := make([]ISubmachineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubmachineContext); ok {
			tst[i] = t.(ISubmachineContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Submachine(i int) ISubmachineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubmachineContext); ok {
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

	return t.(ISubmachineContext)
}

func (s *RegionContext) AllFinal() []IFinalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFinalContext); ok {
			len++
		}
	}

	tst := make([]IFinalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFinalContext); ok {
			tst[i] = t.(IFinalContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Final(i int) IFinalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFinalContext); ok {
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

	return t.(IFinalContext)
}

func (s *RegionContext) AllTerminate() []ITerminateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminateContext); ok {
			len++
		}
	}

	tst := make([]ITerminateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminateContext); ok {
			tst[i] = t.(ITerminateContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Terminate(i int) ITerminateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminateContext); ok {
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

	return t.(ITerminateContext)
}

func (s *RegionContext) AllChoice() []IChoiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChoiceContext); ok {
			len++
		}
	}

	tst := make([]IChoiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChoiceContext); ok {
			tst[i] = t.(IChoiceContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Choice(i int) IChoiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChoiceContext); ok {
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

	return t.(IChoiceContext)
}

func (s *RegionContext) AllJunction() []IJunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJunctionContext); ok {
			len++
		}
	}

	tst := make([]IJunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJunctionContext); ok {
			tst[i] = t.(IJunctionContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Junction(i int) IJunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJunctionContext); ok {
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

	return t.(IJunctionContext)
}

func (s *RegionContext) AllFork() []IForkContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForkContext); ok {
			len++
		}
	}

	tst := make([]IForkContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForkContext); ok {
			tst[i] = t.(IForkContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Fork(i int) IForkContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForkContext); ok {
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

	return t.(IForkContext)
}

func (s *RegionContext) AllJoin() []IJoinContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinContext); ok {
			len++
		}
	}

	tst := make([]IJoinContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinContext); ok {
			tst[i] = t.(IJoinContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) Join(i int) IJoinContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinContext); ok {
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

	return t.(IJoinContext)
}

func (s *RegionContext) AllHistory() []IHistoryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHistoryContext); ok {
			len++
		}
	}

	tst := make([]IHistoryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHistoryContext); ok {
			tst[i] = t.(IHistoryContext)
			i++
		}
	}

	return tst
}

func (s *RegionContext) History(i int) IHistoryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHistoryContext); ok {
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

	return t.(IHistoryContext)
}

func (s *RegionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitRegion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Region() (localctx IRegionContext) {
	localctx = NewRegionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, fsmParserRULE_region)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(181)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(184)
		p.Match(fsmParserREGION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(186)
			p.Stereotype()
		}

	}
	{
		p.SetState(189)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5226983514112) != 0 {
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(190)
				p.Start_()
			}

		case 2:
			{
				p.SetState(191)
				p.State()
			}

		case 3:
			{
				p.SetState(192)
				p.Parallel()
			}

		case 4:
			{
				p.SetState(193)
				p.Submachine()
			}

		case 5:
			{
				p.SetState(194)
				p.Final()
			}

		case 6:
			{
				p.SetState(195)
				p.Terminate()
			}

		case 7:
			{
				p.SetState(196)
				p.Choice()
			}

		case 8:
			{
				p.SetState(197)
				p.Junction()
			}

		case 9:
			{
				p.SetState(198)
				p.Fork()
			}

		case 10:
			{
				p.SetState(199)
				p.Join()
			}

		case 11:
			{
				p.SetState(200)
				p.History()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(206)
		p.Match(fsmParserT__1)
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

// IChoiceContext is an interface to support dynamic dispatch.
type IChoiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHOICE() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	AllBranch() []IBranchContext
	Branch(i int) IBranchContext
	Note() antlr.TerminalNode
	INITIAL() antlr.TerminalNode
	Stereotype() IStereotypeContext

	// IsChoiceContext differentiates from other interfaces.
	IsChoiceContext()
}

type ChoiceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChoiceContext() *ChoiceContext {
	var p = new(ChoiceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_choice
	return p
}

func InitEmptyChoiceContext(p *ChoiceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_choice
}

func (*ChoiceContext) IsChoiceContext() {}

func NewChoiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceContext {
	var p = new(ChoiceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_choice

	return p
}

func (s *ChoiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceContext) CHOICE() antlr.TerminalNode {
	return s.GetToken(fsmParserCHOICE, 0)
}

func (s *ChoiceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *ChoiceContext) AllBranch() []IBranchContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBranchContext); ok {
			len++
		}
	}

	tst := make([]IBranchContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBranchContext); ok {
			tst[i] = t.(IBranchContext)
			i++
		}
	}

	return tst
}

func (s *ChoiceContext) Branch(i int) IBranchContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranchContext); ok {
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

	return t.(IBranchContext)
}

func (s *ChoiceContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *ChoiceContext) INITIAL() antlr.TerminalNode {
	return s.GetToken(fsmParserINITIAL, 0)
}

func (s *ChoiceContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *ChoiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChoiceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitChoice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Choice() (localctx IChoiceContext) {
	localctx = NewChoiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, fsmParserRULE_choice)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(208)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(211)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(214)
		p.Match(fsmParserCHOICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(216)
			p.Stereotype()
		}

	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__0:
		{
			p.SetState(219)
			p.Match(fsmParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4399120253088) != 0 {
			{
				p.SetState(220)
				p.Branch()
			}

			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(226)
			p.Match(fsmParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__4, fsmParserT__6, fsmParserGOTO, fsmParserNote:
		{
			p.SetState(227)
			p.Branch()
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

// IJunctionContext is an interface to support dynamic dispatch.
type IJunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JUNCTION() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	AllBranch() []IBranchContext
	Branch(i int) IBranchContext
	Note() antlr.TerminalNode
	INITIAL() antlr.TerminalNode
	Stereotype() IStereotypeContext

	// IsJunctionContext differentiates from other interfaces.
	IsJunctionContext()
}

type JunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJunctionContext() *JunctionContext {
	var p = new(JunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_junction
	return p
}

func InitEmptyJunctionContext(p *JunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_junction
}

func (*JunctionContext) IsJunctionContext() {}

func NewJunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JunctionContext {
	var p = new(JunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_junction

	return p
}

func (s *JunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *JunctionContext) JUNCTION() antlr.TerminalNode {
	return s.GetToken(fsmParserJUNCTION, 0)
}

func (s *JunctionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *JunctionContext) AllBranch() []IBranchContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBranchContext); ok {
			len++
		}
	}

	tst := make([]IBranchContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBranchContext); ok {
			tst[i] = t.(IBranchContext)
			i++
		}
	}

	return tst
}

func (s *JunctionContext) Branch(i int) IBranchContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranchContext); ok {
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

	return t.(IBranchContext)
}

func (s *JunctionContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *JunctionContext) INITIAL() antlr.TerminalNode {
	return s.GetToken(fsmParserINITIAL, 0)
}

func (s *JunctionContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *JunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitJunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Junction() (localctx IJunctionContext) {
	localctx = NewJunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, fsmParserRULE_junction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(230)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(233)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(236)
		p.Match(fsmParserJUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(238)
			p.Stereotype()
		}

	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__0:
		{
			p.SetState(241)
			p.Match(fsmParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4399120253088) != 0 {
			{
				p.SetState(242)
				p.Branch()
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(248)
			p.Match(fsmParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__4, fsmParserT__6, fsmParserGOTO, fsmParserNote:
		{
			p.SetState(249)
			p.Branch()
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

// IForkContext is an interface to support dynamic dispatch.
type IForkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FORK() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	AllNote() []antlr.TerminalNode
	Note(i int) antlr.TerminalNode
	Stereotype() IStereotypeContext
	AllGoto_() []IGotoContext
	Goto_(i int) IGotoContext
	AllActions() []IActionsContext
	Actions(i int) IActionsContext

	// IsForkContext differentiates from other interfaces.
	IsForkContext()
}

type ForkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForkContext() *ForkContext {
	var p = new(ForkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_fork
	return p
}

func InitEmptyForkContext(p *ForkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_fork
}

func (*ForkContext) IsForkContext() {}

func NewForkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForkContext {
	var p = new(ForkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_fork

	return p
}

func (s *ForkContext) GetParser() antlr.Parser { return s.parser }

func (s *ForkContext) FORK() antlr.TerminalNode {
	return s.GetToken(fsmParserFORK, 0)
}

func (s *ForkContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *ForkContext) AllNote() []antlr.TerminalNode {
	return s.GetTokens(fsmParserNote)
}

func (s *ForkContext) Note(i int) antlr.TerminalNode {
	return s.GetToken(fsmParserNote, i)
}

func (s *ForkContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *ForkContext) AllGoto_() []IGotoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGotoContext); ok {
			len++
		}
	}

	tst := make([]IGotoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGotoContext); ok {
			tst[i] = t.(IGotoContext)
			i++
		}
	}

	return tst
}

func (s *ForkContext) Goto_(i int) IGotoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGotoContext); ok {
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

	return t.(IGotoContext)
}

func (s *ForkContext) AllActions() []IActionsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionsContext); ok {
			len++
		}
	}

	tst := make([]IActionsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionsContext); ok {
			tst[i] = t.(IActionsContext)
			i++
		}
	}

	return tst
}

func (s *ForkContext) Actions(i int) IActionsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionsContext); ok {
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

	return t.(IActionsContext)
}

func (s *ForkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitFork(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Fork() (localctx IForkContext) {
	localctx = NewForkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, fsmParserRULE_fork)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(252)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(255)
		p.Match(fsmParserFORK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(257)
			p.Stereotype()
		}

	}
	{
		p.SetState(260)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4399120253056) != 0 {
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserNote {
			{
				p.SetState(261)
				p.Match(fsmParserNote)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__6 {
			{
				p.SetState(264)
				p.Actions()
			}

		}
		{
			p.SetState(267)
			p.Goto_()
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)
		p.Match(fsmParserT__1)
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

// IJoinContext is an interface to support dynamic dispatch.
type IJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JOIN() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Goto_() IGotoContext
	Note() antlr.TerminalNode
	Stereotype() IStereotypeContext
	Actions() IActionsContext

	// IsJoinContext differentiates from other interfaces.
	IsJoinContext()
}

type JoinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinContext() *JoinContext {
	var p = new(JoinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_join
	return p
}

func InitEmptyJoinContext(p *JoinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_join
}

func (*JoinContext) IsJoinContext() {}

func NewJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinContext {
	var p = new(JoinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_join

	return p
}

func (s *JoinContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinContext) JOIN() antlr.TerminalNode {
	return s.GetToken(fsmParserJOIN, 0)
}

func (s *JoinContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *JoinContext) Goto_() IGotoContext {
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

func (s *JoinContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *JoinContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *JoinContext) Actions() IActionsContext {
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

func (s *JoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Join() (localctx IJoinContext) {
	localctx = NewJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, fsmParserRULE_join)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(275)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(278)
		p.Match(fsmParserJOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(280)
			p.Stereotype()
		}

	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__6 {
		{
			p.SetState(283)
			p.Actions()
		}

	}
	{
		p.SetState(286)
		p.Goto_()
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

// IPointContext is an interface to support dynamic dispatch.
type IPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKind returns the kind token.
	GetKind() antlr.Token

	// SetKind sets the kind token.
	SetKind(antlr.Token)

	// Getter signatures
	POINT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Goto_() IGotoContext
	ENTRY() antlr.TerminalNode
	EXIT() antlr.TerminalNode
	Note() antlr.TerminalNode
	Stereotype() IStereotypeContext
	Actions() IActionsContext

	// IsPointContext differentiates from other interfaces.
	IsPointContext()
}

type PointContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	kind   antlr.Token
}

func NewEmptyPointContext() *PointContext {
	var p = new(PointContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_point
	return p
}

func InitEmptyPointContext(p *PointContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_point
}

func (*PointContext) IsPointContext() {}

func NewPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointContext {
	var p = new(PointContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_point

	return p
}

func (s *PointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointContext) GetKind() antlr.Token { return s.kind }

func (s *PointContext) SetKind(v antlr.Token) { s.kind = v }

func (s *PointContext) POINT() antlr.TerminalNode {
	return s.GetToken(fsmParserPOINT, 0)
}

func (s *PointContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *PointContext) Goto_() IGotoContext {
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

func (s *PointContext) ENTRY() antlr.TerminalNode {
	return s.GetToken(fsmParserENTRY, 0)
}

func (s *PointContext) EXIT() antlr.TerminalNode {
	return s.GetToken(fsmParserEXIT, 0)
}

func (s *PointContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *PointContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *PointContext) Actions() IActionsContext {
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

func (s *PointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitPoint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Point() (localctx IPointContext) {
	localctx = NewPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, fsmParserRULE_point)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(288)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(291)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PointContext).kind = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == fsmParserENTRY || _la == fsmParserEXIT) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PointContext).kind = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(292)
		p.Match(fsmParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(294)
			p.Stereotype()
		}

	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__6 {
		{
			p.SetState(297)
			p.Actions()
		}

	}
	{
		p.SetState(300)
		p.Goto_()
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

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENTRY() antlr.TerminalNode
	POINT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	Stereotype() IStereotypeContext

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) ENTRY() antlr.TerminalNode {
	return s.GetToken(fsmParserENTRY, 0)
}

func (s *ReferenceContext) POINT() antlr.TerminalNode {
	return s.GetToken(fsmParserPOINT, 0)
}

func (s *ReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *ReferenceContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *ReferenceContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, fsmParserRULE_reference)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(302)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(305)
		p.Match(fsmParserENTRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(fsmParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(308)
			p.Stereotype()
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

// IHistoryContext is an interface to support dynamic dispatch.
type IHistoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKind returns the kind token.
	GetKind() antlr.Token

	// SetKind sets the kind token.
	SetKind(antlr.Token)

	// Getter signatures
	H() antlr.TerminalNode
	H_DEEP() antlr.TerminalNode
	Note() antlr.TerminalNode
	Stereotype() IStereotypeContext
	Goto_() IGotoContext
	Actions() IActionsContext

	// IsHistoryContext differentiates from other interfaces.
	IsHistoryContext()
}

type HistoryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	kind   antlr.Token
}

func NewEmptyHistoryContext() *HistoryContext {
	var p = new(HistoryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_history
	return p
}

func InitEmptyHistoryContext(p *HistoryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_history
}

func (*HistoryContext) IsHistoryContext() {}

func NewHistoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HistoryContext {
	var p = new(HistoryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_history

	return p
}

func (s *HistoryContext) GetParser() antlr.Parser { return s.parser }

func (s *HistoryContext) GetKind() antlr.Token { return s.kind }

func (s *HistoryContext) SetKind(v antlr.Token) { s.kind = v }

func (s *HistoryContext) H() antlr.TerminalNode {
	return s.GetToken(fsmParserH, 0)
}

func (s *HistoryContext) H_DEEP() antlr.TerminalNode {
	return s.GetToken(fsmParserH_DEEP, 0)
}

func (s *HistoryContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *HistoryContext) Stereotype() IStereotypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStereotypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStereotypeContext)
}

func (s *HistoryContext) Goto_() IGotoContext {
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

func (s *HistoryContext) Actions() IActionsContext {
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

func (s *HistoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HistoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HistoryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitHistory(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) History() (localctx IHistoryContext) {
	localctx = NewHistoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, fsmParserRULE_history)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(311)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(314)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*HistoryContext).kind = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == fsmParserH || _la == fsmParserH_DEEP) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*HistoryContext).kind = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(315)
			p.Stereotype()
		}

	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__6 {
			{
				p.SetState(318)
				p.Actions()
			}

		}
		{
			p.SetState(321)
			p.Goto_()
		}

	} else if p.HasError() { // JIM
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

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INITIAL() antlr.TerminalNode
	GOTO() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Actions() IActionsContext

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) INITIAL() antlr.TerminalNode {
	return s.GetToken(fsmParserINITIAL, 0)
}

func (s *StartContext) GOTO() antlr.TerminalNode {
	return s.GetToken(fsmParserGOTO, 0)
}

func (s *StartContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *StartContext) Actions() IActionsContext {
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

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, fsmParserRULE_start)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(fsmParserINITIAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__6 {
		{
			p.SetState(325)
			p.Actions()
		}

	}
	{
		p.SetState(328)
		p.Match(fsmParserGOTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Match(fsmParserIdentifier)
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

// IStereotypeContext is an interface to support dynamic dispatch.
type IStereotypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsStereotypeContext differentiates from other interfaces.
	IsStereotypeContext()
}

type StereotypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStereotypeContext() *StereotypeContext {
	var p = new(StereotypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_stereotype
	return p
}

func InitEmptyStereotypeContext(p *StereotypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_stereotype
}

func (*StereotypeContext) IsStereotypeContext() {}

func NewStereotypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StereotypeContext {
	var p = new(StereotypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_stereotype

	return p
}

func (s *StereotypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StereotypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *StereotypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StereotypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StereotypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitStereotype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Stereotype() (localctx IStereotypeContext) {
	localctx = NewStereotypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, fsmParserRULE_stereotype)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(331)
		p.Match(fsmParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(332)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.Match(fsmParserT__3)
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

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Goto_() IGotoContext
	Note() antlr.TerminalNode
	ELSE() antlr.TerminalNode
	Guard() IGuardContext
	Actions() IActionsContext

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_branch
	return p
}

func InitEmptyBranchContext(p *BranchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_branch
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) Goto_() IGotoContext {
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

func (s *BranchContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *BranchContext) ELSE() antlr.TerminalNode {
	return s.GetToken(fsmParserELSE, 0)
}

func (s *BranchContext) Guard() IGuardContext {
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

func (s *BranchContext) Actions() IActionsContext {
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

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Branch() (localctx IBranchContext) {
	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, fsmParserRULE_branch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(335)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(338)
			p.Match(fsmParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(fsmParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(fsmParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(341)
			p.Guard()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__6 {
		{
			p.SetState(344)
			p.Actions()
		}

	}
	{
		p.SetState(347)
		p.Goto_()
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
	ON() antlr.TerminalNode
	DEFER() antlr.TerminalNode
	INVARIANT() antlr.TerminalNode
	Guard() IGuardContext
	Trigger() ITriggerContext
	Goto_() IGotoContext
	Note() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	ENTRY() antlr.TerminalNode
	EXIT() antlr.TerminalNode
	DO() antlr.TerminalNode

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

func (s *EventContext) ON() antlr.TerminalNode {
	return s.GetToken(fsmParserON, 0)
}

func (s *EventContext) DEFER() antlr.TerminalNode {
	return s.GetToken(fsmParserDEFER, 0)
}

func (s *EventContext) INVARIANT() antlr.TerminalNode {
	return s.GetToken(fsmParserINVARIANT, 0)
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

func (s *EventContext) Trigger() ITriggerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriggerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriggerContext)
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

func (s *EventContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *EventContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *EventContext) ENTRY() antlr.TerminalNode {
	return s.GetToken(fsmParserENTRY, 0)
}

func (s *EventContext) EXIT() antlr.TerminalNode {
	return s.GetToken(fsmParserEXIT, 0)
}

func (s *EventContext) DO() antlr.TerminalNode {
	return s.GetToken(fsmParserDO, 0)
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
	p.EnterRule(localctx, 34, fsmParserRULE_event)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(349)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(352)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*EventContext).name = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&92274688) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*EventContext).name = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(353)
			p.Actions()
		}

	case 2:
		{
			p.SetState(354)
			p.Match(fsmParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*EventContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(fsmParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(fsmParserDEFER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(358)
			p.Match(fsmParserINVARIANT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)
			p.Guard()
		}

	case 4:
		{
			p.SetState(360)
			p.Trigger()
		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__4 {
			{
				p.SetState(361)
				p.Guard()
			}

		}
		{
			p.SetState(364)
			p.Actions()
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(365)
				p.Goto_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		{
			p.SetState(368)
			p.Trigger()
		}
		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__4 {
			{
				p.SetState(369)
				p.Guard()
			}

		}
		{
			p.SetState(372)
			p.Goto_()
		}

	case 6:
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__4 {
			{
				p.SetState(374)
				p.Guard()
			}

		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__6 {
			{
				p.SetState(377)
				p.Actions()
			}

		}
		{
			p.SetState(380)
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

// ITriggerContext is an interface to support dynamic dispatch.
type ITriggerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetDelay returns the delay token.
	GetDelay() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetDelay sets the delay token.
	SetDelay(antlr.Token)

	// Getter signatures
	ON() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	AFTER() antlr.TerminalNode
	Duration() antlr.TerminalNode

	// IsTriggerContext differentiates from other interfaces.
	IsTriggerContext()
}

type TriggerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	delay  antlr.Token
}

func NewEmptyTriggerContext() *TriggerContext {
	var p = new(TriggerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_trigger
	return p
}

func InitEmptyTriggerContext(p *TriggerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_trigger
}

func (*TriggerContext) IsTriggerContext() {}

func NewTriggerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriggerContext {
	var p = new(TriggerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_trigger

	return p
}

func (s *TriggerContext) GetParser() antlr.Parser { return s.parser }

func (s *TriggerContext) GetName() antlr.Token { return s.name }

func (s *TriggerContext) GetDelay() antlr.Token { return s.delay }

func (s *TriggerContext) SetName(v antlr.Token) { s.name = v }

func (s *TriggerContext) SetDelay(v antlr.Token) { s.delay = v }

func (s *TriggerContext) ON() antlr.TerminalNode {
	return s.GetToken(fsmParserON, 0)
}

func (s *TriggerContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *TriggerContext) AFTER() antlr.TerminalNode {
	return s.GetToken(fsmParserAFTER, 0)
}

func (s *TriggerContext) Duration() antlr.TerminalNode {
	return s.GetToken(fsmParserDuration, 0)
}

func (s *TriggerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriggerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriggerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitTrigger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Trigger() (localctx ITriggerContext) {
	localctx = NewTriggerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, fsmParserRULE_trigger)
	var _la int

	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.Match(fsmParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*TriggerContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserAFTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.Match(fsmParserAFTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(fsmParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TriggerContext).delay = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == fsmParserDuration || _la == fsmParserIdentifier) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TriggerContext).delay = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(388)
			p.Match(fsmParserT__8)
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
	p.EnterRule(localctx, 38, fsmParserRULE_actions)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(fsmParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
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
	p.EnterRule(localctx, 40, fsmParserRULE_identifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__9 {
		{
			p.SetState(395)
			p.Match(fsmParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(401)
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
	p.EnterRule(localctx, 42, fsmParserRULE_guard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Match(fsmParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Expression()
	}
	{
		p.SetState(404)
		p.Match(fsmParserT__5)
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
	p.EnterRule(localctx, 44, fsmParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
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
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

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

func (s *Or_expressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(fsmParserOR)
}

func (s *Or_expressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(fsmParserOR, i)
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
	p.EnterRule(localctx, 46, fsmParserRULE_or_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(408)
		p.And_expression()
	}
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserOR {
		{
			p.SetState(409)
			p.Match(fsmParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.And_expression()
		}

		p.SetState(415)
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
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

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

func (s *And_expressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(fsmParserAND)
}

func (s *And_expressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(fsmParserAND, i)
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
	p.EnterRule(localctx, 48, fsmParserRULE_and_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.Not_expression()
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserAND {
		{
			p.SetState(417)
			p.Match(fsmParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Not_expression()
		}

		p.SetState(423)
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
	NOT() antlr.TerminalNode

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

func (s *Not_expressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(fsmParserNOT, 0)
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
	p.EnterRule(localctx, 50, fsmParserRULE_not_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNOT {
		{
			p.SetState(424)
			p.Match(fsmParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(427)
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
	p.EnterRule(localctx, 52, fsmParserRULE_single_expression)
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(429)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(430)
			p.Match(fsmParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Expression()
		}
		{
			p.SetState(432)
			p.Match(fsmParserT__8)
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
	GOTO() antlr.TerminalNode
	FINAL() antlr.TerminalNode
	TERMINATE() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	LOCAL() antlr.TerminalNode
	H() antlr.TerminalNode
	H_DEEP() antlr.TerminalNode

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

func (s *GotoContext) GOTO() antlr.TerminalNode {
	return s.GetToken(fsmParserGOTO, 0)
}

func (s *GotoContext) FINAL() antlr.TerminalNode {
	return s.GetToken(fsmParserFINAL, 0)
}

func (s *GotoContext) TERMINATE() antlr.TerminalNode {
	return s.GetToken(fsmParserTERMINATE, 0)
}

func (s *GotoContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *GotoContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(fsmParserLOCAL, 0)
}

func (s *GotoContext) H() antlr.TerminalNode {
	return s.GetToken(fsmParserH, 0)
}

func (s *GotoContext) H_DEEP() antlr.TerminalNode {
	return s.GetToken(fsmParserH_DEEP, 0)
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
	p.EnterRule(localctx, 54, fsmParserRULE_goto)
	var _la int

	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 75, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(436)
			p.Match(fsmParserGOTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			_la = p.GetTokenStream().LA(1)

			if !(_la == fsmParserFINAL || _la == fsmParserTERMINATE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.Match(fsmParserGOTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserLOCAL {
			{
				p.SetState(439)
				p.Match(fsmParserLOCAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(442)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(443)
			p.Match(fsmParserGOTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserLOCAL {
			{
				p.SetState(444)
				p.Match(fsmParserLOCAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserIdentifier {
			{
				p.SetState(447)
				p.Match(fsmParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(448)
				p.Match(fsmParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(451)
			_la = p.GetTokenStream().LA(1)

			if !(_la == fsmParserH || _la == fsmParserH_DEEP) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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
