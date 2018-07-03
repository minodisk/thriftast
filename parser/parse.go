package parser

import (
	"io"

	"github.com/minodisk/thriftast/ast"
)

func Parse(r io.Reader) *ast.Program {
	l := new(Lexer)
	l.Init(r)
	yyParse(l)
	return l.Program
}
