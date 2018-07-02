%{
package main

import (
	"fmt"
	"os"
	"text/scanner"
)

type Expression interface{}

type Token struct {
    token   int
    literal string
}

type Namespace struct {
	Scope string
	Name  string
}
%}

%union{
    str       string

    token     Token
    expr      Expression

    namespace *Namespace
}

%token <str> IDENTIFIER

%token NAMESPACE

%type <namespace> namespace

%%

namespace
    : NAMESPACE '*' IDENTIFIER
        {
            $$ = &Namespace{
                Scope: "*",
                Name: $3,
            }
        }
    | NAMESPACE IDENTIFIER IDENTIFIER
        {
            $$ = &Namespace{
                Scope: $2,
                Name: $3,
            }
        }

%%

type Lexer struct {
    scanner.Scanner
    result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
    token := int(l.Scan())
    fmt.Println("-----------------------------")
    fmt.Printf("%+v, %s\n", l.Pos(), token, scanner.Ident, scanner.Int, scanner.String, NAMESPACE)
    fmt.Println("-----------------------------")

    literal := l.TokenText()

    switch token {
    case scanner.EOF:
    case scanner.Ident:
      switch literal {
      case "namespace":
        token = NAMESPACE
      }
    }

    lval.token = Token{token: token, literal: literal}

    return token
}

func (l *Lexer) Error(e string) {
    panic(e)
}

func main() {
	file, err := os.Open("test.thrift")
  if (err != nil) {
    panic(err)
  }

    l := new(Lexer)
    l.Init(file)
    yyParse(l)
    fmt.Printf("%#v\n", l.result)
}
