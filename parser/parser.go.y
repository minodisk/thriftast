%{
package parser

import "github.com/minodisk/thriftast/ast"

%}

%union{
  program     *ast.Program
  expressions []ast.Expression
  namespace   *ast.Namespace
  typedef     *ast.Typedef

  identifier  *ast.Identifier
}

// Types
%type <program>     program
%type <expressions> expressions
%type <namespace>   namespace
%type <typedef>     typedef

// Reserved keywords
%token NAMESPACE
%token TYPEDEF

// Tokens
%token <identifier> IDENTIFIER

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
  : NAMESPACE IDENTIFIER IDENTIFIER
    {
      $$ = &ast.Namespace{
        Scope: $2,
        Name: $3,
      }
    }

typedef
  : TYPEDEF IDENTIFIER IDENTIFIER
    {
      $$ = &ast.Typedef{
        DefinitionType: $2,
        Identifier: $3,
      }
    }

%%
