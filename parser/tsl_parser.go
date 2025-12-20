// Code generated from TslParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TslParser
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

type TslParser struct {
	*antlr.BaseParser
}

var TslParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tslparserParserInit() {
	staticData := &TslParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'package'", "'const'", "'var'", "'func'", "'if'", "'else'", "'for'",
		"'true'", "'false'", "'return'", "'and'", "'or'", "'not'", "'null'",
		"'this'", "'new'", "'bar'", "'open'", "'high'", "'low'", "'close'",
		"'buy'", "'sell'", "'sellShort'", "'buyToCover'", "'at'", "'market'",
		"'stop'", "'limit'", "'contracts'", "'int'", "'real'", "'bool'", "'string'",
		"'time'", "'date'", "'timeseries'", "'enum'", "'class'", "'error'",
		"'list'", "'map'", "", "", "", "", "", "'('", "')'", "'['", "']'", "'{'",
		"'}'", "'='", "'*'", "'/'", "'+'", "'-'", "','", "'.'", "':'", "'''",
		"'<>'", "'<'", "'<='", "'>'", "'>='",
	}
	staticData.SymbolicNames = []string{
		"", "PACKAGE", "CONST", "VAR", "FUNC", "IF", "ELSE", "FOR", "TRUE",
		"FALSE", "RETURN", "AND", "OR", "NOT", "NULL", "THIS", "NEW", "BAR",
		"OPEN", "HIGH", "LOW", "CLOSE", "BUY", "SELL", "SELL_SHORT", "BUY_TO_COVER",
		"AT", "MARKET", "STOP", "LIMIT", "CONTRACTS", "INT", "REAL", "BOOL",
		"STRING", "TIME", "DATE", "TIMESERIES", "ENUM", "CLASS", "ERROR", "LIST",
		"MAP", "IDENTIFIER", "INT_VALUE", "REAL_VALUE", "DIGIT", "STRING_VALUE",
		"L_PAREN", "R_PAREN", "L_BRACKET", "R_BRACKET", "L_CURLY", "R_CURLY",
		"EQUAL", "STAR", "SLASH", "PLUS", "MINUS", "COMMA", "DOT", "COLON",
		"APOS", "NOT_EQUAL", "LESS_THAN", "LESS_OR_EQUAL", "GREATER_THAN", "GREATER_OR_EQUAL",
		"WS", "NEWLINE", "BLOCK_COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"script", "packageDef", "constantsDef", "constantDef", "variablesDef",
		"variableDef", "functionDef", "class", "parameters", "parameterDecl",
		"results", "enumDef", "enumItem", "enumValue", "classDef", "property",
		"type", "timeSeriesType", "listType", "mapType", "keyType", "block",
		"statement", "varsDeclaration", "varDeclaration", "varsAssignmentOrFunctionCall",
		"ifStatement", "elseIfBlock", "elseBlock", "returnStatement", "expression",
		"unaryExpression", "expressionInParenthesis", "chainedExpression", "chainItem",
		"paramsExpression", "accessorExpression", "barExpression", "listExpression",
		"mapExpression", "keyValueCouple", "keyValue", "constantValueExpression",
		"timeValue", "dateValue", "boolValue", "errorValue", "fqIdentifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 505, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 1, 0, 3, 0, 98, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 105, 8,
		0, 10, 0, 12, 0, 108, 9, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 116,
		8, 2, 10, 2, 12, 2, 119, 9, 2, 1, 2, 1, 2, 3, 2, 123, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 132, 8, 4, 10, 4, 12, 4, 135, 9, 4,
		1, 4, 1, 4, 3, 4, 139, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6,
		147, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 152, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 164, 8, 8, 10, 8, 12, 8, 167, 9,
		8, 3, 8, 169, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		5, 10, 179, 8, 10, 10, 10, 12, 10, 182, 9, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 5, 11, 188, 8, 11, 10, 11, 12, 11, 191, 9, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 200, 8, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 5, 14, 208, 8, 14, 10, 14, 12, 14, 211, 9, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 229, 8, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 3, 17, 236, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21,
		5, 21, 254, 8, 21, 10, 21, 12, 21, 257, 9, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 265, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23,
		271, 8, 23, 10, 23, 12, 23, 274, 9, 23, 1, 24, 1, 24, 1, 24, 3, 24, 279,
		8, 24, 1, 24, 1, 24, 3, 24, 283, 8, 24, 1, 25, 1, 25, 1, 25, 5, 25, 288,
		8, 25, 10, 25, 12, 25, 291, 9, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 297,
		8, 25, 10, 25, 12, 25, 300, 9, 25, 3, 25, 302, 8, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 5, 26, 308, 8, 26, 10, 26, 12, 26, 311, 9, 26, 1, 26, 3, 26,
		314, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 329, 8, 29, 10, 29, 12, 29, 332,
		9, 29, 3, 29, 334, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3,
		30, 342, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 4, 30, 356, 8, 30, 11, 30, 12, 30, 357, 1,
		30, 1, 30, 1, 30, 4, 30, 363, 8, 30, 11, 30, 12, 30, 364, 5, 30, 367, 8,
		30, 10, 30, 12, 30, 370, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		3, 31, 378, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33, 386,
		8, 33, 1, 33, 1, 33, 1, 33, 5, 33, 391, 8, 33, 10, 33, 12, 33, 394, 9,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 401, 8, 33, 10, 33, 12, 33,
		404, 9, 33, 3, 33, 406, 8, 33, 1, 34, 1, 34, 3, 34, 410, 8, 34, 1, 34,
		3, 34, 413, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 419, 8, 35, 10, 35,
		12, 35, 422, 9, 35, 3, 35, 424, 8, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 436, 8, 37, 1, 38, 1, 38, 1,
		38, 1, 38, 5, 38, 442, 8, 38, 10, 38, 12, 38, 445, 9, 38, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 453, 8, 39, 10, 39, 12, 39, 456, 9,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41,
		3, 41, 468, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3,
		42, 477, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 3, 47, 503, 8, 47, 1, 47, 0, 1,
		60, 48, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 0, 7, 2, 0, 44, 44,
		47, 47, 2, 0, 31, 31, 34, 36, 1, 0, 55, 56, 1, 0, 57, 58, 2, 0, 54, 54,
		63, 67, 1, 0, 18, 21, 1, 0, 8, 9, 534, 0, 97, 1, 0, 0, 0, 2, 109, 1, 0,
		0, 0, 4, 112, 1, 0, 0, 0, 6, 124, 1, 0, 0, 0, 8, 128, 1, 0, 0, 0, 10, 140,
		1, 0, 0, 0, 12, 144, 1, 0, 0, 0, 14, 155, 1, 0, 0, 0, 16, 159, 1, 0, 0,
		0, 18, 172, 1, 0, 0, 0, 20, 175, 1, 0, 0, 0, 22, 183, 1, 0, 0, 0, 24, 194,
		1, 0, 0, 0, 26, 201, 1, 0, 0, 0, 28, 203, 1, 0, 0, 0, 30, 214, 1, 0, 0,
		0, 32, 228, 1, 0, 0, 0, 34, 230, 1, 0, 0, 0, 36, 237, 1, 0, 0, 0, 38, 242,
		1, 0, 0, 0, 40, 249, 1, 0, 0, 0, 42, 251, 1, 0, 0, 0, 44, 264, 1, 0, 0,
		0, 46, 266, 1, 0, 0, 0, 48, 275, 1, 0, 0, 0, 50, 284, 1, 0, 0, 0, 52, 303,
		1, 0, 0, 0, 54, 315, 1, 0, 0, 0, 56, 320, 1, 0, 0, 0, 58, 333, 1, 0, 0,
		0, 60, 341, 1, 0, 0, 0, 62, 377, 1, 0, 0, 0, 64, 379, 1, 0, 0, 0, 66, 405,
		1, 0, 0, 0, 68, 407, 1, 0, 0, 0, 70, 414, 1, 0, 0, 0, 72, 427, 1, 0, 0,
		0, 74, 431, 1, 0, 0, 0, 76, 437, 1, 0, 0, 0, 78, 448, 1, 0, 0, 0, 80, 459,
		1, 0, 0, 0, 82, 467, 1, 0, 0, 0, 84, 476, 1, 0, 0, 0, 86, 478, 1, 0, 0,
		0, 88, 484, 1, 0, 0, 0, 90, 492, 1, 0, 0, 0, 92, 494, 1, 0, 0, 0, 94, 499,
		1, 0, 0, 0, 96, 98, 3, 2, 1, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0,
		98, 106, 1, 0, 0, 0, 99, 105, 3, 4, 2, 0, 100, 105, 3, 8, 4, 0, 101, 105,
		3, 12, 6, 0, 102, 105, 3, 22, 11, 0, 103, 105, 3, 28, 14, 0, 104, 99, 1,
		0, 0, 0, 104, 100, 1, 0, 0, 0, 104, 101, 1, 0, 0, 0, 104, 102, 1, 0, 0,
		0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106,
		107, 1, 0, 0, 0, 107, 1, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 110, 5,
		1, 0, 0, 110, 111, 5, 43, 0, 0, 111, 3, 1, 0, 0, 0, 112, 122, 5, 2, 0,
		0, 113, 117, 5, 52, 0, 0, 114, 116, 3, 6, 3, 0, 115, 114, 1, 0, 0, 0, 116,
		119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120,
		1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 123, 5, 53, 0, 0, 121, 123, 3, 6,
		3, 0, 122, 113, 1, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 5, 1, 0, 0, 0, 124,
		125, 5, 43, 0, 0, 125, 126, 5, 54, 0, 0, 126, 127, 3, 60, 30, 0, 127, 7,
		1, 0, 0, 0, 128, 138, 5, 3, 0, 0, 129, 133, 5, 52, 0, 0, 130, 132, 3, 10,
		5, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0,
		133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136,
		139, 5, 53, 0, 0, 137, 139, 3, 10, 5, 0, 138, 129, 1, 0, 0, 0, 138, 137,
		1, 0, 0, 0, 139, 9, 1, 0, 0, 0, 140, 141, 5, 43, 0, 0, 141, 142, 5, 54,
		0, 0, 142, 143, 3, 60, 30, 0, 143, 11, 1, 0, 0, 0, 144, 146, 5, 4, 0, 0,
		145, 147, 3, 14, 7, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147,
		148, 1, 0, 0, 0, 148, 149, 5, 43, 0, 0, 149, 151, 3, 16, 8, 0, 150, 152,
		3, 20, 10, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1,
		0, 0, 0, 153, 154, 3, 42, 21, 0, 154, 13, 1, 0, 0, 0, 155, 156, 5, 48,
		0, 0, 156, 157, 3, 94, 47, 0, 157, 158, 5, 49, 0, 0, 158, 15, 1, 0, 0,
		0, 159, 168, 5, 48, 0, 0, 160, 165, 3, 18, 9, 0, 161, 162, 5, 59, 0, 0,
		162, 164, 3, 18, 9, 0, 163, 161, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165,
		163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165,
		1, 0, 0, 0, 168, 160, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0,
		0, 0, 170, 171, 5, 49, 0, 0, 171, 17, 1, 0, 0, 0, 172, 173, 5, 43, 0, 0,
		173, 174, 3, 32, 16, 0, 174, 19, 1, 0, 0, 0, 175, 180, 3, 32, 16, 0, 176,
		177, 5, 59, 0, 0, 177, 179, 3, 32, 16, 0, 178, 176, 1, 0, 0, 0, 179, 182,
		1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 21, 1, 0,
		0, 0, 182, 180, 1, 0, 0, 0, 183, 184, 5, 38, 0, 0, 184, 185, 5, 43, 0,
		0, 185, 189, 5, 52, 0, 0, 186, 188, 3, 24, 12, 0, 187, 186, 1, 0, 0, 0,
		188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		192, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 193, 5, 53, 0, 0, 193, 23,
		1, 0, 0, 0, 194, 199, 5, 43, 0, 0, 195, 196, 5, 48, 0, 0, 196, 197, 3,
		26, 13, 0, 197, 198, 5, 49, 0, 0, 198, 200, 1, 0, 0, 0, 199, 195, 1, 0,
		0, 0, 199, 200, 1, 0, 0, 0, 200, 25, 1, 0, 0, 0, 201, 202, 7, 0, 0, 0,
		202, 27, 1, 0, 0, 0, 203, 204, 5, 39, 0, 0, 204, 205, 5, 43, 0, 0, 205,
		209, 5, 52, 0, 0, 206, 208, 3, 30, 15, 0, 207, 206, 1, 0, 0, 0, 208, 211,
		1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212, 1, 0,
		0, 0, 211, 209, 1, 0, 0, 0, 212, 213, 5, 53, 0, 0, 213, 29, 1, 0, 0, 0,
		214, 215, 5, 43, 0, 0, 215, 216, 3, 32, 16, 0, 216, 31, 1, 0, 0, 0, 217,
		229, 5, 31, 0, 0, 218, 229, 5, 32, 0, 0, 219, 229, 5, 33, 0, 0, 220, 229,
		5, 34, 0, 0, 221, 229, 5, 35, 0, 0, 222, 229, 5, 36, 0, 0, 223, 229, 3,
		34, 17, 0, 224, 229, 3, 94, 47, 0, 225, 229, 5, 40, 0, 0, 226, 229, 3,
		36, 18, 0, 227, 229, 3, 38, 19, 0, 228, 217, 1, 0, 0, 0, 228, 218, 1, 0,
		0, 0, 228, 219, 1, 0, 0, 0, 228, 220, 1, 0, 0, 0, 228, 221, 1, 0, 0, 0,
		228, 222, 1, 0, 0, 0, 228, 223, 1, 0, 0, 0, 228, 224, 1, 0, 0, 0, 228,
		225, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229, 33, 1,
		0, 0, 0, 230, 235, 5, 37, 0, 0, 231, 232, 5, 48, 0, 0, 232, 233, 3, 32,
		16, 0, 233, 234, 5, 49, 0, 0, 234, 236, 1, 0, 0, 0, 235, 231, 1, 0, 0,
		0, 235, 236, 1, 0, 0, 0, 236, 35, 1, 0, 0, 0, 237, 238, 5, 41, 0, 0, 238,
		239, 5, 48, 0, 0, 239, 240, 3, 32, 16, 0, 240, 241, 5, 49, 0, 0, 241, 37,
		1, 0, 0, 0, 242, 243, 5, 42, 0, 0, 243, 244, 5, 48, 0, 0, 244, 245, 3,
		40, 20, 0, 245, 246, 5, 59, 0, 0, 246, 247, 3, 32, 16, 0, 247, 248, 5,
		49, 0, 0, 248, 39, 1, 0, 0, 0, 249, 250, 7, 1, 0, 0, 250, 41, 1, 0, 0,
		0, 251, 255, 5, 52, 0, 0, 252, 254, 3, 44, 22, 0, 253, 252, 1, 0, 0, 0,
		254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		258, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 259, 5, 53, 0, 0, 259, 43,
		1, 0, 0, 0, 260, 265, 3, 46, 23, 0, 261, 265, 3, 50, 25, 0, 262, 265, 3,
		52, 26, 0, 263, 265, 3, 58, 29, 0, 264, 260, 1, 0, 0, 0, 264, 261, 1, 0,
		0, 0, 264, 262, 1, 0, 0, 0, 264, 263, 1, 0, 0, 0, 265, 45, 1, 0, 0, 0,
		266, 267, 5, 3, 0, 0, 267, 272, 3, 48, 24, 0, 268, 269, 5, 59, 0, 0, 269,
		271, 3, 48, 24, 0, 270, 268, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270,
		1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 47, 1, 0, 0, 0, 274, 272, 1, 0,
		0, 0, 275, 278, 5, 43, 0, 0, 276, 277, 5, 61, 0, 0, 277, 279, 3, 32, 16,
		0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 282, 1, 0, 0, 0, 280,
		281, 5, 54, 0, 0, 281, 283, 3, 60, 30, 0, 282, 280, 1, 0, 0, 0, 282, 283,
		1, 0, 0, 0, 283, 49, 1, 0, 0, 0, 284, 289, 3, 66, 33, 0, 285, 286, 5, 59,
		0, 0, 286, 288, 3, 66, 33, 0, 287, 285, 1, 0, 0, 0, 288, 291, 1, 0, 0,
		0, 289, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 301, 1, 0, 0, 0, 291,
		289, 1, 0, 0, 0, 292, 293, 5, 54, 0, 0, 293, 298, 3, 60, 30, 0, 294, 295,
		5, 59, 0, 0, 295, 297, 3, 60, 30, 0, 296, 294, 1, 0, 0, 0, 297, 300, 1,
		0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 302, 1, 0, 0,
		0, 300, 298, 1, 0, 0, 0, 301, 292, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302,
		51, 1, 0, 0, 0, 303, 304, 5, 5, 0, 0, 304, 305, 3, 60, 30, 0, 305, 309,
		3, 42, 21, 0, 306, 308, 3, 54, 27, 0, 307, 306, 1, 0, 0, 0, 308, 311, 1,
		0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 313, 1, 0, 0,
		0, 311, 309, 1, 0, 0, 0, 312, 314, 3, 56, 28, 0, 313, 312, 1, 0, 0, 0,
		313, 314, 1, 0, 0, 0, 314, 53, 1, 0, 0, 0, 315, 316, 5, 6, 0, 0, 316, 317,
		5, 5, 0, 0, 317, 318, 3, 60, 30, 0, 318, 319, 3, 42, 21, 0, 319, 55, 1,
		0, 0, 0, 320, 321, 5, 6, 0, 0, 321, 322, 3, 42, 21, 0, 322, 57, 1, 0, 0,
		0, 323, 334, 5, 10, 0, 0, 324, 325, 5, 10, 0, 0, 325, 330, 3, 60, 30, 0,
		326, 327, 5, 59, 0, 0, 327, 329, 3, 60, 30, 0, 328, 326, 1, 0, 0, 0, 329,
		332, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 334,
		1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 333, 323, 1, 0, 0, 0, 333, 324, 1, 0,
		0, 0, 334, 59, 1, 0, 0, 0, 335, 336, 6, 30, -1, 0, 336, 342, 3, 62, 31,
		0, 337, 338, 5, 13, 0, 0, 338, 342, 3, 60, 30, 5, 339, 342, 3, 76, 38,
		0, 340, 342, 3, 78, 39, 0, 341, 335, 1, 0, 0, 0, 341, 337, 1, 0, 0, 0,
		341, 339, 1, 0, 0, 0, 341, 340, 1, 0, 0, 0, 342, 368, 1, 0, 0, 0, 343,
		344, 10, 8, 0, 0, 344, 345, 7, 2, 0, 0, 345, 367, 3, 60, 30, 9, 346, 347,
		10, 7, 0, 0, 347, 348, 7, 3, 0, 0, 348, 367, 3, 60, 30, 8, 349, 350, 10,
		6, 0, 0, 350, 351, 7, 4, 0, 0, 351, 367, 3, 60, 30, 7, 352, 355, 10, 4,
		0, 0, 353, 354, 5, 11, 0, 0, 354, 356, 3, 60, 30, 0, 355, 353, 1, 0, 0,
		0, 356, 357, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358,
		367, 1, 0, 0, 0, 359, 362, 10, 3, 0, 0, 360, 361, 5, 12, 0, 0, 361, 363,
		3, 60, 30, 0, 362, 360, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 362, 1,
		0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 367, 1, 0, 0, 0, 366, 343, 1, 0, 0,
		0, 366, 346, 1, 0, 0, 0, 366, 349, 1, 0, 0, 0, 366, 352, 1, 0, 0, 0, 366,
		359, 1, 0, 0, 0, 367, 370, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 369,
		1, 0, 0, 0, 369, 61, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 371, 372, 7, 3,
		0, 0, 372, 378, 3, 62, 31, 0, 373, 378, 3, 64, 32, 0, 374, 378, 3, 74,
		37, 0, 375, 378, 3, 84, 42, 0, 376, 378, 3, 66, 33, 0, 377, 371, 1, 0,
		0, 0, 377, 373, 1, 0, 0, 0, 377, 374, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0,
		377, 376, 1, 0, 0, 0, 378, 63, 1, 0, 0, 0, 379, 380, 5, 48, 0, 0, 380,
		381, 3, 60, 30, 0, 381, 382, 5, 49, 0, 0, 382, 65, 1, 0, 0, 0, 383, 384,
		5, 15, 0, 0, 384, 386, 5, 60, 0, 0, 385, 383, 1, 0, 0, 0, 385, 386, 1,
		0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 392, 3, 68, 34, 0, 388, 389, 5, 60,
		0, 0, 389, 391, 3, 68, 34, 0, 390, 388, 1, 0, 0, 0, 391, 394, 1, 0, 0,
		0, 392, 390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 406, 1, 0, 0, 0, 394,
		392, 1, 0, 0, 0, 395, 396, 5, 16, 0, 0, 396, 397, 3, 94, 47, 0, 397, 402,
		3, 70, 35, 0, 398, 399, 5, 60, 0, 0, 399, 401, 3, 68, 34, 0, 400, 398,
		1, 0, 0, 0, 401, 404, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 402, 403, 1, 0,
		0, 0, 403, 406, 1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 405, 385, 1, 0, 0, 0,
		405, 395, 1, 0, 0, 0, 406, 67, 1, 0, 0, 0, 407, 409, 5, 43, 0, 0, 408,
		410, 3, 70, 35, 0, 409, 408, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 412,
		1, 0, 0, 0, 411, 413, 3, 72, 36, 0, 412, 411, 1, 0, 0, 0, 412, 413, 1,
		0, 0, 0, 413, 69, 1, 0, 0, 0, 414, 423, 5, 48, 0, 0, 415, 420, 3, 60, 30,
		0, 416, 417, 5, 59, 0, 0, 417, 419, 3, 60, 30, 0, 418, 416, 1, 0, 0, 0,
		419, 422, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0, 421,
		424, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 423, 415, 1, 0, 0, 0, 423, 424,
		1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 426, 5, 49, 0, 0, 426, 71, 1, 0,
		0, 0, 427, 428, 5, 50, 0, 0, 428, 429, 3, 60, 30, 0, 429, 430, 5, 51, 0,
		0, 430, 73, 1, 0, 0, 0, 431, 432, 5, 17, 0, 0, 432, 433, 5, 60, 0, 0, 433,
		435, 7, 5, 0, 0, 434, 436, 3, 72, 36, 0, 435, 434, 1, 0, 0, 0, 435, 436,
		1, 0, 0, 0, 436, 75, 1, 0, 0, 0, 437, 438, 5, 50, 0, 0, 438, 443, 3, 60,
		30, 0, 439, 440, 5, 59, 0, 0, 440, 442, 3, 60, 30, 0, 441, 439, 1, 0, 0,
		0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444,
		446, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 447, 5, 51, 0, 0, 447, 77,
		1, 0, 0, 0, 448, 449, 5, 52, 0, 0, 449, 454, 3, 80, 40, 0, 450, 451, 5,
		59, 0, 0, 451, 453, 3, 80, 40, 0, 452, 450, 1, 0, 0, 0, 453, 456, 1, 0,
		0, 0, 454, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 457, 1, 0, 0, 0,
		456, 454, 1, 0, 0, 0, 457, 458, 5, 53, 0, 0, 458, 79, 1, 0, 0, 0, 459,
		460, 3, 82, 41, 0, 460, 461, 5, 61, 0, 0, 461, 462, 3, 60, 30, 0, 462,
		81, 1, 0, 0, 0, 463, 468, 5, 44, 0, 0, 464, 468, 5, 47, 0, 0, 465, 468,
		3, 86, 43, 0, 466, 468, 3, 88, 44, 0, 467, 463, 1, 0, 0, 0, 467, 464, 1,
		0, 0, 0, 467, 465, 1, 0, 0, 0, 467, 466, 1, 0, 0, 0, 468, 83, 1, 0, 0,
		0, 469, 477, 5, 44, 0, 0, 470, 477, 5, 45, 0, 0, 471, 477, 3, 90, 45, 0,
		472, 477, 5, 47, 0, 0, 473, 477, 3, 86, 43, 0, 474, 477, 3, 88, 44, 0,
		475, 477, 3, 92, 46, 0, 476, 469, 1, 0, 0, 0, 476, 470, 1, 0, 0, 0, 476,
		471, 1, 0, 0, 0, 476, 472, 1, 0, 0, 0, 476, 473, 1, 0, 0, 0, 476, 474,
		1, 0, 0, 0, 476, 475, 1, 0, 0, 0, 477, 85, 1, 0, 0, 0, 478, 479, 5, 62,
		0, 0, 479, 480, 5, 44, 0, 0, 480, 481, 5, 61, 0, 0, 481, 482, 5, 44, 0,
		0, 482, 483, 5, 62, 0, 0, 483, 87, 1, 0, 0, 0, 484, 485, 5, 62, 0, 0, 485,
		486, 5, 44, 0, 0, 486, 487, 5, 58, 0, 0, 487, 488, 5, 44, 0, 0, 488, 489,
		5, 58, 0, 0, 489, 490, 5, 44, 0, 0, 490, 491, 5, 62, 0, 0, 491, 89, 1,
		0, 0, 0, 492, 493, 7, 6, 0, 0, 493, 91, 1, 0, 0, 0, 494, 495, 5, 40, 0,
		0, 495, 496, 5, 48, 0, 0, 496, 497, 5, 47, 0, 0, 497, 498, 5, 49, 0, 0,
		498, 93, 1, 0, 0, 0, 499, 502, 5, 43, 0, 0, 500, 501, 5, 60, 0, 0, 501,
		503, 5, 43, 0, 0, 502, 500, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 95,
		1, 0, 0, 0, 49, 97, 104, 106, 117, 122, 133, 138, 146, 151, 165, 168, 180,
		189, 199, 209, 228, 235, 255, 264, 272, 278, 282, 289, 298, 301, 309, 313,
		330, 333, 341, 357, 364, 366, 368, 377, 385, 392, 402, 405, 409, 412, 420,
		423, 435, 443, 454, 467, 476, 502,
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

// TslParserInit initializes any static state used to implement TslParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTslParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TslParserInit() {
	staticData := &TslParserParserStaticData
	staticData.once.Do(tslparserParserInit)
}

// NewTslParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTslParser(input antlr.TokenStream) *TslParser {
	TslParserInit()
	this := new(TslParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TslParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TslParser.g4"

	return this
}

// TslParser tokens.
const (
	TslParserEOF              = antlr.TokenEOF
	TslParserPACKAGE          = 1
	TslParserCONST            = 2
	TslParserVAR              = 3
	TslParserFUNC             = 4
	TslParserIF               = 5
	TslParserELSE             = 6
	TslParserFOR              = 7
	TslParserTRUE             = 8
	TslParserFALSE            = 9
	TslParserRETURN           = 10
	TslParserAND              = 11
	TslParserOR               = 12
	TslParserNOT              = 13
	TslParserNULL             = 14
	TslParserTHIS             = 15
	TslParserNEW              = 16
	TslParserBAR              = 17
	TslParserOPEN             = 18
	TslParserHIGH             = 19
	TslParserLOW              = 20
	TslParserCLOSE            = 21
	TslParserBUY              = 22
	TslParserSELL             = 23
	TslParserSELL_SHORT       = 24
	TslParserBUY_TO_COVER     = 25
	TslParserAT               = 26
	TslParserMARKET           = 27
	TslParserSTOP             = 28
	TslParserLIMIT            = 29
	TslParserCONTRACTS        = 30
	TslParserINT              = 31
	TslParserREAL             = 32
	TslParserBOOL             = 33
	TslParserSTRING           = 34
	TslParserTIME             = 35
	TslParserDATE             = 36
	TslParserTIMESERIES       = 37
	TslParserENUM             = 38
	TslParserCLASS            = 39
	TslParserERROR            = 40
	TslParserLIST             = 41
	TslParserMAP              = 42
	TslParserIDENTIFIER       = 43
	TslParserINT_VALUE        = 44
	TslParserREAL_VALUE       = 45
	TslParserDIGIT            = 46
	TslParserSTRING_VALUE     = 47
	TslParserL_PAREN          = 48
	TslParserR_PAREN          = 49
	TslParserL_BRACKET        = 50
	TslParserR_BRACKET        = 51
	TslParserL_CURLY          = 52
	TslParserR_CURLY          = 53
	TslParserEQUAL            = 54
	TslParserSTAR             = 55
	TslParserSLASH            = 56
	TslParserPLUS             = 57
	TslParserMINUS            = 58
	TslParserCOMMA            = 59
	TslParserDOT              = 60
	TslParserCOLON            = 61
	TslParserAPOS             = 62
	TslParserNOT_EQUAL        = 63
	TslParserLESS_THAN        = 64
	TslParserLESS_OR_EQUAL    = 65
	TslParserGREATER_THAN     = 66
	TslParserGREATER_OR_EQUAL = 67
	TslParserWS               = 68
	TslParserNEWLINE          = 69
	TslParserBLOCK_COMMENT    = 70
	TslParserLINE_COMMENT     = 71
)

