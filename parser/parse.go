package parser

import (
	"io"

	"github.com/minodisk/thriftast/ast"
)

// Parse reads r and parses content to AST.
func Parse(r io.Reader) *ast.Program {
	l := new(Lexer)
	l.Init(r)
	yyParse(l)
	return l.Program
}
