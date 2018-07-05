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
	ident       *ast.Ident
	dot         *ast.Dot
}