// TslParser rules.
const (
	TslParserRULE_script                       = 0
	TslParserRULE_packageDef                   = 1
	TslParserRULE_constantsDef                 = 2
	TslParserRULE_constantDef                  = 3
	TslParserRULE_variablesDef                 = 4
	TslParserRULE_variableDef                  = 5
	TslParserRULE_functionDef                  = 6
	TslParserRULE_class                        = 7
	TslParserRULE_parameters                   = 8
	TslParserRULE_parameterDecl                = 9
	TslParserRULE_results                      = 10
	TslParserRULE_enumDef                      = 11
	TslParserRULE_enumItem                     = 12
	TslParserRULE_enumValue                    = 13
	TslParserRULE_classDef                     = 14
	TslParserRULE_property                     = 15
	TslParserRULE_type                         = 16
	TslParserRULE_timeSeriesType               = 17
	TslParserRULE_listType                     = 18
	TslParserRULE_mapType                      = 19
	TslParserRULE_keyType                      = 20
	TslParserRULE_block                        = 21
	TslParserRULE_statement                    = 22
	TslParserRULE_varsDeclaration              = 23
	TslParserRULE_varDeclaration               = 24
	TslParserRULE_varsAssignmentOrFunctionCall = 25
	TslParserRULE_ifStatement                  = 26
	TslParserRULE_elseIfBlock                  = 27
	TslParserRULE_elseBlock                    = 28
	TslParserRULE_returnStatement              = 29
	TslParserRULE_expression                   = 30
	TslParserRULE_unaryExpression              = 31
	TslParserRULE_expressionInParenthesis      = 32
	TslParserRULE_chainedExpression            = 33
	TslParserRULE_chainItem                    = 34
	TslParserRULE_paramsExpression             = 35
	TslParserRULE_accessorExpression           = 36
	TslParserRULE_barExpression                = 37
	TslParserRULE_listExpression               = 38
	TslParserRULE_mapExpression                = 39
	TslParserRULE_keyValueCouple               = 40
	TslParserRULE_keyValue                     = 41
	TslParserRULE_constantValueExpression      = 42
	TslParserRULE_timeValue                    = 43
	TslParserRULE_dateValue                    = 44
	TslParserRULE_boolValue                    = 45
	TslParserRULE_errorValue                   = 46
	TslParserRULE_fqIdentifier                 = 47
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PackageDef() IPackageDefContext
	AllConstantsDef() []IConstantsDefContext
	ConstantsDef(i int) IConstantsDefContext
	AllVariablesDef() []IVariablesDefContext
	VariablesDef(i int) IVariablesDefContext
	AllFunctionDef() []IFunctionDefContext
	FunctionDef(i int) IFunctionDefContext
	AllEnumDef() []IEnumDefContext
	EnumDef(i int) IEnumDefContext
	AllClassDef() []IClassDefContext
	ClassDef(i int) IClassDefContext

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_script
	return p
}

func InitEmptyScriptContext(p *ScriptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_script
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) PackageDef() IPackageDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageDefContext)
}

func (s *ScriptContext) AllConstantsDef() []IConstantsDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantsDefContext); ok {
			len++
		}
	}

	tst := make([]IConstantsDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantsDefContext); ok {
			tst[i] = t.(IConstantsDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) ConstantsDef(i int) IConstantsDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantsDefContext); ok {
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

	return t.(IConstantsDefContext)
}

func (s *ScriptContext) AllVariablesDef() []IVariablesDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariablesDefContext); ok {
			len++
		}
	}

	tst := make([]IVariablesDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariablesDefContext); ok {
			tst[i] = t.(IVariablesDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) VariablesDef(i int) IVariablesDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariablesDefContext); ok {
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

	return t.(IVariablesDefContext)
}

func (s *ScriptContext) AllFunctionDef() []IFunctionDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDefContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDefContext); ok {
			tst[i] = t.(IFunctionDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) FunctionDef(i int) IFunctionDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
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

	return t.(IFunctionDefContext)
}

func (s *ScriptContext) AllEnumDef() []IEnumDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumDefContext); ok {
			len++
		}
	}

	tst := make([]IEnumDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumDefContext); ok {
			tst[i] = t.(IEnumDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) EnumDef(i int) IEnumDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
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

	return t.(IEnumDefContext)
}

