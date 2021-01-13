// Code generated by goyacc -o parser.y.go parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

//line parser.y:6
type yySymType struct {
	yys     int
	node    Node
	nodes   []Node
	item    Item
	strings []string
	float   float64
}

const SEMICOLON = 57346
const COMMA = 57347
const COMMENT = 57348
const EOF = 57349
const ERROR = 57350
const ID = 57351
const LEFT_PAREN = 57352
const LEFT_BRACKET = 57353
const NUMBER = 57354
const RIGHT_PAREN = 57355
const RIGHT_BRACKET = 57356
const SPACE = 57357
const STRING = 57358
const QUOTED_STRING = 57359
const operatorsStart = 57360
const ADD = 57361
const DIV = 57362
const GTE = 57363
const GT = 57364
const LT = 57365
const LTE = 57366
const MOD = 57367
const MUL = 57368
const NEQ = 57369
const EQ = 57370
const POW = 57371
const SUB = 57372
const operatorsEnd = 57373
const keywordsStart = 57374
const TRUE = 57375
const FALSE = 57376
const IDENTIFIER = 57377
const AND = 57378
const OR = 57379
const NIL = 57380
const NULL = 57381
const RE = 57382
const JP = 57383
const keywordsEnd = 57384
const startSymbolsStart = 57385
const START_FUNC_EXPRESSION = 57386
const startSymbolsEnd = 57387

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SEMICOLON",
	"COMMA",
	"COMMENT",
	"EOF",
	"ERROR",
	"ID",
	"LEFT_PAREN",
	"LEFT_BRACKET",
	"NUMBER",
	"RIGHT_PAREN",
	"RIGHT_BRACKET",
	"SPACE",
	"STRING",
	"QUOTED_STRING",
	"operatorsStart",
	"ADD",
	"DIV",
	"GTE",
	"GT",
	"LT",
	"LTE",
	"MOD",
	"MUL",
	"NEQ",
	"EQ",
	"POW",
	"SUB",
	"operatorsEnd",
	"keywordsStart",
	"TRUE",
	"FALSE",
	"IDENTIFIER",
	"AND",
	"OR",
	"NIL",
	"NULL",
	"RE",
	"JP",
	"keywordsEnd",
	"startSymbolsStart",
	"START_FUNC_EXPRESSION",
	"startSymbolsEnd",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:347

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 38,
	10, 51,
	-2, 13,
}

const yyPrivate = 57344

const yyLast = 200

var yyAct = [...]int{
	38, 21, 19, 8, 28, 3, 8, 59, 37, 9,
	34, 20, 35, 8, 37, 90, 37, 10, 12, 43,
	45, 64, 63, 9, 9, 96, 18, 37, 88, 95,
	44, 10, 10, 41, 42, 11, 86, 67, 39, 40,
	32, 33, 94, 93, 68, 85, 69, 2, 66, 11,
	11, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 82, 83, 84, 9, 34, 65, 35, 47,
	87, 89, 37, 10, 70, 43, 15, 46, 14, 4,
	1, 29, 27, 30, 31, 23, 44, 64, 92, 41,
	42, 11, 91, 22, 39, 40, 32, 33, 48, 49,
	50, 51, 54, 55, 56, 57, 58, 61, 59, 60,
	24, 26, 62, 5, 17, 52, 53, 48, 49, 50,
	51, 54, 55, 56, 57, 58, 61, 59, 60, 9,
	7, 36, 35, 0, 52, 53, 37, 10, 0, 43,
	0, 48, 49, 25, 0, 0, 6, 56, 57, 13,
	44, 59, 60, 41, 42, 11, 16, 0, 39, 40,
	48, 49, 50, 51, 54, 55, 56, 57, 58, 61,
	59, 60, 0, 0, 0, 0, 0, 52, 48, 49,
	50, 51, 54, 55, 56, 57, 58, 61, 59, 60,
	49, 0, 0, 0, 0, 56, 57, 0, 0, 59,
}

var yyPact = [...]int{
	3, 72, 15, -1000, -1000, 14, -1000, 68, -1000, -1000,
	-1000, 66, 15, -1000, 0, -8, -1000, 64, -1000, 98,
	120, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 57, 38, 56, -1000, 32, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 33, -1000, 0, 56, 56,
	56, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 31, -1000, -1000, 11, -2, 79, -1000, -1000,
	-1000, 170, -22, 122, 122, 159, 141, 122, 122, -22,
	-22, 122, -22, 170, 122, -1000, 120, 30, 29, 16,
	12, -1000, -1000, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]int{
	0, 131, 130, 0, 114, 113, 1, 112, 111, 2,
	26, 143, 110, 93, 85, 84, 4, 83, 82, 81,
	80,
}

var yyR1 = [...]int{
	0, 20, 20, 20, 5, 5, 5, 9, 9, 9,
	9, 9, 9, 19, 1, 1, 16, 17, 17, 15,
	15, 12, 11, 4, 4, 4, 4, 10, 10, 7,
	7, 7, 6, 6, 6, 6, 6, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 2, 18, 18, 13, 13, 14, 14, 3, 3,
	3,
}

