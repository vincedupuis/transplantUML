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
		"", "'fsm'", "'{'", "'}'", "'state'", "'parallel'", "'submachine'",
		"'final'", "'terminate'", "'region'", "'choice'", "'junction'", "'fork'",
		"'join'", "'entry'", "'exit'", "'point'", "'H'", "'H*'", "'<<'", "'>>'",
		"'['", "'else'", "']'", "'do'", "'on'", "'/'", "'after'", "'('", "')'",
		"','", "'or'", "'and'", "'not'", "'goto'", "'.'", "'initial'", "'invariant'",
		"'defer'", "'local'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "Initial", "Invariant", "Defer", "Local", "Duration", "Identifier",
		"Note", "Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"fsm", "state", "parallel", "submachine", "final", "terminate", "region",
		"choice", "junction", "fork", "join", "point", "history", "stereotype",
		"branch", "event", "trigger", "actions", "identifiers", "guard", "expression",
		"or_expression", "and_expression", "not_expression", "single_expression",
		"goto",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 430, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		3, 0, 54, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 70, 8, 0, 10, 0, 12, 0, 73, 9, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 3, 1, 79, 8, 1, 1, 1, 3, 1, 82, 8, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 87, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 102, 8, 1, 10, 1, 12, 1, 105, 9, 1, 1, 1,
		1, 1, 1, 2, 3, 2, 110, 8, 2, 1, 2, 3, 2, 113, 8, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 119, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 125, 8, 2, 10, 2, 12,
		2, 128, 9, 2, 1, 2, 1, 2, 1, 3, 3, 3, 133, 8, 3, 1, 3, 3, 3, 136, 8, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 141, 8, 3, 1, 3, 1, 3, 5, 3, 145, 8, 3, 10, 3,
		12, 3, 148, 9, 3, 1, 3, 1, 3, 1, 4, 3, 4, 153, 8, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 158, 8, 4, 1, 4, 3, 4, 161, 8, 4, 1, 5, 3, 5, 164, 8, 5, 1, 5, 1,
		5, 1, 5, 3, 5, 169, 8, 5, 1, 5, 3, 5, 172, 8, 5, 1, 6, 3, 6, 175, 8, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 180, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 193, 8, 6, 10, 6, 12, 6, 196, 9, 6,
		1, 6, 1, 6, 1, 7, 3, 7, 201, 8, 7, 1, 7, 3, 7, 204, 8, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 209, 8, 7, 1, 7, 1, 7, 5, 7, 213, 8, 7, 10, 7, 12, 7, 216, 9,
		7, 1, 7, 1, 7, 3, 7, 220, 8, 7, 1, 8, 3, 8, 223, 8, 8, 1, 8, 3, 8, 226,
		8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 231, 8, 8, 1, 8, 1, 8, 5, 8, 235, 8, 8, 10,
		8, 12, 8, 238, 9, 8, 1, 8, 1, 8, 3, 8, 242, 8, 8, 1, 9, 3, 9, 245, 8, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 250, 8, 9, 1, 9, 1, 9, 3, 9, 254, 8, 9, 1, 9, 3,
		9, 257, 8, 9, 1, 9, 5, 9, 260, 8, 9, 10, 9, 12, 9, 263, 9, 9, 1, 9, 1,
		9, 1, 10, 3, 10, 268, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 273, 8, 10, 1,
		10, 3, 10, 276, 8, 10, 1, 10, 1, 10, 1, 11, 3, 11, 281, 8, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 287, 8, 11, 1, 11, 3, 11, 290, 8, 11, 1, 11, 1,
		11, 1, 12, 3, 12, 295, 8, 12, 1, 12, 1, 12, 3, 12, 299, 8, 12, 1, 12, 3,
		12, 302, 8, 12, 1, 12, 3, 12, 305, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		14, 3, 14, 312, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 318, 8, 14, 1,
		14, 3, 14, 321, 8, 14, 1, 14, 1, 14, 1, 15, 3, 15, 326, 8, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 338,
		8, 15, 1, 15, 1, 15, 3, 15, 342, 8, 15, 1, 15, 1, 15, 3, 15, 346, 8, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 351, 8, 15, 1, 15, 3, 15, 354, 8, 15, 1, 15,
		3, 15, 357, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 365,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 5, 18, 373, 8, 18, 10,
		18, 12, 18, 376, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 5, 21, 387, 8, 21, 10, 21, 12, 21, 390, 9, 21, 1, 22, 1,
		22, 1, 22, 5, 22, 395, 8, 22, 10, 22, 12, 22, 398, 9, 22, 1, 23, 3, 23,
		401, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 410,
		8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 416, 8, 25, 1, 25, 1, 25, 1,
		25, 3, 25, 421, 8, 25, 1, 25, 1, 25, 3, 25, 425, 8, 25, 1, 25, 3, 25, 428,
		8, 25, 1, 25, 0, 0, 26, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 0, 5, 1, 0, 14, 15,
		1, 0, 17, 18, 2, 0, 14, 15, 24, 24, 1, 0, 40, 41, 1, 0, 7, 8, 509, 0, 53,
		1, 0, 0, 0, 2, 78, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 132, 1, 0, 0, 0,
		8, 152, 1, 0, 0, 0, 10, 163, 1, 0, 0, 0, 12, 174, 1, 0, 0, 0, 14, 200,
		1, 0, 0, 0, 16, 222, 1, 0, 0, 0, 18, 244, 1, 0, 0, 0, 20, 267, 1, 0, 0,
		0, 22, 280, 1, 0, 0, 0, 24, 294, 1, 0, 0, 0, 26, 306, 1, 0, 0, 0, 28, 311,
		1, 0, 0, 0, 30, 325, 1, 0, 0, 0, 32, 364, 1, 0, 0, 0, 34, 366, 1, 0, 0,
		0, 36, 369, 1, 0, 0, 0, 38, 377, 1, 0, 0, 0, 40, 381, 1, 0, 0, 0, 42, 383,
		1, 0, 0, 0, 44, 391, 1, 0, 0, 0, 46, 400, 1, 0, 0, 0, 48, 409, 1, 0, 0,
		0, 50, 427, 1, 0, 0, 0, 52, 54, 5, 42, 0, 0, 53, 52, 1, 0, 0, 0, 53, 54,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 5, 1, 0, 0, 56, 57, 5, 41, 0, 0,
		57, 71, 5, 2, 0, 0, 58, 70, 3, 2, 1, 0, 59, 70, 3, 4, 2, 0, 60, 70, 3,
		6, 3, 0, 61, 70, 3, 8, 4, 0, 62, 70, 3, 10, 5, 0, 63, 70, 3, 14, 7, 0,
		64, 70, 3, 16, 8, 0, 65, 70, 3, 18, 9, 0, 66, 70, 3, 20, 10, 0, 67, 70,
		3, 22, 11, 0, 68, 70, 3, 30, 15, 0, 69, 58, 1, 0, 0, 0, 69, 59, 1, 0, 0,
		0, 69, 60, 1, 0, 0, 0, 69, 61, 1, 0, 0, 0, 69, 62, 1, 0, 0, 0, 69, 63,
		1, 0, 0, 0, 69, 64, 1, 0, 0, 0, 69, 65, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0,
		69, 67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74,
		75, 5, 3, 0, 0, 75, 76, 5, 0, 0, 1, 76, 1, 1, 0, 0, 0, 77, 79, 5, 42, 0,
		0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81, 1, 0, 0, 0, 80, 82,
		5, 36, 0, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0,
		83, 84, 5, 4, 0, 0, 84, 86, 5, 41, 0, 0, 85, 87, 3, 26, 13, 0, 86, 85,
		1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 103, 5, 2, 0, 0,
		89, 102, 3, 2, 1, 0, 90, 102, 3, 4, 2, 0, 91, 102, 3, 6, 3, 0, 92, 102,
		3, 8, 4, 0, 93, 102, 3, 10, 5, 0, 94, 102, 3, 14, 7, 0, 95, 102, 3, 16,
		8, 0, 96, 102, 3, 18, 9, 0, 97, 102, 3, 20, 10, 0, 98, 102, 3, 22, 11,
		0, 99, 102, 3, 24, 12, 0, 100, 102, 3, 30, 15, 0, 101, 89, 1, 0, 0, 0,
		101, 90, 1, 0, 0, 0, 101, 91, 1, 0, 0, 0, 101, 92, 1, 0, 0, 0, 101, 93,
		1, 0, 0, 0, 101, 94, 1, 0, 0, 0, 101, 95, 1, 0, 0, 0, 101, 96, 1, 0, 0,
		0, 101, 97, 1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101,
		100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104,
		1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107, 5, 3,
		0, 0, 107, 3, 1, 0, 0, 0, 108, 110, 5, 42, 0, 0, 109, 108, 1, 0, 0, 0,
		109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 113, 5, 36, 0, 0, 112,
		111, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115,
		5, 5, 0, 0, 115, 116, 5, 4, 0, 0, 116, 118, 5, 41, 0, 0, 117, 119, 3, 26,
		13, 0, 118, 117, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0,
		120, 126, 5, 2, 0, 0, 121, 125, 3, 12, 6, 0, 122, 125, 3, 22, 11, 0, 123,
		125, 3, 30, 15, 0, 124, 121, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 123,
		1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0,
		0, 0, 127, 129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 130, 5, 3, 0, 0,
		130, 5, 1, 0, 0, 0, 131, 133, 5, 42, 0, 0, 132, 131, 1, 0, 0, 0, 132, 133,
		1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 136, 5, 36, 0, 0, 135, 134, 1, 0,
		0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 5, 6, 0, 0,
		138, 140, 5, 41, 0, 0, 139, 141, 3, 26, 13, 0, 140, 139, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 146, 5, 2, 0, 0, 143, 145,
		3, 30, 15, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1,
		0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 146, 1, 0, 0,
		0, 149, 150, 5, 3, 0, 0, 150, 7, 1, 0, 0, 0, 151, 153, 5, 42, 0, 0, 152,
		151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155,
		5, 7, 0, 0, 155, 157, 5, 4, 0, 0, 156, 158, 5, 41, 0, 0, 157, 156, 1, 0,
		0, 0, 157, 158, 1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159, 161, 3, 26, 13,
		0, 160, 159, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 9, 1, 0, 0, 0, 162,
		164, 5, 42, 0, 0, 163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165,
		1, 0, 0, 0, 165, 166, 5, 8, 0, 0, 166, 168, 5, 4, 0, 0, 167, 169, 5, 41,
		0, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 171, 1, 0, 0, 0,
		170, 172, 3, 26, 13, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172,
		11, 1, 0, 0, 0, 173, 175, 5, 42, 0, 0, 174, 173, 1, 0, 0, 0, 174, 175,
		1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 5, 9, 0, 0, 177, 179, 5, 41,
		0, 0, 178, 180, 3, 26, 13, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0,
		0, 180, 181, 1, 0, 0, 0, 181, 194, 5, 2, 0, 0, 182, 193, 3, 2, 1, 0, 183,
		193, 3, 4, 2, 0, 184, 193, 3, 6, 3, 0, 185, 193, 3, 8, 4, 0, 186, 193,
		3, 10, 5, 0, 187, 193, 3, 14, 7, 0, 188, 193, 3, 16, 8, 0, 189, 193, 3,
		18, 9, 0, 190, 193, 3, 20, 10, 0, 191, 193, 3, 24, 12, 0, 192, 182, 1,
		0, 0, 0, 192, 183, 1, 0, 0, 0, 192, 184, 1, 0, 0, 0, 192, 185, 1, 0, 0,
		0, 192, 186, 1, 0, 0, 0, 192, 187, 1, 0, 0, 0, 192, 188, 1, 0, 0, 0, 192,
		189, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 196,
		1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0,
		0, 0, 196, 194, 1, 0, 0, 0, 197, 198, 5, 3, 0, 0, 198, 13, 1, 0, 0, 0,
		199, 201, 5, 42, 0, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		203, 1, 0, 0, 0, 202, 204, 5, 36, 0, 0, 203, 202, 1, 0, 0, 0, 203, 204,
		1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 5, 10, 0, 0, 206, 208, 5, 41,
		0, 0, 207, 209, 3, 26, 13, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0, 0,
		0, 209, 219, 1, 0, 0, 0, 210, 214, 5, 2, 0, 0, 211, 213, 3, 28, 14, 0,
		212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214,
		215, 1, 0, 0, 0, 215, 217, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 220,
		5, 3, 0, 0, 218, 220, 3, 28, 14, 0, 219, 210, 1, 0, 0, 0, 219, 218, 1,
		0, 0, 0, 220, 15, 1, 0, 0, 0, 221, 223, 5, 42, 0, 0, 222, 221, 1, 0, 0,
		0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 226, 5, 36, 0, 0, 225,
		224, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228,
		5, 11, 0, 0, 228, 230, 5, 41, 0, 0, 229, 231, 3, 26, 13, 0, 230, 229, 1,
		0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 241, 1, 0, 0, 0, 232, 236, 5, 2, 0,
		0, 233, 235, 3, 28, 14, 0, 234, 233, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0,
		236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 239, 1, 0, 0, 0, 238,
		236, 1, 0, 0, 0, 239, 242, 5, 3, 0, 0, 240, 242, 3, 28, 14, 0, 241, 232,
		1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 17, 1, 0, 0, 0, 243, 245, 5, 42,
		0, 0, 244, 243, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0,
		246, 247, 5, 12, 0, 0, 247, 249, 5, 41, 0, 0, 248, 250, 3, 26, 13, 0, 249,
		248, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 261,
		5, 2, 0, 0, 252, 254, 5, 42, 0, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0,
		0, 0, 254, 256, 1, 0, 0, 0, 255, 257, 3, 34, 17, 0, 256, 255, 1, 0, 0,
		0, 256, 257, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 3, 50, 25, 0,
		259, 253, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261,
		262, 1, 0, 0, 0, 262, 264, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 264, 265,
		5, 3, 0, 0, 265, 19, 1, 0, 0, 0, 266, 268, 5, 42, 0, 0, 267, 266, 1, 0,
		0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 5, 13, 0, 0,
		270, 272, 5, 41, 0, 0, 271, 273, 3, 26, 13, 0, 272, 271, 1, 0, 0, 0, 272,
		273, 1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 276, 3, 34, 17, 0, 275, 274,
		1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 3, 50,
		25, 0, 278, 21, 1, 0, 0, 0, 279, 281, 5, 42, 0, 0, 280, 279, 1, 0, 0, 0,
		280, 281, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 7, 0, 0, 0, 283,
		284, 5, 16, 0, 0, 284, 286, 5, 41, 0, 0, 285, 287, 3, 26, 13, 0, 286, 285,
		1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 288, 290, 3, 34,
		17, 0, 289, 288, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0,
		291, 292, 3, 50, 25, 0, 292, 23, 1, 0, 0, 0, 293, 295, 5, 42, 0, 0, 294,
		293, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298,
		7, 1, 0, 0, 297, 299, 3, 26, 13, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1,
		0, 0, 0, 299, 304, 1, 0, 0, 0, 300, 302, 3, 34, 17, 0, 301, 300, 1, 0,
		0, 0, 301, 302, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 305, 3, 50, 25,
		0, 304, 301, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 25, 1, 0, 0, 0, 306,
		307, 5, 19, 0, 0, 307, 308, 5, 41, 0, 0, 308, 309, 5, 20, 0, 0, 309, 27,
		1, 0, 0, 0, 310, 312, 5, 42, 0, 0, 311, 310, 1, 0, 0, 0, 311, 312, 1, 0,
		0, 0, 312, 317, 1, 0, 0, 0, 313, 314, 5, 21, 0, 0, 314, 315, 5, 22, 0,
		0, 315, 318, 5, 23, 0, 0, 316, 318, 3, 38, 19, 0, 317, 313, 1, 0, 0, 0,
		317, 316, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 320, 1, 0, 0, 0, 319,
		321, 3, 34, 17, 0, 320, 319, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 322,
		1, 0, 0, 0, 322, 323, 3, 50, 25, 0, 323, 29, 1, 0, 0, 0, 324, 326, 5, 42,
		0, 0, 325, 324, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 356, 1, 0, 0, 0,
		327, 328, 7, 2, 0, 0, 328, 357, 3, 34, 17, 0, 329, 330, 5, 25, 0, 0, 330,
		331, 5, 41, 0, 0, 331, 332, 5, 26, 0, 0, 332, 357, 5, 38, 0, 0, 333, 334,
		5, 37, 0, 0, 334, 357, 3, 38, 19, 0, 335, 337, 3, 32, 16, 0, 336, 338,
		3, 38, 19, 0, 337, 336, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338, 339, 1,
		0, 0, 0, 339, 341, 3, 34, 17, 0, 340, 342, 3, 50, 25, 0, 341, 340, 1, 0,
		0, 0, 341, 342, 1, 0, 0, 0, 342, 357, 1, 0, 0, 0, 343, 345, 3, 32, 16,
		0, 344, 346, 3, 38, 19, 0, 345, 344, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0,
		346, 347, 1, 0, 0, 0, 347, 348, 3, 50, 25, 0, 348, 357, 1, 0, 0, 0, 349,
		351, 3, 38, 19, 0, 350, 349, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 353,
		1, 0, 0, 0, 352, 354, 3, 34, 17, 0, 353, 352, 1, 0, 0, 0, 353, 354, 1,
		0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 357, 3, 50, 25, 0, 356, 327, 1, 0,
		0, 0, 356, 329, 1, 0, 0, 0, 356, 333, 1, 0, 0, 0, 356, 335, 1, 0, 0, 0,
		356, 343, 1, 0, 0, 0, 356, 350, 1, 0, 0, 0, 357, 31, 1, 0, 0, 0, 358, 359,
		5, 25, 0, 0, 359, 365, 5, 41, 0, 0, 360, 361, 5, 27, 0, 0, 361, 362, 5,
		28, 0, 0, 362, 363, 7, 3, 0, 0, 363, 365, 5, 29, 0, 0, 364, 358, 1, 0,
		0, 0, 364, 360, 1, 0, 0, 0, 365, 33, 1, 0, 0, 0, 366, 367, 5, 26, 0, 0,
		367, 368, 3, 36, 18, 0, 368, 35, 1, 0, 0, 0, 369, 374, 5, 41, 0, 0, 370,
		371, 5, 30, 0, 0, 371, 373, 5, 41, 0, 0, 372, 370, 1, 0, 0, 0, 373, 376,
		1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 37, 1, 0,
		0, 0, 376, 374, 1, 0, 0, 0, 377, 378, 5, 21, 0, 0, 378, 379, 3, 40, 20,
		0, 379, 380, 5, 23, 0, 0, 380, 39, 1, 0, 0, 0, 381, 382, 3, 42, 21, 0,
		382, 41, 1, 0, 0, 0, 383, 388, 3, 44, 22, 0, 384, 385, 5, 31, 0, 0, 385,
		387, 3, 44, 22, 0, 386, 384, 1, 0, 0, 0, 387, 390, 1, 0, 0, 0, 388, 386,
		1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 43, 1, 0, 0, 0, 390, 388, 1, 0,
		0, 0, 391, 396, 3, 46, 23, 0, 392, 393, 5, 32, 0, 0, 393, 395, 3, 46, 23,
		0, 394, 392, 1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396,
		397, 1, 0, 0, 0, 397, 45, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399, 401, 5,
		33, 0, 0, 400, 399, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 402, 1, 0, 0,
		0, 402, 403, 3, 48, 24, 0, 403, 47, 1, 0, 0, 0, 404, 410, 5, 41, 0, 0,
		405, 406, 5, 28, 0, 0, 406, 407, 3, 40, 20, 0, 407, 408, 5, 29, 0, 0, 408,
		410, 1, 0, 0, 0, 409, 404, 1, 0, 0, 0, 409, 405, 1, 0, 0, 0, 410, 49, 1,
		0, 0, 0, 411, 412, 5, 34, 0, 0, 412, 428, 7, 4, 0, 0, 413, 415, 5, 34,
		0, 0, 414, 416, 5, 39, 0, 0, 415, 414, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0,
		416, 417, 1, 0, 0, 0, 417, 428, 5, 41, 0, 0, 418, 420, 5, 34, 0, 0, 419,
		421, 5, 39, 0, 0, 420, 419, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0, 421, 424,
		1, 0, 0, 0, 422, 423, 5, 41, 0, 0, 423, 425, 5, 35, 0, 0, 424, 422, 1,
		0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 428, 7, 1, 0,
		0, 427, 411, 1, 0, 0, 0, 427, 413, 1, 0, 0, 0, 427, 418, 1, 0, 0, 0, 428,
		51, 1, 0, 0, 0, 72, 53, 69, 71, 78, 81, 86, 101, 103, 109, 112, 118, 124,
		126, 132, 135, 140, 146, 152, 157, 160, 163, 168, 171, 174, 179, 192, 194,
		200, 203, 208, 214, 219, 222, 225, 230, 236, 241, 244, 249, 253, 256, 261,
		267, 272, 275, 280, 286, 289, 294, 298, 301, 304, 311, 317, 320, 325, 337,
		341, 345, 350, 353, 356, 364, 374, 388, 396, 400, 409, 415, 420, 424, 427,
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
	fsmParserT__20      = 21
	fsmParserT__21      = 22
	fsmParserT__22      = 23
	fsmParserT__23      = 24
	fsmParserT__24      = 25
	fsmParserT__25      = 26
	fsmParserT__26      = 27
	fsmParserT__27      = 28
	fsmParserT__28      = 29
	fsmParserT__29      = 30
	fsmParserT__30      = 31
	fsmParserT__31      = 32
	fsmParserT__32      = 33
	fsmParserT__33      = 34
	fsmParserT__34      = 35
	fsmParserInitial    = 36
	fsmParserInvariant  = 37
	fsmParserDefer      = 38
	fsmParserLocal      = 39
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
	fsmParserRULE_history           = 12
	fsmParserRULE_stereotype        = 13
	fsmParserRULE_branch            = 14
	fsmParserRULE_event             = 15
	fsmParserRULE_trigger           = 16
	fsmParserRULE_actions           = 17
	fsmParserRULE_identifiers       = 18
	fsmParserRULE_guard             = 19
	fsmParserRULE_expression        = 20
	fsmParserRULE_or_expression     = 21
	fsmParserRULE_and_expression    = 22
	fsmParserRULE_not_expression    = 23
	fsmParserRULE_single_expression = 24
	fsmParserRULE_goto              = 25
)