func (s *ScriptContext) AllClassDef() []IClassDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClassDefContext); ok {
			len++
		}
	}

	tst := make([]IClassDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClassDefContext); ok {
			tst[i] = t.(IClassDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) ClassDef(i int) IClassDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDefContext); ok {
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

	return t.(IClassDefContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *TslParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TslParserRULE_script)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserPACKAGE {
		{
			p.SetState(96)
			p.PackageDef()
		}

	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&824633720860) != 0 {
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TslParserCONST:
			{
				p.SetState(99)
				p.ConstantsDef()
			}

		case TslParserVAR:
			{
				p.SetState(100)
				p.VariablesDef()
			}

		case TslParserFUNC:
			{
				p.SetState(101)
				p.FunctionDef()
			}

		case TslParserENUM:
			{
				p.SetState(102)
				p.EnumDef()
			}

		case TslParserCLASS:
			{
				p.SetState(103)
				p.ClassDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(108)
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

// IPackageDefContext is an interface to support dynamic dispatch.
type IPackageDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsPackageDefContext differentiates from other interfaces.
	IsPackageDefContext()
}

type PackageDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageDefContext() *PackageDefContext {
	var p = new(PackageDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_packageDef
	return p
}

func InitEmptyPackageDefContext(p *PackageDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_packageDef
}

func (*PackageDefContext) IsPackageDefContext() {}

func NewPackageDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageDefContext {
	var p = new(PackageDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_packageDef

	return p
}

func (s *PackageDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageDefContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(TslParserPACKAGE, 0)
}

func (s *PackageDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *PackageDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterPackageDef(s)
	}
}

func (s *PackageDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitPackageDef(s)
	}
}

func (p *TslParser) PackageDef() (localctx IPackageDefContext) {
	localctx = NewPackageDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TslParserRULE_packageDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(TslParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(TslParserIDENTIFIER)
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

// IConstantsDefContext is an interface to support dynamic dispatch.
type IConstantsDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllConstantDef() []IConstantDefContext
	ConstantDef(i int) IConstantDefContext

	// IsConstantsDefContext differentiates from other interfaces.
	IsConstantsDefContext()
}

type ConstantsDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantsDefContext() *ConstantsDefContext {
	var p = new(ConstantsDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantsDef
	return p
}

func InitEmptyConstantsDefContext(p *ConstantsDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantsDef
}

func (*ConstantsDefContext) IsConstantsDefContext() {}

func NewConstantsDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantsDefContext {
	var p = new(ConstantsDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_constantsDef

	return p
}

func (s *ConstantsDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantsDefContext) CONST() antlr.TerminalNode {
	return s.GetToken(TslParserCONST, 0)
}

func (s *ConstantsDefContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *ConstantsDefContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *ConstantsDefContext) AllConstantDef() []IConstantDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantDefContext); ok {
			len++
		}
	}

	tst := make([]IConstantDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantDefContext); ok {
			tst[i] = t.(IConstantDefContext)
			i++
		}
	}

	return tst
}

func (s *ConstantsDefContext) ConstantDef(i int) IConstantDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantDefContext); ok {
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

	return t.(IConstantDefContext)
}

func (s *ConstantsDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantsDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantsDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterConstantsDef(s)
	}
}

func (s *ConstantsDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitConstantsDef(s)
	}
}

