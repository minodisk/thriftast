//line ./parser/parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line ./parser/parser.go.y:2
import "github.com/minodisk/thriftast/ast"

//line ./parser/parser.go.y:8
type yySymType struct {
	yys         int
	program     *ast.Program
	expressions []ast.Expression
	include     *ast.Include
	namespace   *ast.Namespace
	typedef     *ast.Typedef
	Const       *ast.Const
	Struct      *ast.Struct
	fields      []*ast.Field
	field       *ast.Field
	req         *ast.Req
	ident       *ast.Ident
	value       ast.Value
	string      *ast.String
	int         *ast.Int
	float       *ast.Float
	equal       *ast.Equal
	dot         *ast.Dot
	colon       *ast.Colon
	brace       *ast.Brace
}

const INCLUDE = 57346
const NAMESPACE = 57347
const TYPEDEF = 57348
const CONST = 57349
const STRUCT = 57350
const REQUIRED = 57351
const OPTIONAL = 57352
const SERVICE = 57353
const ONEWAY = 57354
const IDENT = 57355
const STRING = 57356
const INT = 57357
const FLOAT = 57358
const EQUAL = 57359
const DOT = 57360
const COLON = 57361
const LBRACE = 57362
const RBRACE = 57363

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INCLUDE",
	"NAMESPACE",
	"TYPEDEF",
	"CONST",
	"STRUCT",
	"REQUIRED",
	"OPTIONAL",
	"SERVICE",
	"ONEWAY",
	"IDENT",
	"STRING",
	"INT",
	"FLOAT",
	"EQUAL",
	"DOT",
	"COLON",
	"LBRACE",
	"RBRACE",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser/parser.go.y:221

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 47

var yyAct = [...]int{

	21, 31, 35, 26, 36, 32, 33, 34, 22, 14,
	16, 17, 18, 30, 15, 20, 15, 23, 24, 25,
	42, 15, 15, 13, 15, 27, 19, 8, 9, 10,
	11, 12, 37, 38, 39, 29, 28, 7, 40, 6,
	5, 41, 4, 3, 43, 2, 1,
}
var yyPact = [...]int{

	23, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 9, 11,
	11, 11, 13, -1000, 1, -10, 11, 11, 11, -17,
	-1000, 11, -1000, 11, 11, 8, -2, -9, -19, -1000,
	-15, -1000, -1000, -1000, -1000, -1000, 24, 11, -1000, -1000,
	11, 3, -9, -1000,
}
var yyPgo = [...]int{

	0, 46, 45, 43, 42, 40, 39, 37, 36, 35,
	32, 0, 1,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 3, 3,
	4, 5, 6, 7, 8, 10, 10, 10, 9, 9,
	12, 12, 12, 11, 11, 11,
}
var yyR2 = [...]int{

	0, 1, 0, 1, 1, 1, 1, 1, 2, 3,
	3, 3, 5, 5, 1, 0, 1, 1, 5, 7,
	1, 1, 1, 1, 2, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 4, 5,
	6, 7, 8, 14, -11, 13, -11, -11, -11, 13,
	14, -11, 18, -11, -11, -11, 20, 17, -8, -9,
	15, -12, 14, 15, 16, 21, 19, -10, 9, 10,
	-11, -11, 17, -12,
}
var yyDef = [...]int{

	2, -2, 1, 3, 4, 5, 6, 7, 0, 0,
	0, 0, 0, 8, 0, 23, 0, 0, 0, 0,
	9, 25, 24, 10, 11, 0, 0, 0, 0, 14,
	0, 12, 20, 21, 22, 13, 15, 0, 16, 17,
	0, 18, 0, 19,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:70
		{
			yyVAL.program = &ast.Program{Expressions: yyDollar[1].expressions}
			yylex.(*Lexer).Program = yyVAL.program
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.go.y:77
		{
			yyVAL.expressions = nil
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:81
		{
			yyVAL.expressions = append(yyVAL.expressions, yyDollar[1].include)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:85
		{
			yyVAL.expressions = append(yyVAL.expressions, yyDollar[1].namespace)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:89
		{
			yyVAL.expressions = append(yyVAL.expressions, yyDollar[1].typedef)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:93
		{
			yyVAL.expressions = append(yyVAL.expressions, yyDollar[1].Const)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:97
		{
			yyVAL.expressions = append(yyVAL.expressions, yyDollar[1].Struct)
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.go.y:103
		{
			yyVAL.include = yyDollar[1].include
			yyVAL.include.Path = yyDollar[2].string
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.go.y:108
		{
			yyVAL.include = yyDollar[1].include
			yyVAL.include.Name = yyDollar[2].ident
			yyVAL.include.Path = yyDollar[3].string
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.go.y:116
		{
			yyVAL.namespace = yyDollar[1].namespace
			yyVAL.namespace.Scope = yyDollar[2].ident
			yyVAL.namespace.Name = yyDollar[3].ident
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.go.y:124
		{
			yyVAL.typedef = yyDollar[1].typedef
			yyVAL.typedef.DefinitionType = yyDollar[2].ident
			yyVAL.typedef.Identifier = yyDollar[3].ident
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.go.y:132
		{
			yyVAL.Const = yyDollar[1].Const
			yyVAL.Const.Type = yyDollar[2].ident
			yyVAL.Const.Name = yyDollar[3].ident
			yyVAL.Const.Equal = yyDollar[4].equal
			yyVAL.Const.Value = yyDollar[5].value
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.go.y:142
		{
			yyVAL.Struct = yyDollar[1].Struct
			yyVAL.Struct.Name = yyDollar[2].ident
			yyVAL.Struct.LBrace = yyDollar[3].brace
			yyVAL.Struct.Fields = yyDollar[4].fields
			yyVAL.Struct.RBrace = yyDollar[5].brace
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:152
		{
			yyVAL.fields = append(yyVAL.fields, yyDollar[1].field)
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.go.y:158
		{
			yyVAL.req = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:162
		{
			yyVAL.req = yyDollar[1].req
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:166
		{
			yyVAL.req = yyDollar[1].req
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.go.y:172
		{
			yyVAL.field = ast.NewField()
			yyVAL.field.ID = yyDollar[1].int
			yyVAL.field.Colon = yyDollar[2].colon
			__yyfmt__.Println(yyDollar[3].req)
			yyVAL.field.Req = yyDollar[3].req
			yyVAL.field.Type = yyDollar[4].ident
			yyVAL.field.Name = yyDollar[5].ident
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.go.y:182
		{
			yyVAL.field = ast.NewField()
			yyVAL.field.ID = yyDollar[1].int
			yyVAL.field.Colon = yyDollar[2].colon
			yyVAL.field.Req = yyDollar[3].req
			yyVAL.field.Type = yyDollar[4].ident
			yyVAL.field.Name = yyDollar[5].ident
			yyVAL.field.Equal = yyDollar[6].equal
			yyVAL.field.DefaultValue = yyDollar[7].value
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:195
		{
			yyVAL.value = yyDollar[1].string
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:199
		{
			yyVAL.value = yyDollar[1].int
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:203
		{
			yyVAL.value = yyDollar[1].float
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.go.y:209
		{
			yyVAL.ident = yyDollar[1].ident
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.go.y:213
		{
			yyVAL.ident = yyDollar[1].ident.Append(yyDollar[2].dot)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.go.y:217
		{
			yyVAL.ident = yyDollar[1].ident.Append(yyDollar[2].ident)
		}
	}
	goto yystack /* stack new state and value */
}
