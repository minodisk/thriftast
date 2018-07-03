package parser

import (
	"fmt"
	"text/scanner"

	"github.com/minodisk/thriftast/ast"
)

// Lexer scans Thrift IDL to tokens.
type Lexer struct {
	scanner.Scanner
	Program *ast.Program
}

// Lex lexes chunk to token and returns the type of token.
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
		case "typedef":
			return TYPEDEF
		default:
			lval.identifier = &ast.Identifier{Start: start, End: end, Name: name}
			return IDENTIFIER
		}
	default:
		if name == "" {
			return t
		}
		lval.identifier = &ast.Identifier{Start: start, End: end, Name: name}
		return IDENTIFIER
	}
}

// Error throws lexing error.
func (l *Lexer) Error(e string) {
	panic(e)
}