func (p *TslParser) ConstantsDef() (localctx IConstantsDefContext) {
	localctx = NewConstantsDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TslParserRULE_constantsDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(TslParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TslParserL_CURLY:
		{
			p.SetState(113)
			p.Match(TslParserL_CURLY)
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

		for _la == TslParserIDENTIFIER {
			{
				p.SetState(114)
				p.ConstantDef()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(120)
			p.Match(TslParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserIDENTIFIER:
		{
			p.SetState(121)
			p.ConstantDef()
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

// IConstantDefContext is an interface to support dynamic dispatch.
type IConstantDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext

	// IsConstantDefContext differentiates from other interfaces.
	IsConstantDefContext()
}

type ConstantDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantDefContext() *ConstantDefContext {
	var p = new(ConstantDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantDef
	return p
}

func InitEmptyConstantDefContext(p *ConstantDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantDef
}

func (*ConstantDefContext) IsConstantDefContext() {}

func NewConstantDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDefContext {
	var p = new(ConstantDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_constantDef

	return p
}

func (s *ConstantDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *ConstantDefContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserEQUAL, 0)
}

func (s *ConstantDefContext) Expression() IExpressionContext {
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

func (s *ConstantDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterConstantDef(s)
	}
}

func (s *ConstantDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitConstantDef(s)
	}
}

func (p *TslParser) ConstantDef() (localctx IConstantDefContext) {
	localctx = NewConstantDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TslParserRULE_constantDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(TslParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.expression(0)
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

// IVariablesDefContext is an interface to support dynamic dispatch.
type IVariablesDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllVariableDef() []IVariableDefContext
	VariableDef(i int) IVariableDefContext

	// IsVariablesDefContext differentiates from other interfaces.
	IsVariablesDefContext()
}

type VariablesDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablesDefContext() *VariablesDefContext {
	var p = new(VariablesDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_variablesDef
	return p
}

func InitEmptyVariablesDefContext(p *VariablesDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_variablesDef
}

func (*VariablesDefContext) IsVariablesDefContext() {}

func NewVariablesDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesDefContext {
	var p = new(VariablesDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_variablesDef

	return p
}

func (s *VariablesDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesDefContext) VAR() antlr.TerminalNode {
	return s.GetToken(TslParserVAR, 0)
}

func (s *VariablesDefContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *VariablesDefContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *VariablesDefContext) AllVariableDef() []IVariableDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDefContext); ok {
			len++
		}
	}

	tst := make([]IVariableDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDefContext); ok {
			tst[i] = t.(IVariableDefContext)
			i++
		}
	}

	return tst
}

func (s *VariablesDefContext) VariableDef(i int) IVariableDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefContext); ok {
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

	return t.(IVariableDefContext)
}

func (s *VariablesDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablesDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterVariablesDef(s)
	}
}

func (s *VariablesDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitVariablesDef(s)
	}
}

func (p *TslParser) VariablesDef() (localctx IVariablesDefContext) {
	localctx = NewVariablesDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TslParserRULE_variablesDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(TslParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TslParserL_CURLY:
		{
			p.SetState(129)
			p.Match(TslParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TslParserIDENTIFIER {
			{
				p.SetState(130)
				p.VariableDef()
			}

			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(136)
			p.Match(TslParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserIDENTIFIER:
		{
			p.SetState(137)
			p.VariableDef()
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

// IVariableDefContext is an interface to support dynamic dispatch.
type IVariableDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVariableDefContext differentiates from other interfaces.
	IsVariableDefContext()
}

type VariableDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefContext() *VariableDefContext {
	var p = new(VariableDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_variableDef
	return p
}

func InitEmptyVariableDefContext(p *VariableDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_variableDef
}

func (*VariableDefContext) IsVariableDefContext() {}

func NewVariableDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefContext {
	var p = new(VariableDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_variableDef

	return p
}

func (s *VariableDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *VariableDefContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserEQUAL, 0)
}

func (s *VariableDefContext) Expression() IExpressionContext {
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

func (s *VariableDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterVariableDef(s)
	}
}

func (s *VariableDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitVariableDef(s)
	}
}

func (p *TslParser) VariableDef() (localctx IVariableDefContext) {
	localctx = NewVariableDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TslParserRULE_variableDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(TslParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.expression(0)
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

// IFunctionDefContext is an interface to support dynamic dispatch.
type IFunctionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Parameters() IParametersContext
	Block() IBlockContext
	Class() IClassContext
	Results() IResultsContext

	// IsFunctionDefContext differentiates from other interfaces.
	IsFunctionDefContext()
}

type FunctionDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefContext() *FunctionDefContext {
	var p = new(FunctionDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_functionDef
	return p
}

func InitEmptyFunctionDefContext(p *FunctionDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_functionDef
}

func (*FunctionDefContext) IsFunctionDefContext() {}

func NewFunctionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefContext {
	var p = new(FunctionDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_functionDef

	return p
}

func (s *FunctionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(TslParserFUNC, 0)
}

func (s *FunctionDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *FunctionDefContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefContext) Class() IClassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassContext)
}

func (s *FunctionDefContext) Results() IResultsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultsContext)
}

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterFunctionDef(s)
	}
}

func (s *FunctionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitFunctionDef(s)
	}
}

func (p *TslParser) FunctionDef() (localctx IFunctionDefContext) {
	localctx = NewFunctionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TslParserRULE_functionDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(TslParserFUNC)
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

	if _la == TslParserL_PAREN {
		{
			p.SetState(145)
			p.Class()
		}

	}
	{
		p.SetState(148)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Parameters()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16765404839936) != 0 {
		{
			p.SetState(150)
			p.Results()
		}

	}
	{
		p.SetState(153)
		p.Block()
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

// IClassContext is an interface to support dynamic dispatch.
type IClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	FqIdentifier() IFqIdentifierContext
	R_PAREN() antlr.TerminalNode

	// IsClassContext differentiates from other interfaces.
	IsClassContext()
}

type ClassContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassContext() *ClassContext {
	var p = new(ClassContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_class
	return p
}

func InitEmptyClassContext(p *ClassContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_class
}

func (*ClassContext) IsClassContext() {}

func NewClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassContext {
	var p = new(ClassContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_class

	return p
}

func (s *ClassContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ClassContext) FqIdentifier() IFqIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFqIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFqIdentifierContext)
}

func (s *ClassContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterClass(s)
	}
}

func (s *ClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitClass(s)
	}
}

func (p *TslParser) Class() (localctx IClassContext) {
	localctx = NewClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TslParserRULE_class)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.FqIdentifier()
	}
	{
		p.SetState(157)
		p.Match(TslParserR_PAREN)
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

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllParameterDecl() []IParameterDeclContext
	ParameterDecl(i int) IParameterDeclContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ParametersContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ParametersContext) AllParameterDecl() []IParameterDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterDeclContext); ok {
			len++
		}
	}

	tst := make([]IParameterDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterDeclContext); ok {
			tst[i] = t.(IParameterDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) ParameterDecl(i int) IParameterDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclContext); ok {
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

	return t.(IParameterDeclContext)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *TslParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TslParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(TslParserL_PAREN)
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

	if _la == TslParserIDENTIFIER {
		{
			p.SetState(160)
			p.ParameterDecl()
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TslParserCOMMA {
			{
				p.SetState(161)
				p.Match(TslParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(162)
				p.ParameterDecl()
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(170)
		p.Match(TslParserR_PAREN)
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

// IParameterDeclContext is an interface to support dynamic dispatch.
type IParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Type_() ITypeContext

	// IsParameterDeclContext differentiates from other interfaces.
	IsParameterDeclContext()
}

type ParameterDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclContext() *ParameterDeclContext {
	var p = new(ParameterDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_parameterDecl
	return p
}

func InitEmptyParameterDeclContext(p *ParameterDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_parameterDecl
}

func (*ParameterDeclContext) IsParameterDeclContext() {}

func NewParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclContext {
	var p = new(ParameterDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_parameterDecl

	return p
}

func (s *ParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *ParameterDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterParameterDecl(s)
	}
}

func (s *ParameterDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitParameterDecl(s)
	}
}

func (p *TslParser) ParameterDecl() (localctx IParameterDeclContext) {
	localctx = NewParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TslParserRULE_parameterDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Type_()
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

// IResultsContext is an interface to support dynamic dispatch.
type IResultsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsResultsContext differentiates from other interfaces.
	IsResultsContext()
}

type ResultsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultsContext() *ResultsContext {
	var p = new(ResultsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_results
	return p
}

func InitEmptyResultsContext(p *ResultsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_results
}

func (*ResultsContext) IsResultsContext() {}

func NewResultsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultsContext {
	var p = new(ResultsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_results

	return p
}

func (s *ResultsContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultsContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ResultsContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ResultsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *ResultsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *ResultsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterResults(s)
	}
}

func (s *ResultsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitResults(s)
	}
}

func (p *TslParser) Results() (localctx IResultsContext) {
	localctx = NewResultsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TslParserRULE_results)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Type_()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserCOMMA {
		{
			p.SetState(176)
			p.Match(TslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Type_()
		}

		p.SetState(182)
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

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllEnumItem() []IEnumItemContext
	EnumItem(i int) IEnumItemContext

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumDef
	return p
}

func InitEmptyEnumDefContext(p *EnumDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumDef
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(TslParserENUM, 0)
}

func (s *EnumDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *EnumDefContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *EnumDefContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *EnumDefContext) AllEnumItem() []IEnumItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumItemContext); ok {
			len++
		}
	}

	tst := make([]IEnumItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumItemContext); ok {
			tst[i] = t.(IEnumItemContext)
			i++
		}
	}

	return tst
}

func (s *EnumDefContext) EnumItem(i int) IEnumItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumItemContext); ok {
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

	return t.(IEnumItemContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (p *TslParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TslParserRULE_enumDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(TslParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(TslParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserIDENTIFIER {
		{
			p.SetState(186)
			p.EnumItem()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(192)
		p.Match(TslParserR_CURLY)
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

// IEnumItemContext is an interface to support dynamic dispatch.
type IEnumItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	EnumValue() IEnumValueContext
	R_PAREN() antlr.TerminalNode

	// IsEnumItemContext differentiates from other interfaces.
	IsEnumItemContext()
}

type EnumItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumItemContext() *EnumItemContext {
	var p = new(EnumItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumItem
	return p
}

func InitEmptyEnumItemContext(p *EnumItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumItem
}

func (*EnumItemContext) IsEnumItemContext() {}

func NewEnumItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumItemContext {
	var p = new(EnumItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_enumItem

	return p
}

func (s *EnumItemContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *EnumItemContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *EnumItemContext) EnumValue() IEnumValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumItemContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *EnumItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterEnumItem(s)
	}
}

func (s *EnumItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitEnumItem(s)
	}
}

func (p *TslParser) EnumItem() (localctx IEnumItemContext) {
	localctx = NewEnumItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TslParserRULE_enumItem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserL_PAREN {
		{
			p.SetState(195)
			p.Match(TslParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.EnumValue()
		}
		{
			p.SetState(197)
			p.Match(TslParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_VALUE() antlr.TerminalNode
	STRING_VALUE() antlr.TerminalNode

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumValue
	return p
}

func InitEmptyEnumValueContext(p *EnumValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumValue
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, 0)
}

func (s *EnumValueContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING_VALUE, 0)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (p *TslParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TslParserRULE_enumValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TslParserINT_VALUE || _la == TslParserSTRING_VALUE) {
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

// IClassDefContext is an interface to support dynamic dispatch.
type IClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllProperty() []IPropertyContext
	Property(i int) IPropertyContext

	// IsClassDefContext differentiates from other interfaces.
	IsClassDefContext()
}

type ClassDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDefContext() *ClassDefContext {
	var p = new(ClassDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_classDef
	return p
}

func InitEmptyClassDefContext(p *ClassDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_classDef
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) CLASS() antlr.TerminalNode {
	return s.GetToken(TslParserCLASS, 0)
}

func (s *ClassDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *ClassDefContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *ClassDefContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *ClassDefContext) AllProperty() []IPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyContext); ok {
			len++
		}
	}

	tst := make([]IPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyContext); ok {
			tst[i] = t.(IPropertyContext)
			i++
		}
	}

	return tst
}

func (s *ClassDefContext) Property(i int) IPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
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

	return t.(IPropertyContext)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (p *TslParser) ClassDef() (localctx IClassDefContext) {
	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TslParserRULE_classDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(TslParserCLASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(TslParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserIDENTIFIER {
		{
			p.SetState(206)
			p.Property()
		}

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(212)
		p.Match(TslParserR_CURLY)
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

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Type_() ITypeContext

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_property
	return p
}

func InitEmptyPropertyContext(p *PropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_property
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *PropertyContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *TslParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TslParserRULE_property)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Type_()
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	REAL() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TIME() antlr.TerminalNode
	DATE() antlr.TerminalNode
	TimeSeriesType() ITimeSeriesTypeContext
	FqIdentifier() IFqIdentifierContext
	ERROR() antlr.TerminalNode
	ListType() IListTypeContext
	MapType() IMapTypeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(TslParserINT, 0)
}

func (s *TypeContext) REAL() antlr.TerminalNode {
	return s.GetToken(TslParserREAL, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TslParserBOOL, 0)
}

func (s *TypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING, 0)
}

func (s *TypeContext) TIME() antlr.TerminalNode {
	return s.GetToken(TslParserTIME, 0)
}

func (s *TypeContext) DATE() antlr.TerminalNode {
	return s.GetToken(TslParserDATE, 0)
}

func (s *TypeContext) TimeSeriesType() ITimeSeriesTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeSeriesTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeSeriesTypeContext)
}

func (s *TypeContext) FqIdentifier() IFqIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFqIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFqIdentifierContext)
}

func (s *TypeContext) ERROR() antlr.TerminalNode {
	return s.GetToken(TslParserERROR, 0)
}

func (s *TypeContext) ListType() IListTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *TypeContext) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *TslParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TslParserRULE_type)
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TslParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(TslParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(TslParserREAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(219)
			p.Match(TslParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.Match(TslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserTIME:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(221)
			p.Match(TslParserTIME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserDATE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(222)
			p.Match(TslParserDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserTIMESERIES:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(223)
			p.TimeSeriesType()
		}

	case TslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(224)
			p.FqIdentifier()
		}

	case TslParserERROR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(225)
			p.Match(TslParserERROR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserLIST:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(226)
			p.ListType()
		}

	case TslParserMAP:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(227)
			p.MapType()
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

// ITimeSeriesTypeContext is an interface to support dynamic dispatch.
type ITimeSeriesTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TIMESERIES() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	Type_() ITypeContext
	R_PAREN() antlr.TerminalNode

	// IsTimeSeriesTypeContext differentiates from other interfaces.
	IsTimeSeriesTypeContext()
}

type TimeSeriesTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeSeriesTypeContext() *TimeSeriesTypeContext {
	var p = new(TimeSeriesTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_timeSeriesType
	return p
}

func InitEmptyTimeSeriesTypeContext(p *TimeSeriesTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_timeSeriesType
}

func (*TimeSeriesTypeContext) IsTimeSeriesTypeContext() {}

func NewTimeSeriesTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeSeriesTypeContext {
	var p = new(TimeSeriesTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_timeSeriesType

	return p
}

func (s *TimeSeriesTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeSeriesTypeContext) TIMESERIES() antlr.TerminalNode {
	return s.GetToken(TslParserTIMESERIES, 0)
}

func (s *TimeSeriesTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *TimeSeriesTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TimeSeriesTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *TimeSeriesTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeSeriesTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeSeriesTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterTimeSeriesType(s)
	}
}

func (s *TimeSeriesTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitTimeSeriesType(s)
	}
}

func (p *TslParser) TimeSeriesType() (localctx ITimeSeriesTypeContext) {
	localctx = NewTimeSeriesTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TslParserRULE_timeSeriesType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(TslParserTIMESERIES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserL_PAREN {
		{
			p.SetState(231)
			p.Match(TslParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Type_()
		}
		{
			p.SetState(233)
			p.Match(TslParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIST() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	Type_() ITypeContext
	R_PAREN() antlr.TerminalNode

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_listType
	return p
}

func InitEmptyListTypeContext(p *ListTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_listType
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) LIST() antlr.TerminalNode {
	return s.GetToken(TslParserLIST, 0)
}

func (s *ListTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ListTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ListTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitListType(s)
	}
}

func (p *TslParser) ListType() (localctx IListTypeContext) {
	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TslParserRULE_listType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(TslParserLIST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Type_()
	}
	{
		p.SetState(240)
		p.Match(TslParserR_PAREN)
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

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	KeyType() IKeyTypeContext
	COMMA() antlr.TerminalNode
	Type_() ITypeContext
	R_PAREN() antlr.TerminalNode

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_mapType
	return p
}

func InitEmptyMapTypeContext(p *MapTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_mapType
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) MAP() antlr.TerminalNode {
	return s.GetToken(TslParserMAP, 0)
}

func (s *MapTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *MapTypeContext) KeyType() IKeyTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, 0)
}

func (s *MapTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MapTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (p *TslParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TslParserRULE_mapType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(TslParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.KeyType()
	}
	{
		p.SetState(245)
		p.Match(TslParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Type_()
	}
	{
		p.SetState(247)
		p.Match(TslParserR_PAREN)
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

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TIME() antlr.TerminalNode
	DATE() antlr.TerminalNode

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_keyType
	return p
}

func InitEmptyKeyTypeContext(p *KeyTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_keyType
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(TslParserINT, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING, 0)
}

func (s *KeyTypeContext) TIME() antlr.TerminalNode {
	return s.GetToken(TslParserTIME, 0)
}

func (s *KeyTypeContext) DATE() antlr.TerminalNode {
	return s.GetToken(TslParserDATE, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *TslParser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TslParserRULE_keyType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&122406567936) != 0) {
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *BlockContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *TslParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TslParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(TslParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796093121576) != 0 {
		{
			p.SetState(252)
			p.Statement()
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(258)
		p.Match(TslParserR_CURLY)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarsDeclaration() IVarsDeclarationContext
	VarsAssignmentOrFunctionCall() IVarsAssignmentOrFunctionCallContext
	IfStatement() IIfStatementContext
	ReturnStatement() IReturnStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarsDeclaration() IVarsDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsDeclarationContext)
}

func (s *StatementContext) VarsAssignmentOrFunctionCall() IVarsAssignmentOrFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsAssignmentOrFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsAssignmentOrFunctionCallContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *TslParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TslParserRULE_statement)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TslParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.VarsDeclaration()
		}

	case TslParserTHIS, TslParserNEW, TslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.VarsAssignmentOrFunctionCall()
		}

	case TslParserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(262)
			p.IfStatement()
		}

	case TslParserRETURN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(263)
			p.ReturnStatement()
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

// IVarsDeclarationContext is an interface to support dynamic dispatch.
type IVarsDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	AllVarDeclaration() []IVarDeclarationContext
	VarDeclaration(i int) IVarDeclarationContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVarsDeclarationContext differentiates from other interfaces.
	IsVarsDeclarationContext()
}

type VarsDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsDeclarationContext() *VarsDeclarationContext {
	var p = new(VarsDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_varsDeclaration
	return p
}

func InitEmptyVarsDeclarationContext(p *VarsDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_varsDeclaration
}

func (*VarsDeclarationContext) IsVarsDeclarationContext() {}

func NewVarsDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsDeclarationContext {
	var p = new(VarsDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_varsDeclaration

	return p
}

func (s *VarsDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(TslParserVAR, 0)
}

func (s *VarsDeclarationContext) AllVarDeclaration() []IVarDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclarationContext); ok {
			tst[i] = t.(IVarDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *VarsDeclarationContext) VarDeclaration(i int) IVarDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
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

	return t.(IVarDeclarationContext)
}

func (s *VarsDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *VarsDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *VarsDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterVarsDeclaration(s)
	}
}

func (s *VarsDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitVarsDeclaration(s)
	}
}

func (p *TslParser) VarsDeclaration() (localctx IVarsDeclarationContext) {
	localctx = NewVarsDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TslParserRULE_varsDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(TslParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.VarDeclaration()
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserCOMMA {
		{
			p.SetState(268)
			p.Match(TslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.VarDeclaration()
		}

		p.SetState(274)
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

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_varDeclaration
	return p
}

func InitEmptyVarDeclarationContext(p *VarDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_varDeclaration
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *VarDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(TslParserCOLON, 0)
}

func (s *VarDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserEQUAL, 0)
}

func (s *VarDeclarationContext) Expression() IExpressionContext {
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

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (p *TslParser) VarDeclaration() (localctx IVarDeclarationContext) {
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TslParserRULE_varDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserCOLON {
		{
			p.SetState(276)
			p.Match(TslParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Type_()
		}

	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserEQUAL {
		{
			p.SetState(280)
			p.Match(TslParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.expression(0)
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

// IVarsAssignmentOrFunctionCallContext is an interface to support dynamic dispatch.
type IVarsAssignmentOrFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChainedExpression() []IChainedExpressionContext
	ChainedExpression(i int) IChainedExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsVarsAssignmentOrFunctionCallContext differentiates from other interfaces.
	IsVarsAssignmentOrFunctionCallContext()
}

type VarsAssignmentOrFunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsAssignmentOrFunctionCallContext() *VarsAssignmentOrFunctionCallContext {
	var p = new(VarsAssignmentOrFunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_varsAssignmentOrFunctionCall
	return p
}

func InitEmptyVarsAssignmentOrFunctionCallContext(p *VarsAssignmentOrFunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_varsAssignmentOrFunctionCall
}

func (*VarsAssignmentOrFunctionCallContext) IsVarsAssignmentOrFunctionCallContext() {}

func NewVarsAssignmentOrFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsAssignmentOrFunctionCallContext {
	var p = new(VarsAssignmentOrFunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_varsAssignmentOrFunctionCall

	return p
}

func (s *VarsAssignmentOrFunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsAssignmentOrFunctionCallContext) AllChainedExpression() []IChainedExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChainedExpressionContext); ok {
			len++
		}
	}

	tst := make([]IChainedExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChainedExpressionContext); ok {
			tst[i] = t.(IChainedExpressionContext)
			i++
		}
	}

	return tst
}

func (s *VarsAssignmentOrFunctionCallContext) ChainedExpression(i int) IChainedExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainedExpressionContext); ok {
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

	return t.(IChainedExpressionContext)
}

func (s *VarsAssignmentOrFunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *VarsAssignmentOrFunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *VarsAssignmentOrFunctionCallContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserEQUAL, 0)
}

func (s *VarsAssignmentOrFunctionCallContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *VarsAssignmentOrFunctionCallContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *VarsAssignmentOrFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsAssignmentOrFunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsAssignmentOrFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterVarsAssignmentOrFunctionCall(s)
	}
}

func (s *VarsAssignmentOrFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitVarsAssignmentOrFunctionCall(s)
	}
}

func (p *TslParser) VarsAssignmentOrFunctionCall() (localctx IVarsAssignmentOrFunctionCallContext) {
	localctx = NewVarsAssignmentOrFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TslParserRULE_varsAssignmentOrFunctionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.ChainedExpression()
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserCOMMA {
		{
			p.SetState(285)
			p.Match(TslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.ChainedExpression()
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserEQUAL {
		{
			p.SetState(292)
			p.Match(TslParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.expression(0)
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TslParserCOMMA {
			{
				p.SetState(294)
				p.Match(TslParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(295)
				p.expression(0)
			}

			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext
	AllElseIfBlock() []IElseIfBlockContext
	ElseIfBlock(i int) IElseIfBlockContext
	ElseBlock() IElseBlockContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(TslParserIF, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
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

func (s *IfStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) AllElseIfBlock() []IElseIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IElseIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfBlockContext); ok {
			tst[i] = t.(IElseIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIfBlock(i int) IElseIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfBlockContext); ok {
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

	return t.(IElseIfBlockContext)
}

func (s *IfStatementContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *TslParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, TslParserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(TslParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.expression(0)
	}
	{
		p.SetState(305)
		p.Block()
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(306)
				p.ElseIfBlock()
			}

		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserELSE {
		{
			p.SetState(312)
			p.ElseBlock()
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

// IElseIfBlockContext is an interface to support dynamic dispatch.
type IElseIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

	// IsElseIfBlockContext differentiates from other interfaces.
	IsElseIfBlockContext()
}

type ElseIfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfBlockContext() *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_elseIfBlock
	return p
}

func InitEmptyElseIfBlockContext(p *ElseIfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_elseIfBlock
}

func (*ElseIfBlockContext) IsElseIfBlockContext() {}

func NewElseIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_elseIfBlock

	return p
}

func (s *ElseIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TslParserELSE, 0)
}

func (s *ElseIfBlockContext) IF() antlr.TerminalNode {
	return s.GetToken(TslParserIF, 0)
}

func (s *ElseIfBlockContext) Expression() IExpressionContext {
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

func (s *ElseIfBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitElseIfBlock(s)
	}
}

func (p *TslParser) ElseIfBlock() (localctx IElseIfBlockContext) {
	localctx = NewElseIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TslParserRULE_elseIfBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(TslParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Match(TslParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.expression(0)
	}
	{
		p.SetState(318)
		p.Block()
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

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TslParserELSE, 0)
}

func (s *ElseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (p *TslParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, TslParserRULE_elseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(TslParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Block()
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

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(TslParserRETURN, 0)
}

func (s *ReturnStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ReturnStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *ReturnStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *TslParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, TslParserRULE_returnStatement)
	var _la int

	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)
			p.Match(TslParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Match(TslParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.expression(0)
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TslParserCOMMA {
			{
				p.SetState(326)
				p.Match(TslParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(327)
				p.expression(0)
			}

			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	NOT() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	ListExpression() IListExpressionContext
	MapExpression() IMapExpressionContext
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LESS_THAN() antlr.TerminalNode
	LESS_OR_EQUAL() antlr.TerminalNode
	GREATER_THAN() antlr.TerminalNode
	GREATER_OR_EQUAL() antlr.TerminalNode
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

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
	p.RuleIndex = TslParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TslParserNOT, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionContext) ListExpression() IListExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListExpressionContext)
}

func (s *ExpressionContext) MapExpression() IMapExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapExpressionContext)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(TslParserSTAR, 0)
}

func (s *ExpressionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(TslParserSLASH, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TslParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TslParserMINUS, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserEQUAL, 0)
}

func (s *ExpressionContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserNOT_EQUAL, 0)
}

func (s *ExpressionContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(TslParserLESS_THAN, 0)
}

func (s *ExpressionContext) LESS_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserLESS_OR_EQUAL, 0)
}

func (s *ExpressionContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(TslParserGREATER_THAN, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserGREATER_OR_EQUAL, 0)
}

func (s *ExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(TslParserAND)
}

func (s *ExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(TslParserAND, i)
}

func (s *ExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(TslParserOR)
}

func (s *ExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(TslParserOR, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TslParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TslParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, TslParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TslParserTRUE, TslParserFALSE, TslParserTHIS, TslParserNEW, TslParserBAR, TslParserERROR, TslParserIDENTIFIER, TslParserINT_VALUE, TslParserREAL_VALUE, TslParserSTRING_VALUE, TslParserL_PAREN, TslParserPLUS, TslParserMINUS, TslParserAPOS:
		{
			p.SetState(336)
			p.UnaryExpression()
		}

	case TslParserNOT:
		{
			p.SetState(337)
			p.Match(TslParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.expression(5)
		}

	case TslParserL_BRACKET:
		{
			p.SetState(339)
			p.ListExpression()
		}

	case TslParserL_CURLY:
		{
			p.SetState(340)
			p.MapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(366)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(343)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(344)
					_la = p.GetTokenStream().LA(1)

					if !(_la == TslParserSTAR || _la == TslParserSLASH) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(345)
					p.expression(9)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(346)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(347)
					_la = p.GetTokenStream().LA(1)

					if !(_la == TslParserPLUS || _la == TslParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(348)
					p.expression(8)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(349)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(350)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-54)) & ^0x3f) == 0 && ((int64(1)<<(_la-54))&15873) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(351)
					p.expression(7)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(352)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				p.SetState(355)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(353)
							p.Match(TslParserAND)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(354)
							p.expression(0)
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}

					p.SetState(357)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(359)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				p.SetState(362)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(360)
							p.Match(TslParserOR)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(361)
							p.expression(0)
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}

					p.SetState(364)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	ExpressionInParenthesis() IExpressionInParenthesisContext
	BarExpression() IBarExpressionContext
	ConstantValueExpression() IConstantValueExpressionContext
	ChainedExpression() IChainedExpressionContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TslParserMINUS, 0)
}

func (s *UnaryExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TslParserPLUS, 0)
}

func (s *UnaryExpressionContext) ExpressionInParenthesis() IExpressionInParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionInParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionInParenthesisContext)
}

func (s *UnaryExpressionContext) BarExpression() IBarExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBarExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBarExpressionContext)
}

func (s *UnaryExpressionContext) ConstantValueExpression() IConstantValueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantValueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantValueExpressionContext)
}

func (s *UnaryExpressionContext) ChainedExpression() IChainedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChainedExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (p *TslParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, TslParserRULE_unaryExpression)
	var _la int

	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TslParserPLUS, TslParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TslParserPLUS || _la == TslParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(372)
			p.UnaryExpression()
		}

	case TslParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.ExpressionInParenthesis()
		}

	case TslParserBAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(374)
			p.BarExpression()
		}

	case TslParserTRUE, TslParserFALSE, TslParserERROR, TslParserINT_VALUE, TslParserREAL_VALUE, TslParserSTRING_VALUE, TslParserAPOS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(375)
			p.ConstantValueExpression()
		}

	case TslParserTHIS, TslParserNEW, TslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(376)
			p.ChainedExpression()
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

// IExpressionInParenthesisContext is an interface to support dynamic dispatch.
type IExpressionInParenthesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	Expression() IExpressionContext
	R_PAREN() antlr.TerminalNode

	// IsExpressionInParenthesisContext differentiates from other interfaces.
	IsExpressionInParenthesisContext()
}

type ExpressionInParenthesisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionInParenthesisContext() *ExpressionInParenthesisContext {
	var p = new(ExpressionInParenthesisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_expressionInParenthesis
	return p
}

func InitEmptyExpressionInParenthesisContext(p *ExpressionInParenthesisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_expressionInParenthesis
}

func (*ExpressionInParenthesisContext) IsExpressionInParenthesisContext() {}

func NewExpressionInParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionInParenthesisContext {
	var p = new(ExpressionInParenthesisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_expressionInParenthesis

	return p
}

func (s *ExpressionInParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionInParenthesisContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ExpressionInParenthesisContext) Expression() IExpressionContext {
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

func (s *ExpressionInParenthesisContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ExpressionInParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionInParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionInParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterExpressionInParenthesis(s)
	}
}

func (s *ExpressionInParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitExpressionInParenthesis(s)
	}
}

func (p *TslParser) ExpressionInParenthesis() (localctx IExpressionInParenthesisContext) {
	localctx = NewExpressionInParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, TslParserRULE_expressionInParenthesis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
		p.expression(0)
	}
	{
		p.SetState(381)
		p.Match(TslParserR_PAREN)
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

// IChainedExpressionContext is an interface to support dynamic dispatch.
type IChainedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChainItem() []IChainItemContext
	ChainItem(i int) IChainItemContext
	THIS() antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	NEW() antlr.TerminalNode
	FqIdentifier() IFqIdentifierContext
	ParamsExpression() IParamsExpressionContext

	// IsChainedExpressionContext differentiates from other interfaces.
	IsChainedExpressionContext()
}

type ChainedExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainedExpressionContext() *ChainedExpressionContext {
	var p = new(ChainedExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_chainedExpression
	return p
}

func InitEmptyChainedExpressionContext(p *ChainedExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_chainedExpression
}

func (*ChainedExpressionContext) IsChainedExpressionContext() {}

func NewChainedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainedExpressionContext {
	var p = new(ChainedExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_chainedExpression

	return p
}

func (s *ChainedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainedExpressionContext) AllChainItem() []IChainItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChainItemContext); ok {
			len++
		}
	}

	tst := make([]IChainItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChainItemContext); ok {
			tst[i] = t.(IChainItemContext)
			i++
		}
	}

	return tst
}

