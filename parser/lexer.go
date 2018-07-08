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

	token := l.TokenText()
	length := len(token)
	end := ast.NewPos(l.Pos())
	start := &ast.Pos{
		Offset: end.Offset - length,
		Line:   end.Line,
		Column: end.Column - length,
	}

	fmt.Println("-----------------------------")
	fmt.Println(t, token, length, scanner.String)
	fmt.Printf("%s ~ %s\n", start, end)

	switch t {
	case scanner.Ident:
		switch token {
		case "include":
			lval.include = ast.NewInclude(
				start,
				end,
			)
			return INCLUDE
		case "namespace":
			lval.namespace = ast.NewNamespace(
				start,
				end,
			)
			return NAMESPACE
		case "typedef":
			lval.typedef = ast.NewTypedef(start, end)
			return TYPEDEF
		case "const":
			lval.Const = ast.NewConst(start, end)
			return CONST
		default:
			lval.ident = ast.NewIdent(start, end, token)
			return IDENT
		}
	case scanner.String:
		lval.string = ast.NewString(start, end, token)
		return STRING
	case scanner.Int:
		lval.int = ast.NewInt(start, end, token)
		return INT
	case scanner.Float:
		lval.float = ast.NewFloat(start, end, token)
		return FLOAT
	case 46: // .
		lval.dot = ast.NewDot(start, end)
		return DOT
	case 61: // =
		lval.equal = ast.NewEqual(start, end)
		return EQUAL
	default:
		if token == "" {
			return t
		}
		lval.ident = ast.NewIdent(start, end, token)
		return IDENT
	}
}

// Error throws lexing error.
func (l *Lexer) Error(e string) {
	panic(e)
}
