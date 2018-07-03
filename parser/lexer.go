package parser

import (
	"fmt"
	"text/scanner"

	"github.com/minodisk/thriftast/ast"
)

type Lexer struct {
	scanner.Scanner
	Program *ast.Program
}

func (l *Lexer) Lex(lval *yySymType) int {
	t := int(l.Scan())
	fmt.Println("-----------------------------")
	fmt.Printf("%s, %d %d %d %d %d\n", l.Pos(), t, scanner.Ident, scanner.Int, scanner.String, scanner.Char)

	name := l.TokenText()
	length := len(name)
	start := l.Pos()
	end := scanner.Position{
		Filename: start.Filename,
		Offset:   start.Offset + length,
		Line:     start.Line,
		Column:   start.Column + length,
	}

	switch t {
	case scanner.Ident:
		switch name {
		case "namespace":
			return NAMESPACE
		default:
			lval.token = &ast.Identifier{Start: start, End: end, Name: name}
			return IDENTIFIER
		}
	default:
		if name == "" {
			return t
		}
		lval.token = &ast.Identifier{Start: start, End: end, Name: name}
		return IDENTIFIER
	}
}

func (l *Lexer) Error(e string) {
	panic(e)
}
