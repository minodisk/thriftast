%{
package parser

import "github.com/minodisk/thriftast/ast"

%}

%union{
  program     *ast.Program
  expressions []ast.Expression
  namespace   *ast.Namespace
  typedef     *ast.Typedef
  ident       *ast.Ident
  dot         *ast.Dot
}

// Types
%type <program>     program
%type <expressions> expressions
%type <namespace>   namespace
%type <typedef>     typedef
%type <ident>       ident
%type <dot>         dot

// Tokens
%token <namespace>  NAMESPACE
%token <typedef>    TYPEDEF
%token <ident>      IDENT
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
  | expressions namespace
    {
      $$ = append($1, $2)
    }
  | expressions typedef
    {
      $$ = append($1, $2)
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

dot
  : DOT
    {
      $$ =$1
    }

%%
