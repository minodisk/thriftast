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
  required    *ast.Req
  optional    *ast.Req
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
%type <required>    required
%type <optional>    optional
%type <ident>       ident
%type <value>       value
%type <string>      string
%type <int>         int
%type <float>       float
%type <dot>         dot
%type <brace>       brace

// Keywords
%token <include>    INCLUDE
%token <namespace>  NAMESPACE
%token <typedef>    TYPEDEF
%token <Const>      CONST
%token <Struct>     STRUCT
%token <required>   REQUIRED
%token <optional>   OPTIONAL
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
%token <lparen>     LPAREN
%token <rparen>     RPAREN

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
  | include
    {
      $$ = append($$, $1)
    }
  | namespace
    {
      $$ = append($$, $1)
    }
  | typedef
    {
      $$ = append($$, $1)
    }
  | Const
    {
      $$ = append($$, $1)
    }
  | Struct
    {
      $$ = append($$, $1)
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

field
  : int COLON req ident ident
    {
      $$ = ast.NewField()
      $$.ID = $1
      $$.Colon = $2
      $$.Req = $3
      $$.Type = $4
      $$.Name = $5
    }
  | int COLON optional ident ident EQUAL value
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

req
  :
    {
      $$ = nil
    }
  | required
    {
      $$ = $1
    }
  | optional
    {
      $$ = $1
    }

required
  : REQUIRED
    {
      $$ = $1
    }

optional
  : OPTIONAL
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
      $$ = $1
    }

brace
  : LBRACE
    {
      $$ = $1
    }
  | RBRACE
    {
      $$ = $1
    }

%%
