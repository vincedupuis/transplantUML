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
		"Note", "Comment", "Blank",
	}
	staticData.RuleNames = []string{
		"fsm", "state", "parallel", "submachine", "region", "choice", "junction",
		"fork", "join", "point", "stereotype", "branch", "event", "trigger",
		"actions", "identifiers", "guard", "expression", "or_expression", "and_expression",
		"not_expression", "single_expression", "goto",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 375, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 3, 0, 48, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 62, 8, 0, 10, 0,
		12, 0, 65, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 3, 1, 71, 8, 1, 1, 1, 3, 1, 74,
		8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 79, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 91, 8, 1, 10, 1, 12, 1, 94, 9, 1, 1, 1,
		1, 1, 1, 2, 3, 2, 99, 8, 2, 1, 2, 3, 2, 102, 8, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 108, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 114, 8, 2, 10, 2, 12,
		2, 117, 9, 2, 1, 2, 1, 2, 1, 3, 3, 3, 122, 8, 3, 1, 3, 3, 3, 125, 8, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 130, 8, 3, 1, 3, 1, 3, 5, 3, 134, 8, 3, 10, 3,
		12, 3, 137, 9, 3, 1, 3, 1, 3, 1, 4, 3, 4, 142, 8, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 147, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		157, 8, 4, 10, 4, 12, 4, 160, 9, 4, 1, 4, 1, 4, 1, 5, 3, 5, 165, 8, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 170, 8, 5, 1, 5, 1, 5, 5, 5, 174, 8, 5, 10, 5,
		12, 5, 177, 9, 5, 1, 5, 1, 5, 3, 5, 181, 8, 5, 1, 6, 3, 6, 184, 8, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 189, 8, 6, 1, 6, 1, 6, 5, 6, 193, 8, 6, 10, 6, 12,
		6, 196, 9, 6, 1, 6, 1, 6, 3, 6, 200, 8, 6, 1, 7, 3, 7, 203, 8, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 208, 8, 7, 1, 7, 1, 7, 3, 7, 212, 8, 7, 1, 7, 3, 7, 215,
		8, 7, 1, 7, 5, 7, 218, 8, 7, 10, 7, 12, 7, 221, 9, 7, 1, 7, 1, 7, 1, 8,
		3, 8, 226, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 231, 8, 8, 1, 8, 3, 8, 234, 8,
		8, 1, 8, 1, 8, 1, 9, 3, 9, 239, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 245,
		8, 9, 1, 9, 3, 9, 248, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 3, 11, 257, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 263, 8, 11, 1,
		11, 3, 11, 266, 8, 11, 1, 11, 1, 11, 1, 12, 3, 12, 271, 8, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 283,
		8, 12, 1, 12, 1, 12, 3, 12, 287, 8, 12, 1, 12, 1, 12, 3, 12, 291, 8, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 296, 8, 12, 1, 12, 3, 12, 299, 8, 12, 1, 12,
		3, 12, 302, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 310,
		8, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 318, 8, 15, 10,
		15, 12, 15, 321, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 5, 18, 332, 8, 18, 10, 18, 12, 18, 335, 9, 18, 1, 19, 1,
		19, 1, 19, 5, 19, 340, 8, 19, 10, 19, 12, 19, 343, 9, 19, 1, 20, 3, 20,
		346, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 355,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 361, 8, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 366, 8, 22, 1, 22, 1, 22, 3, 22, 370, 8, 22, 1, 22, 3, 22, 373,
		8, 22, 1, 22, 0, 0, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 5, 1, 0, 12, 13, 2, 0, 12, 13,
		20, 20, 1, 0, 40, 41, 1, 0, 31, 32, 1, 0, 34, 35, 437, 0, 47, 1, 0, 0,
		0, 2, 70, 1, 0, 0, 0, 4, 98, 1, 0, 0, 0, 6, 121, 1, 0, 0, 0, 8, 141, 1,
		0, 0, 0, 10, 164, 1, 0, 0, 0, 12, 183, 1, 0, 0, 0, 14, 202, 1, 0, 0, 0,
		16, 225, 1, 0, 0, 0, 18, 238, 1, 0, 0, 0, 20, 251, 1, 0, 0, 0, 22, 256,
		1, 0, 0, 0, 24, 270, 1, 0, 0, 0, 26, 309, 1, 0, 0, 0, 28, 311, 1, 0, 0,
		0, 30, 314, 1, 0, 0, 0, 32, 322, 1, 0, 0, 0, 34, 326, 1, 0, 0, 0, 36, 328,
		1, 0, 0, 0, 38, 336, 1, 0, 0, 0, 40, 345, 1, 0, 0, 0, 42, 354, 1, 0, 0,
		0, 44, 372, 1, 0, 0, 0, 46, 48, 5, 42, 0, 0, 47, 46, 1, 0, 0, 0, 47, 48,
		1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 5, 1, 0, 0, 50, 51, 5, 41, 0, 0,
		51, 63, 5, 2, 0, 0, 52, 62, 3, 2, 1, 0, 53, 62, 3, 4, 2, 0, 54, 62, 3,
		6, 3, 0, 55, 62, 3, 10, 5, 0, 56, 62, 3, 12, 6, 0, 57, 62, 3, 14, 7, 0,
		58, 62, 3, 16, 8, 0, 59, 62, 3, 18, 9, 0, 60, 62, 3, 24, 12, 0, 61, 52,
		1, 0, 0, 0, 61, 53, 1, 0, 0, 0, 61, 54, 1, 0, 0, 0, 61, 55, 1, 0, 0, 0,
		61, 56, 1, 0, 0, 0, 61, 57, 1, 0, 0, 0, 61, 58, 1, 0, 0, 0, 61, 59, 1,
		0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63,
		64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 67, 5, 3, 0,
		0, 67, 68, 5, 0, 0, 1, 68, 1, 1, 0, 0, 0, 69, 71, 5, 42, 0, 0, 70, 69,
		1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 74, 5, 36, 0, 0,
		73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 5,
		4, 0, 0, 76, 78, 5, 41, 0, 0, 77, 79, 3, 20, 10, 0, 78, 77, 1, 0, 0, 0,
		78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 92, 5, 2, 0, 0, 81, 91, 3,
		2, 1, 0, 82, 91, 3, 4, 2, 0, 83, 91, 3, 6, 3, 0, 84, 91, 3, 10, 5, 0, 85,
		91, 3, 12, 6, 0, 86, 91, 3, 14, 7, 0, 87, 91, 3, 16, 8, 0, 88, 91, 3, 18,
		9, 0, 89, 91, 3, 24, 12, 0, 90, 81, 1, 0, 0, 0, 90, 82, 1, 0, 0, 0, 90,
		83, 1, 0, 0, 0, 90, 84, 1, 0, 0, 0, 90, 85, 1, 0, 0, 0, 90, 86, 1, 0, 0,
		0, 90, 87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94,
		1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 95, 96, 5, 3, 0, 0, 96, 3, 1, 0, 0, 0, 97, 99, 5, 42,
		0, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100,
		102, 5, 36, 0, 0, 101, 100, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103,
		1, 0, 0, 0, 103, 104, 5, 5, 0, 0, 104, 105, 5, 4, 0, 0, 105, 107, 5, 41,
		0, 0, 106, 108, 3, 20, 10, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0,
		0, 108, 109, 1, 0, 0, 0, 109, 115, 5, 2, 0, 0, 110, 114, 3, 8, 4, 0, 111,
		114, 3, 18, 9, 0, 112, 114, 3, 24, 12, 0, 113, 110, 1, 0, 0, 0, 113, 111,
		1, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0,
		0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0,
		118, 119, 5, 3, 0, 0, 119, 5, 1, 0, 0, 0, 120, 122, 5, 42, 0, 0, 121, 120,
		1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 125, 5, 36,
		0, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0,
		126, 127, 5, 6, 0, 0, 127, 129, 5, 41, 0, 0, 128, 130, 3, 20, 10, 0, 129,
		128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 135,
		5, 2, 0, 0, 132, 134, 3, 24, 12, 0, 133, 132, 1, 0, 0, 0, 134, 137, 1,
		0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 138, 1, 0, 0,
		0, 137, 135, 1, 0, 0, 0, 138, 139, 5, 3, 0, 0, 139, 7, 1, 0, 0, 0, 140,
		142, 5, 42, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143,
		1, 0, 0, 0, 143, 144, 5, 7, 0, 0, 144, 146, 5, 41, 0, 0, 145, 147, 3, 20,
		10, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0,
		148, 158, 5, 2, 0, 0, 149, 157, 3, 2, 1, 0, 150, 157, 3, 4, 2, 0, 151,
		157, 3, 6, 3, 0, 152, 157, 3, 10, 5, 0, 153, 157, 3, 12, 6, 0, 154, 157,
		3, 14, 7, 0, 155, 157, 3, 16, 8, 0, 156, 149, 1, 0, 0, 0, 156, 150, 1,
		0, 0, 0, 156, 151, 1, 0, 0, 0, 156, 152, 1, 0, 0, 0, 156, 153, 1, 0, 0,
		0, 156, 154, 1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 161, 1, 0, 0, 0, 160, 158,
		1, 0, 0, 0, 161, 162, 5, 3, 0, 0, 162, 9, 1, 0, 0, 0, 163, 165, 5, 42,
		0, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0,
		166, 167, 5, 8, 0, 0, 167, 169, 5, 41, 0, 0, 168, 170, 3, 20, 10, 0, 169,
		168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 180, 1, 0, 0, 0, 171, 175,
		5, 2, 0, 0, 172, 174, 3, 22, 11, 0, 173, 172, 1, 0, 0, 0, 174, 177, 1,
		0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 1, 0, 0,
		0, 177, 175, 1, 0, 0, 0, 178, 181, 5, 3, 0, 0, 179, 181, 3, 22, 11, 0,
		180, 171, 1, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 11, 1, 0, 0, 0, 182, 184,
		5, 42, 0, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0,
		0, 0, 185, 186, 5, 9, 0, 0, 186, 188, 5, 41, 0, 0, 187, 189, 3, 20, 10,
		0, 188, 187, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 199, 1, 0, 0, 0, 190,
		194, 5, 2, 0, 0, 191, 193, 3, 22, 11, 0, 192, 191, 1, 0, 0, 0, 193, 196,
		1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0,
		0, 0, 196, 194, 1, 0, 0, 0, 197, 200, 5, 3, 0, 0, 198, 200, 3, 22, 11,
		0, 199, 190, 1, 0, 0, 0, 199, 198, 1, 0, 0, 0, 200, 13, 1, 0, 0, 0, 201,
		203, 5, 42, 0, 0, 202, 201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204,
		1, 0, 0, 0, 204, 205, 5, 10, 0, 0, 205, 207, 5, 41, 0, 0, 206, 208, 3,
		20, 10, 0, 207, 206, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 1, 0,
		0, 0, 209, 219, 5, 2, 0, 0, 210, 212, 5, 42, 0, 0, 211, 210, 1, 0, 0, 0,
		211, 212, 1, 0, 0, 0, 212, 214, 1, 0, 0, 0, 213, 215, 3, 28, 14, 0, 214,
		213, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 218,
		3, 44, 22, 0, 217, 211, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1,
		0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 222, 1, 0, 0, 0, 221, 219, 1, 0, 0,
		0, 222, 223, 5, 3, 0, 0, 223, 15, 1, 0, 0, 0, 224, 226, 5, 42, 0, 0, 225,
		224, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228,
		5, 11, 0, 0, 228, 230, 5, 41, 0, 0, 229, 231, 3, 20, 10, 0, 230, 229, 1,
		0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232, 234, 3, 28, 14,
		0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235,
		236, 3, 44, 22, 0, 236, 17, 1, 0, 0, 0, 237, 239, 5, 42, 0, 0, 238, 237,
		1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 7, 0,
		0, 0, 241, 242, 5, 14, 0, 0, 242, 244, 5, 41, 0, 0, 243, 245, 3, 20, 10,
		0, 244, 243, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 247, 1, 0, 0, 0, 246,
		248, 3, 28, 14, 0, 247, 246, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 250, 3, 44, 22, 0, 250, 19, 1, 0, 0, 0, 251, 252, 5, 15,
		0, 0, 252, 253, 5, 41, 0, 0, 253, 254, 5, 16, 0, 0, 254, 21, 1, 0, 0, 0,
		255, 257, 5, 42, 0, 0, 256, 255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257,
		262, 1, 0, 0, 0, 258, 259, 5, 17, 0, 0, 259, 260, 5, 18, 0, 0, 260, 263,
		5, 19, 0, 0, 261, 263, 3, 32, 16, 0, 262, 258, 1, 0, 0, 0, 262, 261, 1,
		0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0, 0, 0, 264, 266, 3, 28, 14,
		0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267,
		268, 3, 44, 22, 0, 268, 23, 1, 0, 0, 0, 269, 271, 5, 42, 0, 0, 270, 269,
		1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 301, 1, 0, 0, 0, 272, 273, 7, 1,
		0, 0, 273, 302, 3, 28, 14, 0, 274, 275, 5, 21, 0, 0, 275, 276, 5, 41, 0,
		0, 276, 277, 5, 22, 0, 0, 277, 302, 5, 38, 0, 0, 278, 279, 5, 37, 0, 0,
		279, 302, 3, 32, 16, 0, 280, 282, 3, 26, 13, 0, 281, 283, 3, 32, 16, 0,
		282, 281, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284,
		286, 3, 28, 14, 0, 285, 287, 3, 44, 22, 0, 286, 285, 1, 0, 0, 0, 286, 287,
		1, 0, 0, 0, 287, 302, 1, 0, 0, 0, 288, 290, 3, 26, 13, 0, 289, 291, 3,
		32, 16, 0, 290, 289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 293, 3, 44, 22, 0, 293, 302, 1, 0, 0, 0, 294, 296, 3, 32, 16,
		0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 1, 0, 0, 0, 297,
		299, 3, 28, 14, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 300,
		1, 0, 0, 0, 300, 302, 3, 44, 22, 0, 301, 272, 1, 0, 0, 0, 301, 274, 1,
		0, 0, 0, 301, 278, 1, 0, 0, 0, 301, 280, 1, 0, 0, 0, 301, 288, 1, 0, 0,
		0, 301, 295, 1, 0, 0, 0, 302, 25, 1, 0, 0, 0, 303, 304, 5, 21, 0, 0, 304,
		310, 5, 41, 0, 0, 305, 306, 5, 23, 0, 0, 306, 307, 5, 24, 0, 0, 307, 308,
		7, 2, 0, 0, 308, 310, 5, 25, 0, 0, 309, 303, 1, 0, 0, 0, 309, 305, 1, 0,
		0, 0, 310, 27, 1, 0, 0, 0, 311, 312, 5, 22, 0, 0, 312, 313, 3, 30, 15,
		0, 313, 29, 1, 0, 0, 0, 314, 319, 5, 41, 0, 0, 315, 316, 5, 26, 0, 0, 316,
		318, 5, 41, 0, 0, 317, 315, 1, 0, 0, 0, 318, 321, 1, 0, 0, 0, 319, 317,
		1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 31, 1, 0, 0, 0, 321, 319, 1, 0,
		0, 0, 322, 323, 5, 17, 0, 0, 323, 324, 3, 34, 17, 0, 324, 325, 5, 19, 0,
		0, 325, 33, 1, 0, 0, 0, 326, 327, 3, 36, 18, 0, 327, 35, 1, 0, 0, 0, 328,
		333, 3, 38, 19, 0, 329, 330, 5, 27, 0, 0, 330, 332, 3, 38, 19, 0, 331,
		329, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333, 334,
		1, 0, 0, 0, 334, 37, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 341, 3, 40,
		20, 0, 337, 338, 5, 28, 0, 0, 338, 340, 3, 40, 20, 0, 339, 337, 1, 0, 0,
		0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342,
		39, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344, 346, 5, 29, 0, 0, 345, 344,
		1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 348, 3, 42,
		21, 0, 348, 41, 1, 0, 0, 0, 349, 355, 5, 41, 0, 0, 350, 351, 5, 24, 0,
		0, 351, 352, 3, 34, 17, 0, 352, 353, 5, 25, 0, 0, 353, 355, 1, 0, 0, 0,
		354, 349, 1, 0, 0, 0, 354, 350, 1, 0, 0, 0, 355, 43, 1, 0, 0, 0, 356, 357,
		5, 30, 0, 0, 357, 373, 7, 3, 0, 0, 358, 360, 5, 30, 0, 0, 359, 361, 5,
		39, 0, 0, 360, 359, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 362, 1, 0, 0,
		0, 362, 373, 5, 41, 0, 0, 363, 365, 5, 30, 0, 0, 364, 366, 5, 39, 0, 0,
		365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 369, 1, 0, 0, 0, 367,
		368, 5, 41, 0, 0, 368, 370, 5, 33, 0, 0, 369, 367, 1, 0, 0, 0, 369, 370,
		1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 373, 7, 4, 0, 0, 372, 356, 1, 0,
		0, 0, 372, 358, 1, 0, 0, 0, 372, 363, 1, 0, 0, 0, 373, 45, 1, 0, 0, 0,
		60, 47, 61, 63, 70, 73, 78, 90, 92, 98, 101, 107, 113, 115, 121, 124, 129,
		135, 141, 146, 156, 158, 164, 169, 175, 180, 183, 188, 194, 199, 202, 207,
		211, 214, 219, 225, 230, 233, 238, 244, 247, 256, 262, 265, 270, 282, 286,
		290, 295, 298, 301, 309, 319, 333, 341, 345, 354, 360, 365, 369, 372,
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
	Note() antlr.TerminalNode
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(46)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(49)
		p.Match(fsmParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4605294559088) != 0 {
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(52)
				p.State()
			}

		case 2:
			{
				p.SetState(53)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(54)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(55)
				p.Choice()
			}

		case 5:
			{
				p.SetState(56)
				p.Junction()
			}

		case 6:
			{
				p.SetState(57)
				p.Fork()
			}

		case 7:
			{
				p.SetState(58)
				p.Join()
			}

		case 8:
			{
				p.SetState(59)
				p.Point()
			}

		case 9:
			{
				p.SetState(60)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(fsmParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
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
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(69)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(72)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(75)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(77)
			p.Stereotype()
		}

	}
	{
		p.SetState(80)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4605294559088) != 0 {
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(81)
				p.State()
			}

		case 2:
			{
				p.SetState(82)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(83)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(84)
				p.Choice()
			}

		case 5:
			{
				p.SetState(85)
				p.Junction()
			}

		case 6:
			{
				p.SetState(86)
				p.Fork()
			}

		case 7:
			{
				p.SetState(87)
				p.Join()
			}

		case 8:
			{
				p.SetState(88)
				p.Point()
			}

		case 9:
			{
				p.SetState(89)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(97)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(100)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(103)
		p.Match(fsmParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(fsmParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(106)
			p.Stereotype()
		}

	}
	{
		p.SetState(109)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4536575078528) != 0 {
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(110)
				p.Region()
			}

		case 2:
			{
				p.SetState(111)
				p.Point()
			}

		case 3:
			{
				p.SetState(112)
				p.Event()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(120)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserInitial {
		{
			p.SetState(123)
			p.Match(fsmParserInitial)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(126)
		p.Match(fsmParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(128)
			p.Stereotype()
		}

	}
	{
		p.SetState(131)
		p.Match(fsmParserT__1)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4536575078400) != 0 {
		{
			p.SetState(132)
			p.Event()
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
	Note() antlr.TerminalNode
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
	{
		p.SetState(143)
		p.Match(fsmParserT__6)
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

	if _la == fsmParserT__14 {
		{
			p.SetState(145)
			p.Stereotype()
		}

	}
	{
		p.SetState(148)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4466765991792) != 0 {
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(149)
				p.State()
			}

		case 2:
			{
				p.SetState(150)
				p.Parallel()
			}

		case 3:
			{
				p.SetState(151)
				p.Submachine()
			}

		case 4:
			{
				p.SetState(152)
				p.Choice()
			}

		case 5:
			{
				p.SetState(153)
				p.Junction()
			}

		case 6:
			{
				p.SetState(154)
				p.Fork()
			}

		case 7:
			{
				p.SetState(155)
				p.Join()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
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
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(163)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(166)
		p.Match(fsmParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(168)
			p.Stereotype()
		}

	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__1:
		{
			p.SetState(171)
			p.Match(fsmParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4399124578304) != 0 {
			{
				p.SetState(172)
				p.Branch()
			}

			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(178)
			p.Match(fsmParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__16, fsmParserT__21, fsmParserT__29, fsmParserNote:
		{
			p.SetState(179)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(182)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(185)
		p.Match(fsmParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(187)
			p.Stereotype()
		}

	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__1:
		{
			p.SetState(190)
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

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4399124578304) != 0 {
			{
				p.SetState(191)
				p.Branch()
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

	case fsmParserT__16, fsmParserT__21, fsmParserT__29, fsmParserNote:
		{
			p.SetState(198)
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
	p.EnterRule(localctx, 14, fsmParserRULE_fork)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(201)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(204)
		p.Match(fsmParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(206)
			p.Stereotype()
		}

	}
	{
		p.SetState(209)
		p.Match(fsmParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4399124447232) != 0 {
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserNote {
			{
				p.SetState(210)
				p.Match(fsmParserNote)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
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
	p.EnterRule(localctx, 16, fsmParserRULE_join)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(224)
			p.Match(fsmParserNote)
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

	if _la == fsmParserT__14 {
		{
			p.SetState(229)
			p.Stereotype()
		}

	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__21 {
		{
			p.SetState(232)
			p.Actions()
		}

	}
	{
		p.SetState(235)
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
	p.EnterRule(localctx, 18, fsmParserRULE_point)
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
	{
		p.SetState(240)

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
		p.SetState(241)
		p.Match(fsmParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__14 {
		{
			p.SetState(243)
			p.Stereotype()
		}

	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__21 {
		{
			p.SetState(246)
			p.Actions()
		}

	}
	{
		p.SetState(249)
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
		p.SetState(251)
		p.Match(fsmParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
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
	p.EnterRule(localctx, 22, fsmParserRULE_branch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(255)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(258)
			p.Match(fsmParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(fsmParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Match(fsmParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(261)
			p.Guard()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__21 {
		{
			p.SetState(264)
			p.Actions()
		}

	}
	{
		p.SetState(267)
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
	p.EnterRule(localctx, 24, fsmParserRULE_event)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserNote {
		{
			p.SetState(269)
			p.Match(fsmParserNote)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(272)

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
			p.SetState(273)
			p.Actions()
		}

	case 2:
		{
			p.SetState(274)
			p.Match(fsmParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)

			var _m = p.Match(fsmParserIdentifier)

			localctx.(*EventContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(fsmParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Match(fsmParserDefer)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(278)
			p.Match(fsmParserInvariant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Guard()
		}

	case 4:
		{
			p.SetState(280)
			p.Trigger()
		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__16 {
			{
				p.SetState(281)
				p.Guard()
			}

		}
		{
			p.SetState(284)
			p.Actions()
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(285)
				p.Goto_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		{
			p.SetState(288)
			p.Trigger()
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__16 {
			{
				p.SetState(289)
				p.Guard()
			}

		}
		{
			p.SetState(292)
			p.Goto_()
		}

	case 6:
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__16 {
			{
				p.SetState(294)
				p.Guard()
			}

		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserT__21 {
			{
				p.SetState(297)
				p.Actions()
			}

		}
		{
			p.SetState(300)
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

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.Match(fsmParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)

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
			p.SetState(305)
			p.Match(fsmParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(fsmParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)

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
			p.SetState(308)
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
		p.SetState(311)
		p.Match(fsmParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
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
		p.SetState(314)
		p.Match(fsmParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__25 {
		{
			p.SetState(315)
			p.Match(fsmParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Match(fsmParserIdentifier)
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
		p.SetState(322)
		p.Match(fsmParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Expression()
	}
	{
		p.SetState(324)
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
		p.SetState(326)
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
		p.SetState(328)
		p.And_expression()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__26 {
		{
			p.SetState(329)
			p.Match(fsmParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.And_expression()
		}

		p.SetState(335)
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
		p.SetState(336)
		p.Not_expression()
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fsmParserT__27 {
		{
			p.SetState(337)
			p.Match(fsmParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Not_expression()
		}

		p.SetState(343)
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
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fsmParserT__28 {
		{
			p.SetState(344)
			p.Match(fsmParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(347)
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
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fsmParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fsmParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(350)
			p.Match(fsmParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Expression()
		}
		{
			p.SetState(352)
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

	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.Match(fsmParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
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
			p.SetState(358)
			p.Match(fsmParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserLocal {
			{
				p.SetState(359)
				p.Match(fsmParserLocal)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(362)
			p.Match(fsmParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(363)
			p.Match(fsmParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserLocal {
			{
				p.SetState(364)
				p.Match(fsmParserLocal)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fsmParserIdentifier {
			{
				p.SetState(367)
				p.Match(fsmParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(368)
				p.Match(fsmParserT__32)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(371)
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
