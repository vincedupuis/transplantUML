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
		"", "'{'", "'}'", "':'", "'<<'", "'>>'", "'['", "']'", "'/'", "'('",
		"')'", "','", "'.'", "'fsm'", "'state'", "'parallel'", "'submachine'",
		"'region'", "'final'", "'terminate'", "'choice'", "'junction'", "'fork'",
		"'join'", "'entry'", "'exit'", "'point'", "'do'", "'on'", "'after'",
		"'else'", "'goto'", "'local'", "'initial'", "'defer'", "'invariant'",
		"'and'", "'or'", "'not'", "'H'", "'H*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "FSM", "STATE",
		"PARALLEL", "SUBMACHINE", "REGION", "FINAL", "TERMINATE", "CHOICE",
		"JUNCTION", "FORK", "JOIN", "ENTRY", "EXIT", "POINT", "DO", "ON", "AFTER",
		"ELSE", "GOTO", "LOCAL", "INITIAL", "DEFER", "INVARIANT", "AND", "OR",
		"NOT", "H", "H_DEEP", "Duration", "Identifier", "Note", "Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"fsm", "state", "parallel", "submachine", "final", "terminate", "region",
		"choice", "junction", "fork", "join", "point", "reference", "outlet",
		"history", "start", "stereotype", "branch", "event", "trigger", "actions",
		"identifiers", "guard", "expression", "or_expression", "and_expression",
		"not_expression", "single_expression", "goto",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 471, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 3, 0, 60, 8, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 5, 0, 78, 8, 0, 10, 0, 12, 0, 81, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1,
		3, 1, 87, 8, 1, 1, 1, 3, 1, 90, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 95, 8, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 111, 8, 1, 10, 1, 12, 1, 114, 9, 1, 1, 1, 1, 1, 1, 2,
		3, 2, 119, 8, 2, 1, 2, 3, 2, 122, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 128,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 134, 8, 2, 10, 2, 12, 2, 137, 9, 2,
		1, 2, 1, 2, 1, 3, 3, 3, 142, 8, 3, 1, 3, 3, 3, 145, 8, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 151, 8, 3, 1, 3, 3, 3, 154, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 160, 8, 3, 10, 3, 12, 3, 163, 9, 3, 1, 3, 1, 3, 1, 4, 3, 4, 168,
		8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 173, 8, 4, 1, 4, 3, 4, 176, 8, 4, 1, 5, 3,
		5, 179, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 184, 8, 5, 1, 5, 3, 5, 187, 8, 5,
		1, 6, 3, 6, 190, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 195, 8, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 209, 8,
		6, 10, 6, 12, 6, 212, 9, 6, 1, 6, 1, 6, 1, 7, 3, 7, 217, 8, 7, 1, 7, 3,
		7, 220, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 225, 8, 7, 1, 7, 1, 7, 5, 7, 229,
		8, 7, 10, 7, 12, 7, 232, 9, 7, 1, 7, 1, 7, 3, 7, 236, 8, 7, 1, 8, 3, 8,
		239, 8, 8, 1, 8, 3, 8, 242, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 247, 8, 8, 1,
		8, 1, 8, 5, 8, 251, 8, 8, 10, 8, 12, 8, 254, 9, 8, 1, 8, 1, 8, 3, 8, 258,
		8, 8, 1, 9, 3, 9, 261, 8, 9, 1, 9, 1, 9, 1, 9, 3, 9, 266, 8, 9, 1, 9, 1,
		9, 3, 9, 270, 8, 9, 1, 9, 3, 9, 273, 8, 9, 1, 9, 5, 9, 276, 8, 9, 10, 9,
		12, 9, 279, 9, 9, 1, 9, 1, 9, 1, 10, 3, 10, 284, 8, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 289, 8, 10, 1, 10, 3, 10, 292, 8, 10, 1, 10, 1, 10, 1, 11, 3,
		11, 297, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 303, 8, 11, 1, 11, 3,
		11, 306, 8, 11, 1, 11, 1, 11, 1, 12, 3, 12, 311, 8, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 317, 8, 12, 1, 13, 3, 13, 320, 8, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 326, 8, 13, 1, 14, 3, 14, 329, 8, 14, 1, 14, 1, 14, 3,
		14, 333, 8, 14, 1, 14, 3, 14, 336, 8, 14, 1, 14, 3, 14, 339, 8, 14, 1,
		15, 1, 15, 3, 15, 343, 8, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 3, 17, 353, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 359,
		8, 17, 1, 17, 3, 17, 362, 8, 17, 1, 17, 1, 17, 1, 18, 3, 18, 367, 8, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 379, 8, 18, 1, 18, 1, 18, 3, 18, 383, 8, 18, 1, 18, 1, 18, 3, 18, 387,
		8, 18, 1, 18, 1, 18, 1, 18, 3, 18, 392, 8, 18, 1, 18, 3, 18, 395, 8, 18,
		1, 18, 3, 18, 398, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3,
		19, 406, 8, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 414, 8,
		21, 10, 21, 12, 21, 417, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 5, 24, 428, 8, 24, 10, 24, 12, 24, 431, 9, 24, 1,
		25, 1, 25, 1, 25, 5, 25, 436, 8, 25, 10, 25, 12, 25, 439, 9, 25, 1, 26,
		3, 26, 442, 8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3,
		27, 451, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 457, 8, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 462, 8, 28, 1, 28, 1, 28, 3, 28, 466, 8, 28, 1, 28, 3,
		28, 469, 8, 28, 1, 28, 0, 0, 29, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		0, 5, 1, 0, 24, 25, 1, 0, 39, 40, 2, 0, 24, 25, 27, 27, 1, 0, 41, 42, 1,
		0, 18, 19, 559, 0, 59, 1, 0, 0, 0, 2, 86, 1, 0, 0, 0, 4, 118, 1, 0, 0,
		0, 6, 141, 1, 0, 0, 0, 8, 167, 1, 0, 0, 0, 10, 178, 1, 0, 0, 0, 12, 189,
		1, 0, 0, 0, 14, 216, 1, 0, 0, 0, 16, 238, 1, 0, 0, 0, 18, 260, 1, 0, 0,
		0, 20, 283, 1, 0, 0, 0, 22, 296, 1, 0, 0, 0, 24, 310, 1, 0, 0, 0, 26, 319,
		1, 0, 0, 0, 28, 328, 1, 0, 0, 0, 30, 340, 1, 0, 0, 0, 32, 347, 1, 0, 0,
		0, 34, 352, 1, 0, 0, 0, 36, 366, 1, 0, 0, 0, 38, 405, 1, 0, 0, 0, 40, 407,
		1, 0, 0, 0, 42, 410, 1, 0, 0, 0, 44, 418, 1, 0, 0, 0, 46, 422, 1, 0, 0,
		0, 48, 424, 1, 0, 0, 0, 50, 432, 1, 0, 0, 0, 52, 441, 1, 0, 0, 0, 54, 450,
		1, 0, 0, 0, 56, 468, 1, 0, 0, 0, 58, 60, 5, 43, 0, 0, 59, 58, 1, 0, 0,
		0, 59, 60, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 5, 13, 0, 0, 62, 63,
		5, 42, 0, 0, 63, 79, 5, 1, 0, 0, 64, 78, 3, 30, 15, 0, 65, 78, 3, 2, 1,
		0, 66, 78, 3, 4, 2, 0, 67, 78, 3, 6, 3, 0, 68, 78, 3, 8, 4, 0, 69, 78,
		3, 10, 5, 0, 70, 78, 3, 14, 7, 0, 71, 78, 3, 16, 8, 0, 72, 78, 3, 18, 9,
		0, 73, 78, 3, 20, 10, 0, 74, 78, 3, 22, 11, 0, 75, 78, 3, 26, 13, 0, 76,
		78, 3, 36, 18, 0, 77, 64, 1, 0, 0, 0, 77, 65, 1, 0, 0, 0, 77, 66, 1, 0,
		0, 0, 77, 67, 1, 0, 0, 0, 77, 68, 1, 0, 0, 0, 77, 69, 1, 0, 0, 0, 77, 70,
		1, 0, 0, 0, 77, 71, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 77, 73, 1, 0, 0, 0,
		77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81,
		79, 1, 0, 0, 0, 82, 83, 5, 2, 0, 0, 83, 84, 5, 0, 0, 1, 84, 1, 1, 0, 0,
		0, 85, 87, 5, 43, 0, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 89,
		1, 0, 0, 0, 88, 90, 5, 33, 0, 0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0,
		90, 91, 1, 0, 0, 0, 91, 92, 5, 14, 0, 0, 92, 94, 5, 42, 0, 0, 93, 95, 3,
		32, 16, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 112, 5, 1, 0, 0, 97, 111, 3, 30, 15, 0, 98, 111, 3, 2, 1, 0, 99, 111,
		3, 4, 2, 0, 100, 111, 3, 6, 3, 0, 101, 111, 3, 8, 4, 0, 102, 111, 3, 10,
		5, 0, 103, 111, 3, 14, 7, 0, 104, 111, 3, 16, 8, 0, 105, 111, 3, 18, 9,
		0, 106, 111, 3, 20, 10, 0, 107, 111, 3, 22, 11, 0, 108, 111, 3, 28, 14,
		0, 109, 111, 3, 36, 18, 0, 110, 97, 1, 0, 0, 0, 110, 98, 1, 0, 0, 0, 110,
		99, 1, 0, 0, 0, 110, 100, 1, 0, 0, 0, 110, 101, 1, 0, 0, 0, 110, 102, 1,
		0, 0, 0, 110, 103, 1, 0, 0, 0, 110, 104, 1, 0, 0, 0, 110, 105, 1, 0, 0,
		0, 110, 106, 1, 0, 0, 0, 110, 107, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110,
		109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113,
		1, 0, 0, 0, 113, 115, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 2,
		0, 0, 116, 3, 1, 0, 0, 0, 117, 119, 5, 43, 0, 0, 118, 117, 1, 0, 0, 0,
		118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120, 122, 5, 33, 0, 0, 121,
		120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124,
		5, 15, 0, 0, 124, 125, 5, 14, 0, 0, 125, 127, 5, 42, 0, 0, 126, 128, 3,
		32, 16, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0,
		0, 0, 129, 135, 5, 1, 0, 0, 130, 134, 3, 12, 6, 0, 131, 134, 3, 22, 11,
		0, 132, 134, 3, 36, 18, 0, 133, 130, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0,
		133, 132, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135,
		136, 1, 0, 0, 0, 136, 138, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139,
		5, 2, 0, 0, 139, 5, 1, 0, 0, 0, 140, 142, 5, 43, 0, 0, 141, 140, 1, 0,
		0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143, 145, 5, 33, 0, 0,
		144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		147, 5, 16, 0, 0, 147, 150, 5, 42, 0, 0, 148, 149, 5, 3, 0, 0, 149, 151,
		5, 42, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0,
		0, 0, 152, 154, 3, 32, 16, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0,
		0, 154, 155, 1, 0, 0, 0, 155, 161, 5, 1, 0, 0, 156, 160, 3, 22, 11, 0,
		157, 160, 3, 24, 12, 0, 158, 160, 3, 36, 18, 0, 159, 156, 1, 0, 0, 0, 159,
		157, 1, 0, 0, 0, 159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159,
		1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0,
		0, 0, 164, 165, 5, 2, 0, 0, 165, 7, 1, 0, 0, 0, 166, 168, 5, 43, 0, 0,
		167, 166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169,
		170, 5, 18, 0, 0, 170, 172, 5, 14, 0, 0, 171, 173, 5, 42, 0, 0, 172, 171,
		1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 175, 1, 0, 0, 0, 174, 176, 3, 32,
		16, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 9, 1, 0, 0, 0,
		177, 179, 5, 43, 0, 0, 178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 181, 5, 19, 0, 0, 181, 183, 5, 14, 0, 0, 182, 184,
		5, 42, 0, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 1, 0,
		0, 0, 185, 187, 3, 32, 16, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0,
		0, 187, 11, 1, 0, 0, 0, 188, 190, 5, 43, 0, 0, 189, 188, 1, 0, 0, 0, 189,
		190, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 192, 5, 17, 0, 0, 192, 194,
		5, 42, 0, 0, 193, 195, 3, 32, 16, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1,
		0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 210, 5, 1, 0, 0, 197, 209, 3, 30, 15,
		0, 198, 209, 3, 2, 1, 0, 199, 209, 3, 4, 2, 0, 200, 209, 3, 6, 3, 0, 201,
		209, 3, 8, 4, 0, 202, 209, 3, 10, 5, 0, 203, 209, 3, 14, 7, 0, 204, 209,
		3, 16, 8, 0, 205, 209, 3, 18, 9, 0, 206, 209, 3, 20, 10, 0, 207, 209, 3,
		28, 14, 0, 208, 197, 1, 0, 0, 0, 208, 198, 1, 0, 0, 0, 208, 199, 1, 0,
		0, 0, 208, 200, 1, 0, 0, 0, 208, 201, 1, 0, 0, 0, 208, 202, 1, 0, 0, 0,
		208, 203, 1, 0, 0, 0, 208, 204, 1, 0, 0, 0, 208, 205, 1, 0, 0, 0, 208,
		206, 1, 0, 0, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208,
		1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213, 1, 0, 0, 0, 212, 210, 1, 0,
		0, 0, 213, 214, 5, 2, 0, 0, 214, 13, 1, 0, 0, 0, 215, 217, 5, 43, 0, 0,
		216, 215, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218,
		220, 5, 33, 0, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221,
		1, 0, 0, 0, 221, 222, 5, 20, 0, 0, 222, 224, 5, 42, 0, 0, 223, 225, 3,
		32, 16, 0, 224, 223, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 235, 1, 0,
		0, 0, 226, 230, 5, 1, 0, 0, 227, 229, 3, 34, 17, 0, 228, 227, 1, 0, 0,
		0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231,
		233, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 236, 5, 2, 0, 0, 234, 236,
		3, 34, 17, 0, 235, 226, 1, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 15, 1, 0,
		0, 0, 237, 239, 5, 43, 0, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0,
		239, 241, 1, 0, 0, 0, 240, 242, 5, 33, 0, 0, 241, 240, 1, 0, 0, 0, 241,
		242, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 5, 21, 0, 0, 244, 246,
		5, 42, 0, 0, 245, 247, 3, 32, 16, 0, 246, 245, 1, 0, 0, 0, 246, 247, 1,
		0, 0, 0, 247, 257, 1, 0, 0, 0, 248, 252, 5, 1, 0, 0, 249, 251, 3, 34, 17,
		0, 250, 249, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252,
		253, 1, 0, 0, 0, 253, 255, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 258,
		5, 2, 0, 0, 256, 258, 3, 34, 17, 0, 257, 248, 1, 0, 0, 0, 257, 256, 1,
		0, 0, 0, 258, 17, 1, 0, 0, 0, 259, 261, 5, 43, 0, 0, 260, 259, 1, 0, 0,
		0, 260, 261, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 5, 22, 0, 0, 263,
		265, 5, 42, 0, 0, 264, 266, 3, 32, 16, 0, 265, 264, 1, 0, 0, 0, 265, 266,
		1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 277, 5, 1, 0, 0, 268, 270, 5, 43,
		0, 0, 269, 268, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 272, 1, 0, 0, 0,
		271, 273, 3, 40, 20, 0, 272, 271, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273,
		274, 1, 0, 0, 0, 274, 276, 3, 56, 28, 0, 275, 269, 1, 0, 0, 0, 276, 279,
		1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 280, 1, 0,
		0, 0, 279, 277, 1, 0, 0, 0, 280, 281, 5, 2, 0, 0, 281, 19, 1, 0, 0, 0,
		282, 284, 5, 43, 0, 0, 283, 282, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284,
		285, 1, 0, 0, 0, 285, 286, 5, 23, 0, 0, 286, 288, 5, 42, 0, 0, 287, 289,
		3, 32, 16, 0, 288, 287, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 1,
		0, 0, 0, 290, 292, 3, 40, 20, 0, 291, 290, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 293, 1, 0, 0, 0, 293, 294, 3, 56, 28, 0, 294, 21, 1, 0, 0, 0,
		295, 297, 5, 43, 0, 0, 296, 295, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297,
		298, 1, 0, 0, 0, 298, 299, 7, 0, 0, 0, 299, 300, 5, 26, 0, 0, 300, 302,
		5, 42, 0, 0, 301, 303, 3, 32, 16, 0, 302, 301, 1, 0, 0, 0, 302, 303, 1,
		0, 0, 0, 303, 305, 1, 0, 0, 0, 304, 306, 3, 40, 20, 0, 305, 304, 1, 0,
		0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 3, 56, 28,
		0, 308, 23, 1, 0, 0, 0, 309, 311, 5, 43, 0, 0, 310, 309, 1, 0, 0, 0, 310,
		311, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 313, 5, 24, 0, 0, 313, 314,
		5, 26, 0, 0, 314, 316, 5, 42, 0, 0, 315, 317, 3, 32, 16, 0, 316, 315, 1,
		0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 25, 1, 0, 0, 0, 318, 320, 5, 43, 0,
		0, 319, 318, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		322, 5, 25, 0, 0, 322, 323, 5, 26, 0, 0, 323, 325, 5, 42, 0, 0, 324, 326,
		3, 32, 16, 0, 325, 324, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 27, 1, 0,
		0, 0, 327, 329, 5, 43, 0, 0, 328, 327, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0,
		329, 330, 1, 0, 0, 0, 330, 332, 7, 1, 0, 0, 331, 333, 3, 32, 16, 0, 332,
		331, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 338, 1, 0, 0, 0, 334, 336,
		3, 40, 20, 0, 335, 334, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 337, 1,
		0, 0, 0, 337, 339, 3, 56, 28, 0, 338, 335, 1, 0, 0, 0, 338, 339, 1, 0,
		0, 0, 339, 29, 1, 0, 0, 0, 340, 342, 5, 33, 0, 0, 341, 343, 3, 40, 20,
		0, 342, 341, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344,
		345, 5, 31, 0, 0, 345, 346, 5, 42, 0, 0, 346, 31, 1, 0, 0, 0, 347, 348,
		5, 4, 0, 0, 348, 349, 5, 42, 0, 0, 349, 350, 5, 5, 0, 0, 350, 33, 1, 0,
		0, 0, 351, 353, 5, 43, 0, 0, 352, 351, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0,
		353, 358, 1, 0, 0, 0, 354, 355, 5, 6, 0, 0, 355, 356, 5, 30, 0, 0, 356,
		359, 5, 7, 0, 0, 357, 359, 3, 44, 22, 0, 358, 354, 1, 0, 0, 0, 358, 357,
		1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 361, 1, 0, 0, 0, 360, 362, 3, 40,
		20, 0, 361, 360, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0,
		363, 364, 3, 56, 28, 0, 364, 35, 1, 0, 0, 0, 365, 367, 5, 43, 0, 0, 366,
		365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 397, 1, 0, 0, 0, 368, 369,
		7, 2, 0, 0, 369, 398, 3, 40, 20, 0, 370, 371, 5, 28, 0, 0, 371, 372, 5,
		42, 0, 0, 372, 373, 5, 8, 0, 0, 373, 398, 5, 34, 0, 0, 374, 375, 5, 35,
		0, 0, 375, 398, 3, 44, 22, 0, 376, 378, 3, 38, 19, 0, 377, 379, 3, 44,
		22, 0, 378, 377, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0,
		380, 382, 3, 40, 20, 0, 381, 383, 3, 56, 28, 0, 382, 381, 1, 0, 0, 0, 382,
		383, 1, 0, 0, 0, 383, 398, 1, 0, 0, 0, 384, 386, 3, 38, 19, 0, 385, 387,
		3, 44, 22, 0, 386, 385, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 388, 1,
		0, 0, 0, 388, 389, 3, 56, 28, 0, 389, 398, 1, 0, 0, 0, 390, 392, 3, 44,
		22, 0, 391, 390, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 394, 1, 0, 0, 0,
		393, 395, 3, 40, 20, 0, 394, 393, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395,
		396, 1, 0, 0, 0, 396, 398, 3, 56, 28, 0, 397, 368, 1, 0, 0, 0, 397, 370,
		1, 0, 0, 0, 397, 374, 1, 0, 0, 0, 397, 376, 1, 0, 0, 0, 397, 384, 1, 0,
		0, 0, 397, 391, 1, 0, 0, 0, 398, 37, 1, 0, 0, 0, 399, 400, 5, 28, 0, 0,
		400, 406, 5, 42, 0, 0, 401, 402, 5, 29, 0, 0, 402, 403, 5, 9, 0, 0, 403,
		404, 7, 3, 0, 0, 404, 406, 5, 10, 0, 0, 405, 399, 1, 0, 0, 0, 405, 401,
		1, 0, 0, 0, 406, 39, 1, 0, 0, 0, 407, 408, 5, 8, 0, 0, 408, 409, 3, 42,
		21, 0, 409, 41, 1, 0, 0, 0, 410, 415, 5, 42, 0, 0, 411, 412, 5, 11, 0,
		0, 412, 414, 5, 42, 0, 0, 413, 411, 1, 0, 0, 0, 414, 417, 1, 0, 0, 0, 415,
		413, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 43, 1, 0, 0, 0, 417, 415, 1,
		0, 0, 0, 418, 419, 5, 6, 0, 0, 419, 420, 3, 46, 23, 0, 420, 421, 5, 7,
		0, 0, 421, 45, 1, 0, 0, 0, 422, 423, 3, 48, 24, 0, 423, 47, 1, 0, 0, 0,
		424, 429, 3, 50, 25, 0, 425, 426, 5, 37, 0, 0, 426, 428, 3, 50, 25, 0,
		427, 425, 1, 0, 0, 0, 428, 431, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429,
		430, 1, 0, 0, 0, 430, 49, 1, 0, 0, 0, 431, 429, 1, 0, 0, 0, 432, 437, 3,
		52, 26, 0, 433, 434, 5, 36, 0, 0, 434, 436, 3, 52, 26, 0, 435, 433, 1,
		0, 0, 0, 436, 439, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0,
		0, 438, 51, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 440, 442, 5, 38, 0, 0, 441,
		440, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 444,
		3, 54, 27, 0, 444, 53, 1, 0, 0, 0, 445, 451, 5, 42, 0, 0, 446, 447, 5,
		9, 0, 0, 447, 448, 3, 46, 23, 0, 448, 449, 5, 10, 0, 0, 449, 451, 1, 0,
		0, 0, 450, 445, 1, 0, 0, 0, 450, 446, 1, 0, 0, 0, 451, 55, 1, 0, 0, 0,
		452, 453, 5, 31, 0, 0, 453, 469, 7, 4, 0, 0, 454, 456, 5, 31, 0, 0, 455,
		457, 5, 32, 0, 0, 456, 455, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 458,
		1, 0, 0, 0, 458, 469, 5, 42, 0, 0, 459, 461, 5, 31, 0, 0, 460, 462, 5,
		32, 0, 0, 461, 460, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0, 462, 465, 1, 0, 0,
		0, 463, 464, 5, 42, 0, 0, 464, 466, 5, 12, 0, 0, 465, 463, 1, 0, 0, 0,
		465, 466, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 469, 7, 1, 0, 0, 468,
		452, 1, 0, 0, 0, 468, 454, 1, 0, 0, 0, 468, 459, 1, 0, 0, 0, 469, 57, 1,
		0, 0, 0, 79, 59, 77, 79, 86, 89, 94, 110, 112, 118, 121, 127, 133, 135,
		141, 144, 150, 153, 159, 161, 167, 172, 175, 178, 183, 186, 189, 194, 208,
		210, 216, 219, 224, 230, 235, 238, 241, 246, 252, 257, 260, 265, 269, 272,
		277, 283, 288, 291, 296, 302, 305, 310, 316, 319, 325, 328, 332, 335, 338,
		342, 352, 358, 361, 366, 378, 382, 386, 391, 394, 397, 405, 415, 429, 437,
		441, 450, 456, 461, 465, 468,
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
	fsmParserFSM        = 13
	fsmParserSTATE      = 14
	fsmParserPARALLEL   = 15
	fsmParserSUBMACHINE = 16
	fsmParserREGION     = 17
	fsmParserFINAL      = 18
	fsmParserTERMINATE  = 19
	fsmParserCHOICE     = 20
	fsmParserJUNCTION   = 21
	fsmParserFORK       = 22
	fsmParserJOIN       = 23
	fsmParserENTRY      = 24
	fsmParserEXIT       = 25
	fsmParserPOINT      = 26
	fsmParserDO         = 27
	fsmParserON         = 28
	fsmParserAFTER      = 29
	fsmParserELSE       = 30
	fsmParserGOTO       = 31
	fsmParserLOCAL      = 32
	fsmParserINITIAL    = 33
	fsmParserDEFER      = 34
	fsmParserINVARIANT  = 35
	fsmParserAND        = 36
	fsmParserOR         = 37
	fsmParserNOT        = 38
	fsmParserH          = 39
	fsmParserH_DEEP     = 40
	fsmParserDuration   = 41
	fsmParserIdentifier = 42
	fsmParserNote       = 43
	fsmParserComment    = 44
	fsmParserBlank      = 45
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
	fsmParserRULE_outlet            = 13
	fsmParserRULE_history           = 14
	fsmParserRULE_start             = 15
	fsmParserRULE_stereotype        = 16
	fsmParserRULE_branch            = 17
	fsmParserRULE_event             = 18
	fsmParserRULE_trigger           = 19
	fsmParserRULE_actions           = 20
	fsmParserRULE_identifiers       = 21
	fsmParserRULE_guard             = 22
	fsmParserRULE_expression        = 23
	fsmParserRULE_or_expression     = 24
	fsmParserRULE_and_expression    = 25
	fsmParserRULE_not_expression    = 26
	fsmParserRULE_single_expression = 27
	fsmParserRULE_goto              = 28
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
	AllOutlet() []IOutletContext
	Outlet(i int) IOutletContext
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