var yyR2 = [...]int{
	0, 2, 2, 1, 1, 3, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 4, 3, 2, 1, 0, 1, 3, 3,
	1, 0, 1, 1, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 1, 2, 4, 4, 4, 4, 1, 1,
	4,
}

var yyChk = [...]int{
	-1000, -20, 44, 2, 7, -5, -11, -2, -3, 9,
	17, 35, 4, -11, 10, 10, -11, -4, -10, -9,
	11, -6, -13, -14, -12, -11, -8, -18, -16, -19,
	-17, -15, 40, 41, 10, 12, -1, 16, -3, 38,
	39, 33, 34, 19, 30, -16, 13, 5, 19, 20,
	21, 22, 36, 37, 23, 24, 25, 26, 27, 29,
	30, 28, -7, -6, -3, 10, 10, -9, 12, 13,
	-10, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, 14, 5, -16, 17, -16,
	17, 13, -6, 13, 13, 13, 13,
}

var yyDef = [...]int{
	0, -2, 0, 3, 2, 1, 4, 0, 51, 58,
	59, 0, 0, 6, 26, 0, 5, 0, 25, 27,
	31, 7, 8, 9, 10, 11, 12, 32, 33, 34,
	35, 36, 0, 0, 0, 52, 0, 16, -2, 17,
	18, 19, 20, 14, 15, 0, 22, 24, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 30, 13, 0, 0, 0, 53, 60,
	23, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 50, 28, 0, 0, 0, 0,
	0, 21, 29, 54, 55, 56, 57,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:80
		{
			yylex.(*parser).parseResult = yyDollar[2].nodes
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:85
		{
			yylex.(*parser).unexpected("", "")
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:91
		{
			yyVAL.nodes = Funcs{yyDollar[1].node}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:95
		{
			yyDollar[1].nodes = append(yyDollar[1].nodes, yyDollar[3].node)
			yyVAL.nodes = yyDollar[1].nodes
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:100
		{
			yyDollar[1].nodes = append(yyDollar[1].nodes, yyDollar[2].node)
			yyVAL.nodes = yyDollar[1].nodes
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:111
		{
			yyVAL.node = &Identifier{Name: yyDollar[1].item.Val}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:122
		{
			yyVAL.node = &StringLiteral{Val: yylex.(*parser).unquoteString(yyDollar[1].item.Val)}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:128
		{
			yyVAL.node = &NilLiteral{}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:132
		{
			yyVAL.node = &NilLiteral{}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:138
		{
			yyVAL.node = &BoolLiteral{Val: true}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:142
		{
			yyVAL.node = &BoolLiteral{Val: false}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:148
		{
			yyVAL.node = &ParenExpr{Param: yyDollar[2].node}
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:154
		{
			yyVAL.node = yylex.(*parser).newFunc(yyDollar[1].item.Val, yyDollar[3].nodes)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:162
		{
			yyVAL.nodes = append(yyVAL.nodes, yyDollar[3].node)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:167
		{
			yyVAL.nodes = []Node{yyDollar[1].node}
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:171
		{
			yyVAL.nodes = nil
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:178
		{
			yyVAL.node = yyDollar[1].node
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:182
		{
			yyVAL.node = getFuncArgList(yyDollar[2].node.(NodeList))
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:188
		{
			nl := yyVAL.node.(NodeList)
			nl = append(nl, yyDollar[3].node)
			yyVAL.node = nl
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:194
		{
			yyVAL.node = NodeList{yyDollar[1].node}
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:198
		{
			yyVAL.node = NodeList{}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:211
		{
			yyVAL.node = yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:215
		{
			yyVAL.node = yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:219
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:225
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:231
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:237
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:243
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:249
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:255
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:260
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:265
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:271
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:276
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:281
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:291
		{
			yyVAL.item = yyDollar[1].item
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:298
		{
			yyVAL.node = yylex.(*parser).number(yyDollar[1].item.Val)
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:302
		{
			num := yylex.(*parser).number(yyDollar[2].item.Val)
			switch yyDollar[1].item.Typ {
			case ADD: // pass
			case SUB:
				if num.IsInt {
					num.Int = -num.Int
				} else {
					num.Float = -num.Float
				}
			}
			yyVAL.node = num
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:318
		{
			yyVAL.node = &Regex{Regex: yyDollar[3].node.(*StringLiteral).Val}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:322
		{
			yyVAL.node = &Regex{Regex: yylex.(*parser).unquoteString(yyDollar[3].item.Val)}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:328
		{
			yyVAL.node = &Jspath{Jspath: yyDollar[3].node.(*StringLiteral).Val}
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:332
		{
			yyVAL.node = &Jspath{Jspath: yylex.(*parser).unquoteString(yyDollar[3].item.Val)}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:339
		{
			yyVAL.item.Val = yylex.(*parser).unquoteString(yyDollar[1].item.Val)
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:343
		{
			yyVAL.item.Val = yyDollar[3].node.(*StringLiteral).Val
		}
	}
	goto yystack /* stack new state and value */
}