func (s *ChainedExpressionContext) ChainItem(i int) IChainItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainItemContext); ok {
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

	return t.(IChainItemContext)
}

func (s *ChainedExpressionContext) THIS() antlr.TerminalNode {
	return s.GetToken(TslParserTHIS, 0)
}

func (s *ChainedExpressionContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TslParserDOT)
}

func (s *ChainedExpressionContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TslParserDOT, i)
}

func (s *ChainedExpressionContext) NEW() antlr.TerminalNode {
	return s.GetToken(TslParserNEW, 0)
}

func (s *ChainedExpressionContext) FqIdentifier() IFqIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFqIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFqIdentifierContext)
}

func (s *ChainedExpressionContext) ParamsExpression() IParamsExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsExpressionContext)
}

func (s *ChainedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterChainedExpression(s)
	}
}

func (s *ChainedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitChainedExpression(s)
	}
}

func (p *TslParser) ChainedExpression() (localctx IChainedExpressionContext) {
	localctx = NewChainedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, TslParserRULE_chainedExpression)
	var _la int

	var _alt int

	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TslParserTHIS, TslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TslParserTHIS {
			{
				p.SetState(383)
				p.Match(TslParserTHIS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(384)
				p.Match(TslParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(387)
			p.ChainItem()
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(388)
					p.Match(TslParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(389)
					p.ChainItem()
				}

			}
			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case TslParserNEW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.Match(TslParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.FqIdentifier()
		}
		{
			p.SetState(397)
			p.ParamsExpression()
		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(398)
					p.Match(TslParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(399)
					p.ChainItem()
				}

			}
			p.SetState(404)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
			if p.HasError() {
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

// IChainItemContext is an interface to support dynamic dispatch.
type IChainItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ParamsExpression() IParamsExpressionContext
	AccessorExpression() IAccessorExpressionContext

	// IsChainItemContext differentiates from other interfaces.
	IsChainItemContext()
}

type ChainItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainItemContext() *ChainItemContext {
	var p = new(ChainItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_chainItem
	return p
}

func InitEmptyChainItemContext(p *ChainItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_chainItem
}

func (*ChainItemContext) IsChainItemContext() {}

func NewChainItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainItemContext {
	var p = new(ChainItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_chainItem

	return p
}

func (s *ChainItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *ChainItemContext) ParamsExpression() IParamsExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsExpressionContext)
}

func (s *ChainItemContext) AccessorExpression() IAccessorExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessorExpressionContext)
}

func (s *ChainItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterChainItem(s)
	}
}

func (s *ChainItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitChainItem(s)
	}
}

func (p *TslParser) ChainItem() (localctx IChainItemContext) {
	localctx = NewChainItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, TslParserRULE_chainItem)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(408)
			p.ParamsExpression()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(411)
			p.AccessorExpression()
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

// IParamsExpressionContext is an interface to support dynamic dispatch.
type IParamsExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamsExpressionContext differentiates from other interfaces.
	IsParamsExpressionContext()
}

type ParamsExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsExpressionContext() *ParamsExpressionContext {
	var p = new(ParamsExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_paramsExpression
	return p
}

func InitEmptyParamsExpressionContext(p *ParamsExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_paramsExpression
}

func (*ParamsExpressionContext) IsParamsExpressionContext() {}

func NewParamsExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsExpressionContext {
	var p = new(ParamsExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_paramsExpression

	return p
}

func (s *ParamsExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ParamsExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ParamsExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ParamsExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ParamsExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *ParamsExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *ParamsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterParamsExpression(s)
	}
}

func (s *ParamsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitParamsExpression(s)
	}
}

func (p *TslParser) ParamsExpression() (localctx IParamsExpressionContext) {
	localctx = NewParamsExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, TslParserRULE_paramsExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5050145966817256192) != 0 {
		{
			p.SetState(415)
			p.expression(0)
		}
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TslParserCOMMA {
			{
				p.SetState(416)
				p.Match(TslParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(417)
				p.expression(0)
			}

			p.SetState(422)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(425)
		p.Match(TslParserR_PAREN)
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

// IAccessorExpressionContext is an interface to support dynamic dispatch.
type IAccessorExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	R_BRACKET() antlr.TerminalNode

	// IsAccessorExpressionContext differentiates from other interfaces.
	IsAccessorExpressionContext()
}

type AccessorExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorExpressionContext() *AccessorExpressionContext {
	var p = new(AccessorExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_accessorExpression
	return p
}

func InitEmptyAccessorExpressionContext(p *AccessorExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_accessorExpression
}

func (*AccessorExpressionContext) IsAccessorExpressionContext() {}

func NewAccessorExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorExpressionContext {
	var p = new(AccessorExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_accessorExpression

	return p
}

func (s *AccessorExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorExpressionContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(TslParserL_BRACKET, 0)
}

func (s *AccessorExpressionContext) Expression() IExpressionContext {
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

func (s *AccessorExpressionContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(TslParserR_BRACKET, 0)
}

func (s *AccessorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterAccessorExpression(s)
	}
}

func (s *AccessorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitAccessorExpression(s)
	}
}

func (p *TslParser) AccessorExpression() (localctx IAccessorExpressionContext) {
	localctx = NewAccessorExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, TslParserRULE_accessorExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(427)
		p.Match(TslParserL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.expression(0)
	}
	{
		p.SetState(429)
		p.Match(TslParserR_BRACKET)
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

// IBarExpressionContext is an interface to support dynamic dispatch.
type IBarExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BAR() antlr.TerminalNode
	DOT() antlr.TerminalNode
	OPEN() antlr.TerminalNode
	HIGH() antlr.TerminalNode
	LOW() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	AccessorExpression() IAccessorExpressionContext

	// IsBarExpressionContext differentiates from other interfaces.
	IsBarExpressionContext()
}

type BarExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBarExpressionContext() *BarExpressionContext {
	var p = new(BarExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_barExpression
	return p
}

func InitEmptyBarExpressionContext(p *BarExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_barExpression
}

func (*BarExpressionContext) IsBarExpressionContext() {}

func NewBarExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BarExpressionContext {
	var p = new(BarExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_barExpression

	return p
}

func (s *BarExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BarExpressionContext) BAR() antlr.TerminalNode {
	return s.GetToken(TslParserBAR, 0)
}

func (s *BarExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(TslParserDOT, 0)
}

func (s *BarExpressionContext) OPEN() antlr.TerminalNode {
	return s.GetToken(TslParserOPEN, 0)
}

func (s *BarExpressionContext) HIGH() antlr.TerminalNode {
	return s.GetToken(TslParserHIGH, 0)
}

func (s *BarExpressionContext) LOW() antlr.TerminalNode {
	return s.GetToken(TslParserLOW, 0)
}

func (s *BarExpressionContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(TslParserCLOSE, 0)
}

func (s *BarExpressionContext) AccessorExpression() IAccessorExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessorExpressionContext)
}

func (s *BarExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BarExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BarExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterBarExpression(s)
	}
}

func (s *BarExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitBarExpression(s)
	}
}

func (p *TslParser) BarExpression() (localctx IBarExpressionContext) {
	localctx = NewBarExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, TslParserRULE_barExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.Match(TslParserBAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(432)
		p.Match(TslParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3932160) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(434)
			p.AccessorExpression()
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

// IListExpressionContext is an interface to support dynamic dispatch.
type IListExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	R_BRACKET() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListExpressionContext differentiates from other interfaces.
	IsListExpressionContext()
}

type ListExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListExpressionContext() *ListExpressionContext {
	var p = new(ListExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_listExpression
	return p
}

func InitEmptyListExpressionContext(p *ListExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_listExpression
}

func (*ListExpressionContext) IsListExpressionContext() {}

func NewListExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListExpressionContext {
	var p = new(ListExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_listExpression

	return p
}

func (s *ListExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ListExpressionContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(TslParserL_BRACKET, 0)
}

func (s *ListExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ListExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ListExpressionContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(TslParserR_BRACKET, 0)
}

func (s *ListExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *ListExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *ListExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterListExpression(s)
	}
}

func (s *ListExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitListExpression(s)
	}
}

func (p *TslParser) ListExpression() (localctx IListExpressionContext) {
	localctx = NewListExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, TslParserRULE_listExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.Match(TslParserL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
		p.expression(0)
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserCOMMA {
		{
			p.SetState(439)
			p.Match(TslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.expression(0)
		}

		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(446)
		p.Match(TslParserR_BRACKET)
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

// IMapExpressionContext is an interface to support dynamic dispatch.
type IMapExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	AllKeyValueCouple() []IKeyValueCoupleContext
	KeyValueCouple(i int) IKeyValueCoupleContext
	R_CURLY() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMapExpressionContext differentiates from other interfaces.
	IsMapExpressionContext()
}

type MapExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapExpressionContext() *MapExpressionContext {
	var p = new(MapExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_mapExpression
	return p
}

func InitEmptyMapExpressionContext(p *MapExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_mapExpression
}

func (*MapExpressionContext) IsMapExpressionContext() {}

func NewMapExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapExpressionContext {
	var p = new(MapExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_mapExpression

	return p
}

func (s *MapExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MapExpressionContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *MapExpressionContext) AllKeyValueCouple() []IKeyValueCoupleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeyValueCoupleContext); ok {
			len++
		}
	}

	tst := make([]IKeyValueCoupleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeyValueCoupleContext); ok {
			tst[i] = t.(IKeyValueCoupleContext)
			i++
		}
	}

	return tst
}

func (s *MapExpressionContext) KeyValueCouple(i int) IKeyValueCoupleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValueCoupleContext); ok {
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

	return t.(IKeyValueCoupleContext)
}

func (s *MapExpressionContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *MapExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *MapExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *MapExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterMapExpression(s)
	}
}

func (s *MapExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitMapExpression(s)
	}
}

func (p *TslParser) MapExpression() (localctx IMapExpressionContext) {
	localctx = NewMapExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, TslParserRULE_mapExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(TslParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.KeyValueCouple()
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserCOMMA {
		{
			p.SetState(450)
			p.Match(TslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)
			p.KeyValueCouple()
		}

		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(457)
		p.Match(TslParserR_CURLY)
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

// IKeyValueCoupleContext is an interface to support dynamic dispatch.
type IKeyValueCoupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KeyValue() IKeyValueContext
	COLON() antlr.TerminalNode
	Expression() IExpressionContext

	// IsKeyValueCoupleContext differentiates from other interfaces.
	IsKeyValueCoupleContext()
}

type KeyValueCoupleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueCoupleContext() *KeyValueCoupleContext {
	var p = new(KeyValueCoupleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_keyValueCouple
	return p
}

func InitEmptyKeyValueCoupleContext(p *KeyValueCoupleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_keyValueCouple
}

func (*KeyValueCoupleContext) IsKeyValueCoupleContext() {}

func NewKeyValueCoupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueCoupleContext {
	var p = new(KeyValueCoupleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_keyValueCouple

	return p
}

func (s *KeyValueCoupleContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueCoupleContext) KeyValue() IKeyValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *KeyValueCoupleContext) COLON() antlr.TerminalNode {
	return s.GetToken(TslParserCOLON, 0)
}

func (s *KeyValueCoupleContext) Expression() IExpressionContext {
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

func (s *KeyValueCoupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueCoupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueCoupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterKeyValueCouple(s)
	}
}

func (s *KeyValueCoupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitKeyValueCouple(s)
	}
}

func (p *TslParser) KeyValueCouple() (localctx IKeyValueCoupleContext) {
	localctx = NewKeyValueCoupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, TslParserRULE_keyValueCouple)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(459)
		p.KeyValue()
	}
	{
		p.SetState(460)
		p.Match(TslParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(461)
		p.expression(0)
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

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_VALUE() antlr.TerminalNode
	STRING_VALUE() antlr.TerminalNode
	TimeValue() ITimeValueContext
	DateValue() IDateValueContext

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_keyValue
	return p
}

func InitEmptyKeyValueContext(p *KeyValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_keyValue
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, 0)
}

func (s *KeyValueContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING_VALUE, 0)
}

func (s *KeyValueContext) TimeValue() ITimeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
}

