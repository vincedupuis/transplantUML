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
		"'region'", "'choice'", "'junction'", "'fork'", "'join'", "'entry'",
		"'exit'", "'point'", "'<<'", "'>>'", "'['", "'else'", "']'", "'do'",
		"'on'", "'/'", "'after'", "'('", "')'", "','", "'or'", "'and'", "'not'",
		"'goto'", "'final'", "'terminate'", "'.'", "'H'", "'H*'", "'initial'",
		"'invariant'", "'defer'", "'local'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "Initial", "Invariant", "Defer", "Local", "Duration", "Identifier",
		"Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"fsm", "state", "parallel", "submachine", "region", "choice", "junction",
		"fork", "join", "point", "stereotype", "branch", "event", "trigger",
		"actions", "identifiers", "guard", "expression", "or_expression", "and_expression",
		"not_expression", "single_expression", "goto",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 336, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 59, 8, 0, 10, 0, 12, 0, 62, 9, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 3, 1, 68, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 73, 8, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 85, 8, 1, 10,
		1, 12, 1, 88, 9, 1, 1, 1, 1, 1, 1, 2, 3, 2, 93, 8, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 99, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 105, 8, 2, 10, 2, 12,
		2, 108, 9, 2, 1, 2, 1, 2, 1, 3, 3, 3, 113, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		118, 8, 3, 1, 3, 1, 3, 5, 3, 122, 8, 3, 10, 3, 12, 3, 125, 9, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 142, 8, 4, 10, 4, 12, 4, 145, 9, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 3, 5, 152, 8, 5, 1, 5, 1, 5, 5, 5, 156, 8, 5, 10, 5,
		12, 5, 159, 9, 5, 1, 5, 1, 5, 3, 5, 163, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6,
		168, 8, 6, 1, 6, 1, 6, 5, 6, 172, 8, 6, 10, 6, 12, 6, 175, 9, 6, 1, 6,
		1, 6, 3, 6, 179, 8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 184, 8, 7, 1, 7, 1, 7, 3,
		7, 188, 8, 7, 1, 7, 5, 7, 191, 8, 7, 10, 7, 12, 7, 194, 9, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 3, 8, 201, 8, 8, 1, 8, 3, 8, 204, 8, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 212, 8, 9, 1, 9, 3, 9, 215, 8, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 227,
		8, 11, 1, 11, 3, 11, 230, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 244, 8, 12, 1, 12,
		1, 12, 3, 12, 248, 8, 12, 1, 12, 1, 12, 3, 12, 252, 8, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 257, 8, 12, 1, 12, 3, 12, 260, 8, 12, 1, 12, 3, 12, 263,
		8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 271, 8, 13, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 279, 8, 15, 10, 15, 12, 15,
		282, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 5, 18, 293, 8, 18, 10, 18, 12, 18, 296, 9, 18, 1, 19, 1, 19, 1, 19,
		5, 19, 301, 8, 19, 10, 19, 12, 19, 304, 9, 19, 1, 20, 3, 20, 307, 8, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 316, 8, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 3, 22, 322, 8, 22, 1, 22, 1, 22, 1, 22, 3, 22,
		327, 8, 22, 1, 22, 1, 22, 3, 22, 331, 8, 22, 1, 22, 3, 22, 334, 8, 22,
		1, 22, 0, 0, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 0, 5, 1, 0, 12, 13, 2, 0, 12, 13, 20, 20,
		1, 0, 40, 41, 1, 0, 31, 32, 1, 0, 34, 35, 385, 0, 46, 1, 0, 0, 0, 2, 67,
		1, 0, 0, 0, 4, 92, 1, 0, 0, 0, 6, 112, 1, 0, 0, 0, 8, 128, 1, 0, 0, 0,
		10, 148, 1, 0, 0, 0, 12, 164, 1, 0, 0, 0, 14, 180, 1, 0, 0, 0, 16, 197,
		1, 0, 0, 0, 18, 207, 1, 0, 0, 0, 20, 218, 1, 0, 0, 0, 22, 226, 1, 0, 0,
		0, 24, 262, 1, 0, 0, 0, 26, 270, 1, 0, 0, 0, 28, 272, 1, 0, 0, 0, 30, 275,
		1, 0, 0, 0, 32, 283, 1, 0, 0, 0, 34, 287, 1, 0, 0, 0, 36, 289, 1, 0, 0,
		0, 38, 297, 1, 0, 0, 0, 40, 306, 1, 0, 0, 0, 42, 315, 1, 0, 0, 0, 44, 333,
		1, 0, 0, 0, 46, 47, 5, 1, 0, 0, 47, 48, 5, 41, 0, 0, 48, 60, 5, 2, 0, 0,
		49, 59, 3, 2, 1, 0, 50, 59, 3, 4, 2, 0, 51, 59, 3, 6, 3, 0, 52, 59, 3,
		10, 5, 0, 53, 59, 3, 12, 6, 0, 54, 59, 3, 14, 7, 0, 55, 59, 3, 16, 8, 0,
		56, 59, 3, 18, 9, 0, 57, 59, 3, 24, 12, 0, 58, 49, 1, 0, 0, 0, 58, 50,
		1, 0, 0, 0, 58, 51, 1, 0, 0, 0, 58, 52, 1, 0, 0, 0, 58, 53, 1, 0, 0, 0,
		58, 54, 1, 0, 0, 0, 58, 55, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1,
		0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61,
		63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 64, 5, 3, 0, 0, 64, 65, 5, 0, 0,
		1, 65, 1, 1, 0, 0, 0, 66, 68, 5, 36, 0, 0, 67, 66, 1, 0, 0, 0, 67, 68,
		1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 5, 4, 0, 0, 70, 72, 5, 41, 0, 0,
		71, 73, 3, 20, 10, 0, 72, 71, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 1,
		0, 0, 0, 74, 86, 5, 2, 0, 0, 75, 85, 3, 2, 1, 0, 76, 85, 3, 4, 2, 0, 77,
		85, 3, 6, 3, 0, 78, 85, 3, 10, 5, 0, 79, 85, 3, 12, 6, 0, 80, 85, 3, 14,
		7, 0, 81, 85, 3, 16, 8, 0, 82, 85, 3, 18, 9, 0, 83, 85, 3, 24, 12, 0, 84,
		75, 1, 0, 0, 0, 84, 76, 1, 0, 0, 0, 84, 77, 1, 0, 0, 0, 84, 78, 1, 0, 0,
		0, 84, 79, 1, 0, 0, 0, 84, 80, 1, 0, 0, 0, 84, 81, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		86, 87, 1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 5,
		3, 0, 0, 90, 3, 1, 0, 0, 0, 91, 93, 5, 36, 0, 0, 92, 91, 1, 0, 0, 0, 92,
		93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 5, 5, 0, 0, 95, 96, 5, 4, 0,
		0, 96, 98, 5, 41, 0, 0, 97, 99, 3, 20, 10, 0, 98, 97, 1, 0, 0, 0, 98, 99,
		1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 106, 5, 2, 0, 0, 101, 105, 3, 8,
		4, 0, 102, 105, 3, 18, 9, 0, 103, 105, 3, 24, 12, 0, 104, 101, 1, 0, 0,
		0, 104, 102, 1, 0, 0, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 106,
		1, 0, 0, 0, 109, 110, 5, 3, 0, 0, 110, 5, 1, 0, 0, 0, 111, 113, 5, 36,
		0, 0, 112, 111, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0,
		114, 115, 5, 6, 0, 0, 115, 117, 5, 41, 0, 0, 116, 118, 3, 20, 10, 0, 117,
		116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 123,
		5, 2, 0, 0, 120, 122, 3, 24, 12, 0, 121, 120, 1, 0, 0, 0, 122, 125, 1,
		0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0,
		0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 3, 0, 0, 127, 7, 1, 0, 0, 0, 128,
		129, 5, 7, 0, 0, 129, 131, 5, 41, 0, 0, 130, 132, 3, 20, 10, 0, 131, 130,
		1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 143, 5, 2,
		0, 0, 134, 142, 3, 2, 1, 0, 135, 142, 3, 4, 2, 0, 136, 142, 3, 6, 3, 0,
		137, 142, 3, 10, 5, 0, 138, 142, 3, 12, 6, 0, 139, 142, 3, 14, 7, 0, 140,
		142, 3, 16, 8, 0, 141, 134, 1, 0, 0, 0, 141, 135, 1, 0, 0, 0, 141, 136,
		1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 141, 138, 1, 0, 0, 0, 141, 139, 1, 0,
		0, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0,
		143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146,
		147, 5, 3, 0, 0, 147, 9, 1, 0, 0, 0, 148, 149, 5, 8, 0, 0, 149, 151, 5,
		41, 0, 0, 150, 152, 3, 20, 10, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0,
		0, 0, 152, 162, 1, 0, 0, 0, 153, 157, 5, 2, 0, 0, 154, 156, 3, 22, 11,
		0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157,
		158, 1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 163,
		5, 3, 0, 0, 161, 163, 3, 22, 11, 0, 162, 153, 1, 0, 0, 0, 162, 161, 1,
		0, 0, 0, 163, 11, 1, 0, 0, 0, 164, 165, 5, 9, 0, 0, 165, 167, 5, 41, 0,
		0, 166, 168, 3, 20, 10, 0, 167, 166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0,
		168, 178, 1, 0, 0, 0, 169, 173, 5, 2, 0, 0, 170, 172, 3, 22, 11, 0, 171,
		170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174,
		1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 179, 5, 3,
		0, 0, 177, 179, 3, 22, 11, 0, 178, 169, 1, 0, 0, 0, 178, 177, 1, 0, 0,
		0, 179, 13, 1, 0, 0, 0, 180, 181, 5, 10, 0, 0, 181, 183, 5, 41, 0, 0, 182,
		184, 3, 20, 10, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185,
		1, 0, 0, 0, 185, 192, 5, 2, 0, 0, 186, 188, 3, 28, 14, 0, 187, 186, 1,
		0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 3, 44, 22,
		0, 190, 187, 1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192,
		193, 1, 0, 0, 0, 193, 195, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 196,
		5, 3, 0, 0, 196, 15, 1, 0, 0, 0, 197, 198, 5, 11, 0, 0, 198, 200, 5, 41,
		0, 0, 199, 201, 3, 20, 10, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0,
		0, 201, 203, 1, 0, 0, 0, 202, 204, 3, 28, 14, 0, 203, 202, 1, 0, 0, 0,
		203, 204, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 3, 44, 22, 0, 206,
		17, 1, 0, 0, 0, 207, 208, 7, 0, 0, 0, 208, 209, 5, 14, 0, 0, 209, 211,
		5, 41, 0, 0, 210, 212, 3, 20, 10, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1,
		0, 0, 0, 212, 214, 1, 0, 0, 0, 213, 215, 3, 28, 14, 0, 214, 213, 1, 0,
		0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 3, 44, 22,
		0, 217, 19, 1, 0, 0, 0, 218, 219, 5, 15, 0, 0, 219, 220, 5, 41, 0, 0, 220,
		221, 5, 16, 0, 0, 221, 21, 1, 0, 0, 0, 222, 223, 5, 17, 0, 0, 223, 224,
		5, 18, 0, 0, 224, 227, 5, 19, 0, 0, 225, 227, 3, 32, 16, 0, 226, 222, 1,
		0, 0, 0, 226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 229, 1, 0, 0,
		0, 228, 230, 3, 28, 14, 0, 229, 228, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0,
		230, 231, 1, 0, 0, 0, 231, 232, 3, 44, 22, 0, 232, 23, 1, 0, 0, 0, 233,
		234, 7, 1, 0, 0, 234, 263, 3, 28, 14, 0, 235, 236, 5, 21, 0, 0, 236, 237,
		5, 41, 0, 0, 237, 238, 5, 22, 0, 0, 238, 263, 5, 38, 0, 0, 239, 240, 5,
		37, 0, 0, 240, 263, 3, 32, 16, 0, 241, 243, 3, 26, 13, 0, 242, 244, 3,
		32, 16, 0, 243, 242, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 245, 1, 0,
		0, 0, 245, 247, 3, 28, 14, 0, 246, 248, 3, 44, 22, 0, 247, 246, 1, 0, 0,
		0, 247, 248, 1, 0, 0, 0, 248, 263, 1, 0, 0, 0, 249, 251, 3, 26, 13, 0,
		250, 252, 3, 32, 16, 0, 251, 250, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252,
		253, 1, 0, 0, 0, 253, 254, 3, 44, 22, 0, 254, 263, 1, 0, 0, 0, 255, 257,
		3, 32, 16, 0, 256, 255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 259, 1,
		0, 0, 0, 258, 260, 3, 28, 14, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1, 0,
		0, 0, 260, 261, 1, 0, 0, 0, 261, 263, 3, 44, 22, 0, 262, 233, 1, 0, 0,
		0, 262, 235, 1, 0, 0, 0, 262, 239, 1, 0, 0, 0, 262, 241, 1, 0, 0, 0, 262,
		249, 1, 0, 0, 0, 262, 256, 1, 0, 0, 0, 263, 25, 1, 0, 0, 0, 264, 265, 5,
		21, 0, 0, 265, 271, 5, 41, 0, 0, 266, 267, 5, 23, 0, 0, 267, 268, 5, 24,
		0, 0, 268, 269, 7, 2, 0, 0, 269, 271, 5, 25, 0, 0, 270, 264, 1, 0, 0, 0,
		270, 266, 1, 0, 0, 0, 271, 27, 1, 0, 0, 0, 272, 273, 5, 22, 0, 0, 273,
		274, 3, 30, 15, 0, 274, 29, 1, 0, 0, 0, 275, 280, 5, 41, 0, 0, 276, 277,
		5, 26, 0, 0, 277, 279, 5, 41, 0, 0, 278, 276, 1, 0, 0, 0, 279, 282, 1,
		0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 31, 1, 0, 0,
		0, 282, 280, 1, 0, 0, 0, 283, 284, 5, 17, 0, 0, 284, 285, 3, 34, 17, 0,
		285, 286, 5, 19, 0, 0, 286, 33, 1, 0, 0, 0, 287, 288, 3, 36, 18, 0, 288,
		35, 1, 0, 0, 0, 289, 294, 3, 38, 19, 0, 290, 291, 5, 27, 0, 0, 291, 293,
		3, 38, 19, 0, 292, 290, 1, 0, 0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1,
		0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 37, 1, 0, 0, 0, 296, 294, 1, 0, 0,
		0, 297, 302, 3, 40, 20, 0, 298, 299, 5, 28, 0, 0, 299, 301, 3, 40, 20,
		0, 300, 298, 1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302,
		303, 1, 0, 0, 0, 303, 39, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 307, 5,
		29, 0, 0, 306, 305, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 1, 0, 0,
		0, 308, 309, 3, 42, 21, 0, 309, 41, 1, 0, 0, 0, 310, 316, 5, 41, 0, 0,
		311, 312, 5, 24, 0, 0, 312, 313, 3, 34, 17, 0, 313, 314, 5, 25, 0, 0, 314,
		316, 1, 0, 0, 0, 315, 310, 1, 0, 0, 0, 315, 311, 1, 0, 0, 0, 316, 43, 1,
		0, 0, 0, 317, 318, 5, 30, 0, 0, 318, 334, 7, 3, 0, 0, 319, 321, 5, 30,
		0, 0, 320, 322, 5, 39, 0, 0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0,
		322, 323, 1, 0, 0, 0, 323, 334, 5, 41, 0, 0, 324, 326, 5, 30, 0, 0, 325,
		327, 5, 39, 0, 0, 326, 325, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 330,
		1, 0, 0, 0, 328, 329, 5, 41, 0, 0, 329, 331, 5, 33, 0, 0, 330, 328, 1,
		0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 334, 7, 4, 0,
		0, 333, 317, 1, 0, 0, 0, 333, 319, 1, 0, 0, 0, 333, 324, 1, 0, 0, 0, 334,
		45, 1, 0, 0, 0, 47, 58, 60, 67, 72, 84, 86, 92, 98, 104, 106, 112, 117,
		123, 131, 141, 143, 151, 157, 162, 167, 173, 178, 183, 187, 192, 200, 203,
		211, 214, 226, 229, 243, 247, 251, 256, 259, 262, 270, 280, 294, 302, 306,
		315, 321, 326, 330, 333,
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
	fsmParserComment    = 42
	fsmParserBlank      = 43
)