func (s *FsmContext) AllOutlet() []IOutletContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOutletContext); ok {
			len++
		}
	}

	tst := make([]IOutletContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOutletContext); ok {
			tst[i] = t.(IOutletContext)
			i++
		}
	}

	return tst
}

func (s *FsmContext) Outlet(i int) IOutletContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutletContext); ok {
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

	return t.(IOutletContext)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(58)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(61)
		p.Match(fsmParserFSM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8842196664640) != 0 {
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(64)
				p.Start_()
			}

		case 2:
			{
				p.SetState(65)
				p.State()
			}

		case 3:
			{
				p.SetState(66)
				p.Parallel()
			}

		case 4:
			{
				p.SetState(67)
				p.Submachine()
			}

		case 5:
			{
				p.SetState(68)
				p.Final()
			}

		case 6:
			{
				p.SetState(69)
				p.Terminate()
			}

		case 7:
			{
				p.SetState(70)
				p.Choice()
			}

		case 8:
			{
				p.SetState(71)
				p.Junction()
			}

		case 9:
			{
				p.SetState(72)
				p.Fork()
			}

		case 10:
			{
				p.SetState(73)
				p.Join()
			}

		case 11:
			{
				p.SetState(74)
				p.Point()
			}

		case 12:
			{
				p.SetState(75)
				p.Outlet()
			}

		case 13:
			{
				p.SetState(76)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(85)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(88)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(91)
		p.Match(fsmParserSTATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(93)
			p.Stereotype()
		}

	}
	{
		p.SetState(96)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10491464106304) != 0 {
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(97)
				p.Start_()
			}

		case 2:
			{
				p.SetState(98)
				p.State()
			}

		case 3:
			{
				p.SetState(99)
				p.Parallel()
			}

		case 4:
			{
				p.SetState(100)
				p.Submachine()
			}

		case 5:
			{
				p.SetState(101)
				p.Final()
			}

		case 6:
			{
				p.SetState(102)
				p.Terminate()
			}

		case 7:
			{
				p.SetState(103)
				p.Choice()
			}

		case 8:
			{
				p.SetState(104)
				p.Junction()
			}

		case 9:
			{
				p.SetState(105)
				p.Fork()
			}

		case 10:
			{
				p.SetState(106)
				p.Join()
			}

		case 11:
			{
				p.SetState(107)
				p.Point()
			}

		case 12:
			{
				p.SetState(108)
				p.History()
			}

		case 13:
			{
				p.SetState(109)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(117)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(120)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(123)
		p.Match(fsmParserPARALLEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(fsmParserSTATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(126)
			p.Stereotype()
		}

	}
	{
		p.SetState(129)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8833590231360) != 0 {
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(130)
				p.Region()
			}

		case 2:
			{
				p.SetState(131)
				p.Point()
			}

		case 3:
			{
				p.SetState(132)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(138)
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
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
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

func (s *SubmachineContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(fsmParserIdentifier)
}

func (s *SubmachineContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, i)
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(140)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(143)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(146)
		p.Match(fsmParserSUBMACHINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__2 {
		{
			p.SetState(148)
			p.Match(fsmParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(152)
			p.Stereotype()
		}

	}
	{
		p.SetState(155)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8833590100288) != 0 {
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(156)
				p.Point()
			}

		case 2:
			{
				p.SetState(157)
				p.Reference()
			}

		case 3:
			{
				p.SetState(158)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
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
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(166)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(169)
		p.Match(fsmParserFINAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(fsmParserSTATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserIdentifier {
		{
			p.SetState(171)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(174)
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
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(177)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(180)
		p.Match(fsmParserTERMINATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(fsmParserSTATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserIdentifier {
		{
			p.SetState(182)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(185)
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
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(188)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(191)
		p.Match(fsmParserREGION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(193)
			p.Stereotype()
		}

	}
	{
		p.SetState(196)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10453967028224) != 0 {
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(197)
				p.Start_()
			}

		case 2:
			{
				p.SetState(198)
				p.State()
			}

		case 3:
			{
				p.SetState(199)
				p.Parallel()
			}

		case 4:
			{
				p.SetState(200)
				p.Submachine()
			}

		case 5:
			{
				p.SetState(201)
				p.Final()
			}

		case 6:
			{
				p.SetState(202)
				p.Terminate()
			}

		case 7:
			{
				p.SetState(203)
				p.Choice()
			}

		case 8:
			{
				p.SetState(204)
				p.Junction()
			}

		case 9:
			{
				p.SetState(205)
				p.Fork()
			}

		case 10:
			{
				p.SetState(206)
				p.Join()
			}

		case 11:
			{
				p.SetState(207)
				p.History()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(213)
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
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(215)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(218)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(221)
		p.Match(fsmParserCHOICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(223)
			p.Stereotype()
		}

	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__0:
		{
			p.SetState(226)
			p.Match(fsmParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8798240506176) != 0 {
			{
				p.SetState(227)
				p.Branch()
			}

			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(233)
			p.Match(fsmParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__5, fsmParserT__7, fsmParserGOTO, fsmParserNote:
		{
			p.SetState(234)
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
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(237)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserINITIAL {
		{
			p.SetState(240)
			p.Match(fsmParserINITIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(243)
		p.Match(fsmParserJUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(245)
			p.Stereotype()
		}

	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__0:
		{
			p.SetState(248)
			p.Match(fsmParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8798240506176) != 0 {
			{
				p.SetState(249)
				p.Branch()
			}

			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(255)
			p.Match(fsmParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__5, fsmParserT__7, fsmParserGOTO, fsmParserNote:
		{
			p.SetState(256)
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
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(259)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(262)
		p.Match(fsmParserFORK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(264)
			p.Stereotype()
		}

	}
	{
		p.SetState(267)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8798240506112) != 0 {
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserNote {
			{
				p.SetState(268)
				p.Match(fsmParserNote)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__7 {
			{
				p.SetState(271)
				p.Actions()
			}

		}
		{
			p.SetState(274)
			p.Goto_()
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(280)
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
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(282)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(285)
		p.Match(fsmParserJOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(287)
			p.Stereotype()
		}

	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__7 {
		{
			p.SetState(290)
			p.Actions()
		}

	}
	{
		p.SetState(293)
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
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(295)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(298)

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
		p.SetState(299)
		p.Match(fsmParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(301)
			p.Stereotype()
		}

	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__7 {
		{
			p.SetState(304)
			p.Actions()
		}

	}
	{
		p.SetState(307)
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
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(309)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(312)
		p.Match(fsmParserENTRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(fsmParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(315)
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

// IOutletContext is an interface to support dynamic dispatch.
type IOutletContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXIT() antlr.TerminalNode
	POINT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	Stereotype() IStereotypeContext

	// IsOutletContext differentiates from other interfaces.
	IsOutletContext()
}

type OutletContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutletContext() *OutletContext {
	var p = new(OutletContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_outlet
	return p
}

func InitEmptyOutletContext(p *OutletContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fsmParserRULE_outlet
}

func (*OutletContext) IsOutletContext() {}

func NewOutletContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutletContext {
	var p = new(OutletContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fsmParserRULE_outlet

	return p
}

func (s *OutletContext) GetParser() antlr.Parser { return s.parser }

func (s *OutletContext) EXIT() antlr.TerminalNode {
	return s.GetToken(fsmParserEXIT, 0)
}

func (s *OutletContext) POINT() antlr.TerminalNode {
	return s.GetToken(fsmParserPOINT, 0)
}

func (s *OutletContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *OutletContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *OutletContext) Stereotype() IStereotypeContext {
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

func (s *OutletContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutletContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutletContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fsmVisitor:
		return t.VisitOutlet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fsmParser) Outlet() (localctx IOutletContext) {
	localctx = NewOutletContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, fsmParserRULE_outlet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(318)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(321)
		p.Match(fsmParserEXIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(fsmParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(324)
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
	p.EnterRule(localctx, 28, fsmParserRULE_history)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(327)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(330)

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
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__3 {
		{
			p.SetState(331)
			p.Stereotype()
		}

	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__7 {
			{
				p.SetState(334)
				p.Actions()
			}

		}
		{
			p.SetState(337)
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
	p.EnterRule(localctx, 30, fsmParserRULE_start)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(fsmParserINITIAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__7 {
		{
			p.SetState(341)
			p.Actions()
		}

	}
	{
		p.SetState(344)
		p.Match(fsmParserGOTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
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
	p.EnterRule(localctx, 32, fsmParserRULE_stereotype)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Match(fsmParserT__4)
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
	p.EnterRule(localctx, 34, fsmParserRULE_branch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(351)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(354)
			p.Match(fsmParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Match(fsmParserELSE)
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

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(357)
			p.Guard()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__7 {
		{
			p.SetState(360)
			p.Actions()
		}

	}
	{
		p.SetState(363)
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
	p.EnterRule(localctx, 36, fsmParserRULE_event)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(365)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(368)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*EventContext).name = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&184549376) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*EventContext).name = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(369)
			p.Actions()
		}

	case 2:
		{
			p.SetState(370)
			p.Match(fsmParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*EventContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Match(fsmParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Match(fsmParserDEFER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(374)
			p.Match(fsmParserINVARIANT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Guard()
		}

	case 4:
		{
			p.SetState(376)
			p.Trigger()
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__5 {
			{
				p.SetState(377)
				p.Guard()
			}

		}
		{
			p.SetState(380)
			p.Actions()
		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(381)
				p.Goto_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		{
			p.SetState(384)
			p.Trigger()
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__5 {
			{
				p.SetState(385)
				p.Guard()
			}

		}
		{
			p.SetState(388)
			p.Goto_()
		}

	case 6:
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__5 {
			{
				p.SetState(390)
				p.Guard()
			}

		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__7 {
			{
				p.SetState(393)
				p.Actions()
			}

		}
		{
			p.SetState(396)
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
	p.EnterRule(localctx, 38, fsmParserRULE_trigger)
	var _la int

	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(399)
			p.Match(fsmParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)

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
			p.SetState(401)
			p.Match(fsmParserAFTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Match(fsmParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)

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
			p.SetState(404)
			p.Match(fsmParserT__9)
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
	p.EnterRule(localctx, 40, fsmParserRULE_actions)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.Match(fsmParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
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
	p.EnterRule(localctx, 42, fsmParserRULE_identifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__10 {
		{
			p.SetState(411)
			p.Match(fsmParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(417)
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
	p.EnterRule(localctx, 44, fsmParserRULE_guard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(fsmParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.Expression()
	}
	{
		p.SetState(420)
		p.Match(fsmParserT__6)
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
	p.EnterRule(localctx, 46, fsmParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
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
	p.EnterRule(localctx, 48, fsmParserRULE_or_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(424)
		p.And_expression()
	}
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserOR {
		{
			p.SetState(425)
			p.Match(fsmParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.And_expression()
		}

		p.SetState(431)
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
	p.EnterRule(localctx, 50, fsmParserRULE_and_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Not_expression()
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserAND {
		{
			p.SetState(433)
			p.Match(fsmParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Not_expression()
		}

		p.SetState(439)
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
	p.EnterRule(localctx, 52, fsmParserRULE_not_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNOT {
		{
			p.SetState(440)
			p.Match(fsmParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(443)
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
	p.EnterRule(localctx, 54, fsmParserRULE_single_expression)
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.Match(fsmParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Expression()
		}
		{
			p.SetState(448)
			p.Match(fsmParserT__9)
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
	p.EnterRule(localctx, 56, fsmParserRULE_goto)
	var _la int

	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.Match(fsmParserGOTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
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
			p.SetState(454)
			p.Match(fsmParserGOTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserLOCAL {
			{
				p.SetState(455)
				p.Match(fsmParserLOCAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(458)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(459)
			p.Match(fsmParserGOTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(461)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserLOCAL {
			{
				p.SetState(460)
				p.Match(fsmParserLOCAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserIdentifier {
			{
				p.SetState(463)
				p.Match(fsmParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(464)
				p.Match(fsmParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(467)
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