func (s *KeyValueContext) DateValue() IDateValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
}

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (p *TslParser) KeyValue() (localctx IKeyValueContext) {
	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, TslParserRULE_keyValue)
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.Match(TslParserINT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(464)
			p.Match(TslParserSTRING_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(465)
			p.TimeValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(466)
			p.DateValue()
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

// IConstantValueExpressionContext is an interface to support dynamic dispatch.
type IConstantValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_VALUE() antlr.TerminalNode
	REAL_VALUE() antlr.TerminalNode
	BoolValue() IBoolValueContext
	STRING_VALUE() antlr.TerminalNode
	TimeValue() ITimeValueContext
	DateValue() IDateValueContext
	ErrorValue() IErrorValueContext

	// IsConstantValueExpressionContext differentiates from other interfaces.
	IsConstantValueExpressionContext()
}

type ConstantValueExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantValueExpressionContext() *ConstantValueExpressionContext {
	var p = new(ConstantValueExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantValueExpression
	return p
}

func InitEmptyConstantValueExpressionContext(p *ConstantValueExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantValueExpression
}

func (*ConstantValueExpressionContext) IsConstantValueExpressionContext() {}

func NewConstantValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantValueExpressionContext {
	var p = new(ConstantValueExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_constantValueExpression

	return p
}

func (s *ConstantValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantValueExpressionContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, 0)
}

func (s *ConstantValueExpressionContext) REAL_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserREAL_VALUE, 0)
}

func (s *ConstantValueExpressionContext) BoolValue() IBoolValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolValueContext)
}

func (s *ConstantValueExpressionContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING_VALUE, 0)
}

func (s *ConstantValueExpressionContext) TimeValue() ITimeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
}

func (s *ConstantValueExpressionContext) DateValue() IDateValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
}

func (s *ConstantValueExpressionContext) ErrorValue() IErrorValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorValueContext)
}

func (s *ConstantValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterConstantValueExpression(s)
	}
}

func (s *ConstantValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitConstantValueExpression(s)
	}
}

func (p *TslParser) ConstantValueExpression() (localctx IConstantValueExpressionContext) {
	localctx = NewConstantValueExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, TslParserRULE_constantValueExpression)
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(469)
			p.Match(TslParserINT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(470)
			p.Match(TslParserREAL_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(471)
			p.BoolValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(472)
			p.Match(TslParserSTRING_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(473)
			p.TimeValue()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(474)
			p.DateValue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(475)
			p.ErrorValue()
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

// ITimeValueContext is an interface to support dynamic dispatch.
type ITimeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAPOS() []antlr.TerminalNode
	APOS(i int) antlr.TerminalNode
	AllINT_VALUE() []antlr.TerminalNode
	INT_VALUE(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsTimeValueContext differentiates from other interfaces.
	IsTimeValueContext()
}

type TimeValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeValueContext() *TimeValueContext {
	var p = new(TimeValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_timeValue
	return p
}

func InitEmptyTimeValueContext(p *TimeValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_timeValue
}

func (*TimeValueContext) IsTimeValueContext() {}

func NewTimeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeValueContext {
	var p = new(TimeValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_timeValue

	return p
}

func (s *TimeValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeValueContext) AllAPOS() []antlr.TerminalNode {
	return s.GetTokens(TslParserAPOS)
}

func (s *TimeValueContext) APOS(i int) antlr.TerminalNode {
	return s.GetToken(TslParserAPOS, i)
}

func (s *TimeValueContext) AllINT_VALUE() []antlr.TerminalNode {
	return s.GetTokens(TslParserINT_VALUE)
}

func (s *TimeValueContext) INT_VALUE(i int) antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, i)
}

func (s *TimeValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(TslParserCOLON, 0)
}

func (s *TimeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterTimeValue(s)
	}
}

func (s *TimeValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitTimeValue(s)
	}
}

func (p *TslParser) TimeValue() (localctx ITimeValueContext) {
	localctx = NewTimeValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, TslParserRULE_timeValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(478)
		p.Match(TslParserAPOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(479)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(TslParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(482)
		p.Match(TslParserAPOS)
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

// IDateValueContext is an interface to support dynamic dispatch.
type IDateValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAPOS() []antlr.TerminalNode
	APOS(i int) antlr.TerminalNode
	AllINT_VALUE() []antlr.TerminalNode
	INT_VALUE(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsDateValueContext differentiates from other interfaces.
	IsDateValueContext()
}

type DateValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateValueContext() *DateValueContext {
	var p = new(DateValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_dateValue
	return p
}

func InitEmptyDateValueContext(p *DateValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_dateValue
}

func (*DateValueContext) IsDateValueContext() {}

func NewDateValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateValueContext {
	var p = new(DateValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_dateValue

	return p
}

func (s *DateValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateValueContext) AllAPOS() []antlr.TerminalNode {
	return s.GetTokens(TslParserAPOS)
}

func (s *DateValueContext) APOS(i int) antlr.TerminalNode {
	return s.GetToken(TslParserAPOS, i)
}

func (s *DateValueContext) AllINT_VALUE() []antlr.TerminalNode {
	return s.GetTokens(TslParserINT_VALUE)
}

func (s *DateValueContext) INT_VALUE(i int) antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, i)
}

func (s *DateValueContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(TslParserMINUS)
}

func (s *DateValueContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(TslParserMINUS, i)
}

func (s *DateValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterDateValue(s)
	}
}

func (s *DateValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitDateValue(s)
	}
}

func (p *TslParser) DateValue() (localctx IDateValueContext) {
	localctx = NewDateValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, TslParserRULE_dateValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(484)
		p.Match(TslParserAPOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(485)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.Match(TslParserMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(487)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.Match(TslParserMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(490)
		p.Match(TslParserAPOS)
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

// IBoolValueContext is an interface to support dynamic dispatch.
type IBoolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBoolValueContext differentiates from other interfaces.
	IsBoolValueContext()
}

type BoolValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolValueContext() *BoolValueContext {
	var p = new(BoolValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_boolValue
	return p
}

func InitEmptyBoolValueContext(p *BoolValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_boolValue
}

func (*BoolValueContext) IsBoolValueContext() {}

func NewBoolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolValueContext {
	var p = new(BoolValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_boolValue

	return p
}

func (s *BoolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TslParserTRUE, 0)
}

func (s *BoolValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TslParserFALSE, 0)
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (p *TslParser) BoolValue() (localctx IBoolValueContext) {
	localctx = NewBoolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, TslParserRULE_boolValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TslParserTRUE || _la == TslParserFALSE) {
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

// IErrorValueContext is an interface to support dynamic dispatch.
type IErrorValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ERROR() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	STRING_VALUE() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode

	// IsErrorValueContext differentiates from other interfaces.
	IsErrorValueContext()
}

type ErrorValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorValueContext() *ErrorValueContext {
	var p = new(ErrorValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_errorValue
	return p
}

func InitEmptyErrorValueContext(p *ErrorValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_errorValue
}

func (*ErrorValueContext) IsErrorValueContext() {}

func NewErrorValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorValueContext {
	var p = new(ErrorValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_errorValue

	return p
}

func (s *ErrorValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorValueContext) ERROR() antlr.TerminalNode {
	return s.GetToken(TslParserERROR, 0)
}

func (s *ErrorValueContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ErrorValueContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING_VALUE, 0)
}

func (s *ErrorValueContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ErrorValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterErrorValue(s)
	}
}

func (s *ErrorValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitErrorValue(s)
	}
}

func (p *TslParser) ErrorValue() (localctx IErrorValueContext) {
	localctx = NewErrorValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, TslParserRULE_errorValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
		p.Match(TslParserERROR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(495)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(496)
		p.Match(TslParserSTRING_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(497)
		p.Match(TslParserR_PAREN)
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

// IFqIdentifierContext is an interface to support dynamic dispatch.
type IFqIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsFqIdentifierContext differentiates from other interfaces.
	IsFqIdentifierContext()
}

type FqIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFqIdentifierContext() *FqIdentifierContext {
	var p = new(FqIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_fqIdentifier
	return p
}

func InitEmptyFqIdentifierContext(p *FqIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_fqIdentifier
}

func (*FqIdentifierContext) IsFqIdentifierContext() {}

func NewFqIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FqIdentifierContext {
	var p = new(FqIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_fqIdentifier

	return p
}

func (s *FqIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FqIdentifierContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(TslParserIDENTIFIER)
}

func (s *FqIdentifierContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, i)
}

func (s *FqIdentifierContext) DOT() antlr.TerminalNode {
	return s.GetToken(TslParserDOT, 0)
}

func (s *FqIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FqIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FqIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterFqIdentifier(s)
	}
}

func (s *FqIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitFqIdentifier(s)
	}
}

func (p *TslParser) FqIdentifier() (localctx IFqIdentifierContext) {
	localctx = NewFqIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, TslParserRULE_fqIdentifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserDOT {
		{
			p.SetState(500)
			p.Match(TslParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(501)
			p.Match(TslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

func (p *TslParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 30:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TslParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
