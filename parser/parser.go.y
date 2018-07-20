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
  Struct      *ast.Struct
  fields      []*ast.Field
  field       *ast.Field
  req         *ast.Req
  ident       *ast.Ident
  value       ast.Value
  string      *ast.String
  int         *ast.Int
  float       *ast.Float
  equal       *ast.Equal
  dot         *ast.Dot
  colon       *ast.Colon
  brace       *ast.Brace
}

// Keywords
%token <include>    INCLUDE
%token <namespace>  NAMESPACE
%token <typedef>    TYPEDEF
%token <Const>      CONST
%token <Struct>     STRUCT
%token <req>        REQUIRED
%token <req>        OPTIONAL
%token <service>    SERVICE
%token <oneway>     ONEWAY

// Tokens
%token <ident>      IDENT
%token <string>     STRING
%token <int>        INT
%token <float>      FLOAT
%token <equal>      EQUAL
%token <dot>        DOT
%token <colon>      COLON
%token <brace>      LBRACE
%token <brace>      RBRACE
%token SEP

// Types
%type <program>     program
%type <expressions> expressions
%type <include>     include
%type <namespace>   namespace
%type <typedef>     typedef
%type <Const>       Const
%type <Struct>      Struct
%type <fields>      fields
%type <field>       field
%type <req>         req
%type <ident>       ident
%type <value>       value

%%

program
  : expressions
  {
    $$ = &ast.Program{Expressions: $1}
    yylex.(*Lexer).Program = $$
  }

expressions
  :
    {
      $$ = nil
    }
  | expressions include sep
    {
      $$ = append($1, $2)
    }
  | expressions namespace sep
    {
      $$ = append($1, $2)
    }
  | expressions typedef sep
    {
      $$ = append($1, $2)
    }
  | expressions Const sep
    {
      $$ = append($1, $2)
    }
  | expressions Struct sep
    {
      $$ = append($1, $2)
    }

sep
  :
  | SEP
  | sep sep

include
  : INCLUDE STRING
    {
      $$ = $1
      $$.Path = $2
    }
  | INCLUDE ident STRING
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

Struct
  : STRUCT IDENT LBRACE fields RBRACE
    {
      $$ = $1
      $$.Name = $2
      $$.LBrace = $3
      $$.Fields = $4
      $$.RBrace = $5
    }

fields
  : field
    {
      $$ = append($$, $1)
    }

req
  :
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

field
  : INT COLON req ident ident
    {
      $$ = ast.NewField()
      $$.ID = $1
      $$.Colon = $2
      $$.Req = $3
      $$.Type = $4
      $$.Name = $5
    }
  | INT COLON req ident ident EQUAL value
    {
      $$ = ast.NewField()
      $$.ID = $1
      $$.Colon = $2
      $$.Req = $3
      $$.Type = $4
      $$.Name = $5
      $$.Equal = $6
      $$.DefaultValue = $7
    }

value
  : STRING
    {
      $$ = $1
    }
  | INT
    {
      $$ = $1
    }
  | FLOAT
    {
      $$ = $1
    }

ident
  : IDENT
    {
      $$ = $1
    }
  | IDENT DOT
    {
      $$ = $1.Append($2)
    }
  | ident ident
    {
      $$ = $1.Append($2)
    }

%%
