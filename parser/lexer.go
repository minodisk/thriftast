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

	name := l.TokenText()
	length := len(name)
	end := l.Pos()
	start := scanner.Position{
		Filename: end.Filename,
		Offset:   end.Offset - length,
		Line:     end.Line,
		Column:   end.Column - length,
	}

	fmt.Println("-----------------------------")
	fmt.Println(name, length)
	fmt.Printf("%s ~ %s\n", start, end)

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
