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
	required    *ast.Req
	optional    *ast.Req
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
const LPAREN = 57364
const RPAREN = 57365

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
	"LPAREN",
	"RPAREN",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser/parser.go.y:277