// IFsmContext is an interface to support dynamic dispatch.
type IFsmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	EOF() antlr.TerminalNode
	Note() antlr.TerminalNode
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

func (s *FsmContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *FsmContext) EOF() antlr.TerminalNode {
	return s.GetToken(fsmParserEOF, 0)
}

func (s *FsmContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(52)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(55)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4621638630896) != 0 {
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(58)
				p.State()
			}

		case 2:
			{
				p.SetState(59)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(60)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(61)
				p.Final()
			}

		case 5:
			{
				p.SetState(62)
				p.Terminate()
			}

		case 6:
			{
				p.SetState(63)
				p.Choice()
			}

		case 7:
			{
				p.SetState(64)
				p.Junction()
			}

		case 8:
			{
				p.SetState(65)
				p.Fork()
			}

		case 9:
			{
				p.SetState(66)
				p.Join()
			}

		case 10:
			{
				p.SetState(67)
				p.Point()
			}

		case 11:
			{
				p.SetState(68)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
		p.Match(fsmParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
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
	Note() antlr.TerminalNode
	Initial() antlr.TerminalNode
	Stereotype() IStereotypeContext
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

func (s *StateContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *StateContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *StateContext) Initial() antlr.TerminalNode {
	return s.GetToken(fsmParserInitial, 0)
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(77)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(80)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(83)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(85)
			p.Stereotype()
		}

	}
	{
		p.SetState(88)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4621639024112) != 0 {
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(89)
				p.State()
			}

		case 2:
			{
				p.SetState(90)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(91)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(92)
				p.Final()
			}

		case 5:
			{
				p.SetState(93)
				p.Terminate()
			}

		case 6:
			{
				p.SetState(94)
				p.Choice()
			}

		case 7:
			{
				p.SetState(95)
				p.Junction()
			}

		case 8:
			{
				p.SetState(96)
				p.Fork()
			}

		case 9:
			{
				p.SetState(97)
				p.Join()
			}

		case 10:
			{
				p.SetState(98)
				p.Point()
			}

		case 11:
			{
				p.SetState(99)
				p.History()
			}

		case 12:
			{
				p.SetState(100)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
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

// IParallelContext is an interface to support dynamic dispatch.
type IParallelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	Initial() antlr.TerminalNode
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

func (s *ParallelContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *ParallelContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *ParallelContext) Initial() antlr.TerminalNode {
	return s.GetToken(fsmParserInitial, 0)
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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(108)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(111)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(114)
		p.Match(fsmParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(117)
			p.Stereotype()
		}

	}
	{
		p.SetState(120)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4552919138816) != 0 {
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(121)
				p.Region()
			}

		case 2:
			{
				p.SetState(122)
				p.Point()
			}

		case 3:
			{
				p.SetState(123)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
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

// ISubmachineContext is an interface to support dynamic dispatch.
type ISubmachineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	Initial() antlr.TerminalNode
	Stereotype() IStereotypeContext
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

func (s *SubmachineContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
}

func (s *SubmachineContext) Note() antlr.TerminalNode {
	return s.GetToken(fsmParserNote, 0)
}

func (s *SubmachineContext) Initial() antlr.TerminalNode {
	return s.GetToken(fsmParserInitial, 0)
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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(131)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(134)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(137)
		p.Match(fsmParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(139)
			p.Stereotype()
		}

	}
	{
		p.SetState(142)
		p.Match(fsmParserT__1)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4552919138304) != 0 {
		{
			p.SetState(143)
			p.Event()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
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

// IFinalContext is an interface to support dynamic dispatch.
type IFinalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(151)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(154)
		p.Match(fsmParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserIdentifier {
		{
			p.SetState(156)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(159)
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
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(162)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(165)
		p.Match(fsmParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserIdentifier {
		{
			p.SetState(167)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(170)
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
	Identifier() antlr.TerminalNode
	Note() antlr.TerminalNode
	Stereotype() IStereotypeContext
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
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(173)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(176)
		p.Match(fsmParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(178)
			p.Stereotype()
		}

	}
	{
		p.SetState(181)
		p.Match(fsmParserT__1)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4466766396912) != 0 {
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(182)
				p.State()
			}

		case 2:
			{
				p.SetState(183)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(184)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(185)
				p.Final()
			}

		case 5:
			{
				p.SetState(186)
				p.Terminate()
			}

		case 6:
			{
				p.SetState(187)
				p.Choice()
			}

		case 7:
			{
				p.SetState(188)
				p.Junction()
			}

		case 8:
			{
				p.SetState(189)
				p.Fork()
			}

		case 9:
			{
				p.SetState(190)
				p.Join()
			}

		case 10:
			{
				p.SetState(191)
				p.History()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(197)
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

// IChoiceContext is an interface to support dynamic dispatch.
type IChoiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	AllBranch() []IBranchContext
	Branch(i int) IBranchContext
	Note() antlr.TerminalNode
	Initial() antlr.TerminalNode
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

func (s *ChoiceContext) Initial() antlr.TerminalNode {
	return s.GetToken(fsmParserInitial, 0)
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
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(199)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(202)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(205)
		p.Match(fsmParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(207)
			p.Stereotype()
		}

	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__1:
		{
			p.SetState(210)
			p.Match(fsmParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4415295586304) != 0 {
			{
				p.SetState(211)
				p.Branch()
			}

			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(217)
			p.Match(fsmParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__20, fsmParserT__25, fsmParserT__33, fsmParserNote:
		{
			p.SetState(218)
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
	Identifier() antlr.TerminalNode
	AllBranch() []IBranchContext
	Branch(i int) IBranchContext
	Note() antlr.TerminalNode
	Initial() antlr.TerminalNode
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

func (s *JunctionContext) Initial() antlr.TerminalNode {
	return s.GetToken(fsmParserInitial, 0)
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
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(221)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(224)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(227)
		p.Match(fsmParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(fsmParserIdentifier)
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

	if _la == fsmParserT__18 {
		{
			p.SetState(229)
			p.Stereotype()
		}

	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__1:
		{
			p.SetState(232)
			p.Match(fsmParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4415295586304) != 0 {
			{
				p.SetState(233)
				p.Branch()
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(239)
			p.Match(fsmParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__20, fsmParserT__25, fsmParserT__33, fsmParserNote:
		{
			p.SetState(240)
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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(243)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(246)
		p.Match(fsmParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(248)
			p.Stereotype()
		}

	}
	{
		p.SetState(251)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4415293489152) != 0 {
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
		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__25 {
			{
				p.SetState(255)
				p.Actions()
			}

		}
		{
			p.SetState(258)
			p.Goto_()
		}

		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(264)
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

// IJoinContext is an interface to support dynamic dispatch.
type IJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(266)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(269)
		p.Match(fsmParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(271)
			p.Stereotype()
		}

	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__25 {
		{
			p.SetState(274)
			p.Actions()
		}

	}
	{
		p.SetState(277)
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
	Identifier() antlr.TerminalNode
	Goto_() IGotoContext
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
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(279)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(282)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PointContext).kind = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == fsmParserT__13 || _la == fsmParserT__14) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PointContext).kind = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(283)
		p.Match(fsmParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(285)
			p.Stereotype()
		}

	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__25 {
		{
			p.SetState(288)
			p.Actions()
		}

	}
	{
		p.SetState(291)
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
	p.EnterRule(localctx, 24, fsmParserRULE_history)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(293)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(296)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*HistoryContext).kind = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == fsmParserT__16 || _la == fsmParserT__17) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*HistoryContext).kind = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__18 {
		{
			p.SetState(297)
			p.Stereotype()
		}

	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__25 {
			{
				p.SetState(300)
				p.Actions()
			}

		}
		{
			p.SetState(303)
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
	p.EnterRule(localctx, 26, fsmParserRULE_stereotype)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(fsmParserT__18)
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
	{
		p.SetState(308)
		p.Match(fsmParserT__19)
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
	p.EnterRule(localctx, 28, fsmParserRULE_branch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(310)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(313)
			p.Match(fsmParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(fsmParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(fsmParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(316)
			p.Guard()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__25 {
		{
			p.SetState(319)
			p.Actions()
		}

	}
	{
		p.SetState(322)
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
	Defer() antlr.TerminalNode
	Invariant() antlr.TerminalNode
	Guard() IGuardContext
	Trigger() ITriggerContext
	Goto_() IGotoContext
	Note() antlr.TerminalNode
	Identifier() antlr.TerminalNode

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

func (s *EventContext) Defer() antlr.TerminalNode {
	return s.GetToken(fsmParserDefer, 0)
}

func (s *EventContext) Invariant() antlr.TerminalNode {
	return s.GetToken(fsmParserInvariant, 0)
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
	p.EnterRule(localctx, 30, fsmParserRULE_event)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(324)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(327)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*EventContext).name = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16826368) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*EventContext).name = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(328)
			p.Actions()
		}

	case 2:
		{
			p.SetState(329)
			p.Match(fsmParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*EventContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(fsmParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(fsmParserDefer)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(333)
			p.Match(fsmParserInvariant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Guard()
		}

	case 4:
		{
			p.SetState(335)
			p.Trigger()
		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__20 {
			{
				p.SetState(336)
				p.Guard()
			}

		}
		{
			p.SetState(339)
			p.Actions()
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(340)
				p.Goto_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		{
			p.SetState(343)
			p.Trigger()
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__20 {
			{
				p.SetState(344)
				p.Guard()
			}

		}
		{
			p.SetState(347)
			p.Goto_()
		}

	case 6:
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__20 {
			{
				p.SetState(349)
				p.Guard()
			}

		}
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__25 {
			{
				p.SetState(352)
				p.Actions()
			}

		}
		{
			p.SetState(355)
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
	Identifier() antlr.TerminalNode
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

func (s *TriggerContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
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
	p.EnterRule(localctx, 32, fsmParserRULE_trigger)
	var _la int

	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__24:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(358)
			p.Match(fsmParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*TriggerContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(360)
			p.Match(fsmParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.Match(fsmParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)

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
			p.SetState(363)
			p.Match(fsmParserT__28)
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
	p.EnterRule(localctx, 34, fsmParserRULE_actions)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(fsmParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
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
	p.EnterRule(localctx, 36, fsmParserRULE_identifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__29 {
		{
			p.SetState(370)
			p.Match(fsmParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(376)
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
	p.EnterRule(localctx, 38, fsmParserRULE_guard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(fsmParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.Expression()
	}
	{
		p.SetState(379)
		p.Match(fsmParserT__22)
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
	p.EnterRule(localctx, 40, fsmParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
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
	p.EnterRule(localctx, 42, fsmParserRULE_or_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.And_expression()
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__30 {
		{
			p.SetState(384)
			p.Match(fsmParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.And_expression()
		}

		p.SetState(390)
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
	p.EnterRule(localctx, 44, fsmParserRULE_and_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Not_expression()
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__31 {
		{
			p.SetState(392)
			p.Match(fsmParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Not_expression()
		}

		p.SetState(398)
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
	p.EnterRule(localctx, 46, fsmParserRULE_not_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__32 {
		{
			p.SetState(399)
			p.Match(fsmParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(402)
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
	p.EnterRule(localctx, 48, fsmParserRULE_single_expression)
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__27:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(405)
			p.Match(fsmParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Expression()
		}
		{
			p.SetState(407)
			p.Match(fsmParserT__28)
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
	Local() antlr.TerminalNode

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

func (s *GotoContext) Local() antlr.TerminalNode {
	return s.GetToken(fsmParserLocal, 0)
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
	p.EnterRule(localctx, 50, fsmParserRULE_goto)
	var _la int

	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 71, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.Match(fsmParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			_la = p.GetTokenStream().LA(1)

			if !(_la == fsmParserT__6 || _la == fsmParserT__7) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(413)
			p.Match(fsmParserT__33)
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

		if _la == fsmParserLocal {
			{
				p.SetState(414)
				p.Match(fsmParserLocal)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(417)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(418)
			p.Match(fsmParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserLocal {
			{
				p.SetState(419)
				p.Match(fsmParserLocal)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserIdentifier {
			{
				p.SetState(422)
				p.Match(fsmParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(423)
				p.Match(fsmParserT__34)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(426)
			_la = p.GetTokenStream().LA(1)

			if !(_la == fsmParserT__16 || _la == fsmParserT__17) {
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
