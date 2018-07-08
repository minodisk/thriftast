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
  Const       *ast.Const
  ident       *ast.Ident
  value       ast.Value
  string      *ast.String
  int         *ast.Int
  float       *ast.Float
  equal       *ast.Equal
  dot         *ast.Dot
}

// Types
%type <program>     program
%type <expressions> expressions
%type <include>     include
%type <namespace>   namespace
%type <typedef>     typedef
%type <Const>       Const
%type <ident>       ident
%type <value>       value
%type <string>      string
%type <int>         int
%type <float>       float
%type <dot>         dot

// Keywords
%token <include>    INCLUDE
%token <namespace>  NAMESPACE
%token <typedef>    TYPEDEF
%token <Const>      CONST
// Tokens
%token <ident>      IDENT
%token <string>     STRING
%token <int>        INT
%token <float>      FLOAT
%token <equal>      EQUAL
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
  | expressions Const
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

Const
  : CONST ident ident EQUAL value
    {
      $$ = $1
      $$.Type = $2
      $$.Name = $3
      $$.Equal = $4
      $$.Value = $5
    }

function
  : ident IDENT L_BRACKET fields R_BRACKET
    {
      $$ = $1
      $$.returnType = $2
      $$.lBracket = $3

    }

field
  : int ":" field_required type ident

field_required
  : /* not specified */
    {
      $$ = nil
    }
  | REQUIRED
    {
      $$ = $1
    }
  | OPTIONAL
    {
      $$ = $1
    }

value
  : string
    {
      $$ = $1
    }
  | int
    {
      $$ = $1
    }
  | float
    {
      $$ = $1
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

int
  : INT
    {
      $$ = $1
    }

float
  : FLOAT
    {
      $$ = $1
    }

dot
  : DOT
    {
      $$ =$1
    }

%%
