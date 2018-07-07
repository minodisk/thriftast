%{
package parser

import "github.com/minodisk/thriftast/ast"

%}

%union{
  program     *ast.Program
  expressions []ast.Expression
  include     *ast.Include
  namespace   *ast.Namespace
  typedef     *ast.Typedef
  ident       *ast.Ident
  string      *ast.String
  dot         *ast.Dot
}

// Types
%type <program>     program
%type <expressions> expressions
%type <include>     include
%type <namespace>   namespace
%type <typedef>     typedef
%type <ident>       ident
%type <string>      string
%type <dot>         dot

// Keywords
%token <include>    INCLUDE
%token <namespace>  NAMESPACE
%token <typedef>    TYPEDEF
// Tokens
%token <ident>      IDENT
%token <string>     STRING
%token <dot>        DOT

%%

program
  : expressions
  {
    $$ = &ast.Program{Expressions: $1}
    yylex.(*Lexer).Program = $$
  }

expressions
  : /* no expressions */
    {
      $$ = nil
    }
  | expressions include
    {
      $$ = append($1, $2)
    }
  | expressions namespace
    {
      $$ = append($1, $2)
    }
  | expressions typedef
    {
      $$ = append($1, $2)
    }

include
  : INCLUDE string
    {
      $$ = $1
      $$.Path = $2
    }
  | INCLUDE ident string
    {
      $$ = $1
      $$.Name = $2
      $$.Path = $3
    }

namespace
  : NAMESPACE ident ident
    {
      $$ = $1
      $$.Scope = $2
      $$.Name = $3
    }

typedef
  : TYPEDEF ident ident
    {
      $$ = $1
      $$.DefinitionType = $2
      $$.Identifier = $3
    }

ident
  : IDENT
    {
      $$ = $1
    }
  | IDENT dot
    {
      $$ = $1.Append($2)
    }
  | ident ident
    {
      $$ = $1.Append($2)
    }

string
  : STRING
    {
      $$ = $1
    }

dot
  : DOT
    {
      $$ =$1
    }

%%