// fsmParser rules.
const (
	fsmParserRULE_fsm               = 0
	fsmParserRULE_state             = 1
	fsmParserRULE_parallel          = 2
	fsmParserRULE_submachine        = 3
	fsmParserRULE_region            = 4
	fsmParserRULE_choice            = 5
	fsmParserRULE_junction          = 6
	fsmParserRULE_fork              = 7
	fsmParserRULE_join              = 8
	fsmParserRULE_point             = 9
	fsmParserRULE_stereotype        = 10
	fsmParserRULE_branch            = 11
	fsmParserRULE_event             = 12
	fsmParserRULE_trigger           = 13
	fsmParserRULE_actions           = 14
	fsmParserRULE_identifiers       = 15
	fsmParserRULE_guard             = 16
	fsmParserRULE_expression        = 17
	fsmParserRULE_or_expression     = 18
	fsmParserRULE_and_expression    = 19
	fsmParserRULE_not_expression    = 20
	fsmParserRULE_single_expression = 21
	fsmParserRULE_goto              = 22
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
	AllParallel() []IParallelContext
	Parallel(i int) IParallelContext
	AllSubmachine() []ISubmachineContext
	Submachine(i int) ISubmachineContext
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
	{
		p.SetState(46)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(fsmParserT__1)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&207248047984) != 0 {
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(49)
				p.State()
			}

		case 2:
			{
				p.SetState(50)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(51)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(52)
				p.Choice()
			}

		case 5:
			{
				p.SetState(53)
				p.Junction()
			}

		case 6:
			{
				p.SetState(54)
				p.Fork()
			}

		case 7:
			{
				p.SetState(55)
				p.Join()
			}

		case 8:
			{
				p.SetState(56)
				p.Point()
			}

		case 9:
			{
				p.SetState(57)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(fsmParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
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
	Stereotype() IStereotypeContext
	AllState() []IStateContext
	State(i int) IStateContext
	AllParallel() []IParallelContext
	Parallel(i int) IParallelContext
	AllSubmachine() []ISubmachineContext
	Submachine(i int) ISubmachineContext
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(66)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(69)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(71)
			p.Stereotype()
		}

	}
	{
		p.SetState(74)
		p.Match(fsmParserT__1)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&207248047984) != 0 {
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(75)
				p.State()
			}

		case 2:
			{
				p.SetState(76)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(77)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(78)
				p.Choice()
			}

		case 5:
			{
				p.SetState(79)
				p.Junction()
			}

		case 6:
			{
				p.SetState(80)
				p.Fork()
			}

		case 7:
			{
				p.SetState(81)
				p.Join()
			}

		case 8:
			{
				p.SetState(82)
				p.Point()
			}

		case 9:
			{
				p.SetState(83)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
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
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(91)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(94)
		p.Match(fsmParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(97)
			p.Stereotype()
		}

	}
	{
		p.SetState(100)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&138528567424) != 0 {
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(101)
				p.Region()
			}

		case 2:
			{
				p.SetState(102)
				p.Point()
			}

		case 3:
			{
				p.SetState(103)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
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
		p.Match(fsmParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(116)
			p.Stereotype()
		}

	}
	{
		p.SetState(119)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&138528567296) != 0 {
		{
			p.SetState(120)
			p.Event()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
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

// IRegionContext is an interface to support dynamic dispatch.
type IRegionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Stereotype() IStereotypeContext
	AllState() []IStateContext
	State(i int) IStateContext
	AllParallel() []IParallelContext
	Parallel(i int) IParallelContext
	AllSubmachine() []ISubmachineContext
	Submachine(i int) ISubmachineContext
	AllChoice() []IChoiceContext
	Choice(i int) IChoiceContext
	AllJunction() []IJunctionContext
	Junction(i int) IJunctionContext
	AllFork() []IForkContext
	Fork(i int) IForkContext
	AllJoin() []IJoinContext
	Join(i int) IJoinContext

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
	p.EnterRule(localctx, 8, fsmParserRULE_region)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(fsmParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(130)
			p.Stereotype()
		}

	}
	{
		p.SetState(133)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719480688) != 0 {
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(134)
				p.State()
			}

		case 2:
			{
				p.SetState(135)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(136)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(137)
				p.Choice()
			}

		case 5:
			{
				p.SetState(138)
				p.Junction()
			}

		case 6:
			{
				p.SetState(139)
				p.Fork()
			}

		case 7:
			{
				p.SetState(140)
				p.Join()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
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
	p.EnterRule(localctx, 10, fsmParserRULE_choice)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(fsmParserT__7)
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
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(150)
			p.Stereotype()
		}

	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__1:
		{
			p.SetState(153)
			p.Match(fsmParserT__1)
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

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1078067200) != 0 {
			{
				p.SetState(154)
				p.Branch()
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(160)
			p.Match(fsmParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__16, fsmParserT__21, fsmParserT__29:
		{
			p.SetState(161)
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
	p.EnterRule(localctx, 12, fsmParserRULE_junction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(fsmParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(166)
			p.Stereotype()
		}

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__1:
		{
			p.SetState(169)
			p.Match(fsmParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1078067200) != 0 {
			{
				p.SetState(170)
				p.Branch()
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(176)
			p.Match(fsmParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__16, fsmParserT__21, fsmParserT__29:
		{
			p.SetState(177)
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
	p.EnterRule(localctx, 14, fsmParserRULE_fork)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(fsmParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(fsmParserIdentifier)
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

	if _la == fsmParserT__14 {
		{
			p.SetState(182)
			p.Stereotype()
		}

	}
	{
		p.SetState(185)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__21 || _la == fsmParserT__29 {
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__21 {
			{
				p.SetState(186)
				p.Actions()
			}

		}
		{
			p.SetState(189)
			p.Goto_()
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(195)
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
	p.EnterRule(localctx, 16, fsmParserRULE_join)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(fsmParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(199)
			p.Stereotype()
		}

	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__21 {
		{
			p.SetState(202)
			p.Actions()
		}

	}
	{
		p.SetState(205)
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
	p.EnterRule(localctx, 18, fsmParserRULE_point)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PointContext).kind = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == fsmParserT__11 || _la == fsmParserT__12) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PointContext).kind = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(208)
		p.Match(fsmParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(210)
			p.Stereotype()
		}

	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__21 {
		{
			p.SetState(213)
			p.Actions()
		}

	}
	{
		p.SetState(216)
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
	p.EnterRule(localctx, 20, fsmParserRULE_stereotype)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(fsmParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(fsmParserT__15)
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
	p.EnterRule(localctx, 22, fsmParserRULE_branch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(222)
			p.Match(fsmParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(fsmParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(fsmParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(225)
			p.Guard()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__21 {
		{
			p.SetState(228)
			p.Actions()
		}

	}
	{
		p.SetState(231)
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
	Identifier() antlr.TerminalNode
	Invariant() antlr.TerminalNode
	Guard() IGuardContext
	Trigger() ITriggerContext
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

func (s *EventContext) Defer() antlr.TerminalNode {
	return s.GetToken(fsmParserDefer, 0)
}

func (s *EventContext) Identifier() antlr.TerminalNode {
	return s.GetToken(fsmParserIdentifier, 0)
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
	p.EnterRule(localctx, 24, fsmParserRULE_event)
	var _la int

	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*EventContext).name = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1060864) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*EventContext).name = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(234)
			p.Actions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(fsmParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*EventContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(fsmParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(fsmParserDefer)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.Match(fsmParserInvariant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Guard()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(241)
			p.Trigger()
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__16 {
			{
				p.SetState(242)
				p.Guard()
			}

		}
		{
			p.SetState(245)
			p.Actions()
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(246)
				p.Goto_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(249)
			p.Trigger()
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__16 {
			{
				p.SetState(250)
				p.Guard()
			}

		}
		{
			p.SetState(253)
			p.Goto_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__16 {
			{
				p.SetState(255)
				p.Guard()
			}

		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__21 {
			{
				p.SetState(258)
				p.Actions()
			}

		}
		{
			p.SetState(261)
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
	p.EnterRule(localctx, 26, fsmParserRULE_trigger)
	var _la int

	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Match(fsmParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*TriggerContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(fsmParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(fsmParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)

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
			p.SetState(269)
			p.Match(fsmParserT__24)
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
	p.EnterRule(localctx, 28, fsmParserRULE_actions)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(fsmParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
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
	p.EnterRule(localctx, 30, fsmParserRULE_identifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__25 {
		{
			p.SetState(276)
			p.Match(fsmParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(282)
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
	p.EnterRule(localctx, 32, fsmParserRULE_guard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(fsmParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Expression()
	}
	{
		p.SetState(285)
		p.Match(fsmParserT__18)
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
	p.EnterRule(localctx, 34, fsmParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
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
	p.EnterRule(localctx, 36, fsmParserRULE_or_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.And_expression()
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__26 {
		{
			p.SetState(290)
			p.Match(fsmParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.And_expression()
		}

		p.SetState(296)
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
	p.EnterRule(localctx, 38, fsmParserRULE_and_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Not_expression()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__27 {
		{
			p.SetState(298)
			p.Match(fsmParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Not_expression()
		}

		p.SetState(304)
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
	p.EnterRule(localctx, 40, fsmParserRULE_not_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__28 {
		{
			p.SetState(305)
			p.Match(fsmParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(308)
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
	p.EnterRule(localctx, 42, fsmParserRULE_single_expression)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(310)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(311)
			p.Match(fsmParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Expression()
		}
		{
			p.SetState(313)
			p.Match(fsmParserT__24)
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
	p.EnterRule(localctx, 44, fsmParserRULE_goto)
	var _la int

	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Match(fsmParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			_la = p.GetTokenStream().LA(1)

			if !(_la == fsmParserT__30 || _la == fsmParserT__31) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.Match(fsmParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserLocal {
			{
				p.SetState(320)
				p.Match(fsmParserLocal)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.Match(fsmParserT__29)
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

		if _la == fsmParserLocal {
			{
				p.SetState(325)
				p.Match(fsmParserLocal)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserIdentifier {
			{
				p.SetState(328)
				p.Match(fsmParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(329)
				p.Match(fsmParserT__32)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(332)
			_la = p.GetTokenStream().LA(1)

			if !(_la == fsmParserT__33 || _la == fsmParserT__34) {
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
