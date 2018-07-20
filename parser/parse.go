package parser

import (
	"io"

	"github.com/minodisk/thriftast/ast"
)

// Parse reads r and parses content to AST.
func Parse(r io.Reader) *ast.Program {
	l := new(Lexer)
	l.Init(r)
	l.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '
	yyParse(l)
	return l.Program
}
