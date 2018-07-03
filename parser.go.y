%{
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/scanner"
)

type Expression interface {
    Type() string
}

type Program struct {
    Expressions []Expression
}

func (p *Program) Type() string {
    return "Program"
}

func (p *Program) MarshalJSON() ([]byte, error) {
    typed := struct {
        Program
        Type string
    }{
        *p,
        p.Type(),
    }
    return json.Marshal(typed)
}

type Namespace struct {
    Scope *Identifier
    Name  *Identifier
}

func (n *Namespace) Type() string {
    return "Namespace"
}

func (n *Namespace) MarshalJSON() ([]byte, error) {
    typed := struct {
        Namespace
        Type string
    }{
        *n,
        n.Type(),
    }
    return json.Marshal(typed)
}

type Identifier struct {
    Start scanner.Position
    End   scanner.Position
    Name  string
}

func (i *Identifier) Type() string {
    return "Identifier"
}

func (i *Identifier) MarshalJSON() ([]byte, error) {
    typed := struct {
        Identifier
        Type string
    }{
        *i,
        i.Type(),
    }
    return json.Marshal(typed)
}
%}

%union{
    program     *Program
    expressions []Expression
    namespace   *Namespace

    token     *Identifier
}

// Types
%type <program>     program
%type <expressions> expressions
%type <namespace>   namespace

// Reserved keywords
%token NAMESPACE

// Tokens
%token <token> IDENTIFIER

%%

program
    : expressions
    {
        $$ = &Program{Expressions: $1}
        yylex.(*Lexer).Program = $$
    }

expressions
    : /* no expressions */
        {
            $$ = nil
        }
    | expressions namespace
        {
            $$ = append($1, $2)
        }

namespace
    : NAMESPACE IDENTIFIER IDENTIFIER
        {
            $$ = &Namespace{
                Scope: $2,
                Name: $3,
            }
        }

%%

type Lexer struct {
    scanner.Scanner
    Program *Program
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
        Offset: start.Offset + length,
        Line: start.Line,
        Column: start.Column + length,
    }

    switch t {
    case scanner.Ident:
      switch name {
      case "namespace":
        return NAMESPACE
      default:
        lval.token = &Identifier{Start: start, End: end, Name: name}
        return IDENTIFIER
      }
    default:
      if name == "" {
        return t
      }
      lval.token = &Identifier{Start: start, End: end, Name: name}
      return IDENTIFIER
    }
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
    fmt.Println("========================")
    j, _ := json.MarshalIndent(l.Program, "", "  ")
    fmt.Printf("%s\n", j)
}
