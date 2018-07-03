%{
package parser

import "github.com/minodisk/thriftast/ast"

%}

%union{
    program     *ast.Program
    expressions []ast.Expression
    namespace   *ast.Namespace

    token     *ast.Identifier
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

namespace
    : NAMESPACE IDENTIFIER IDENTIFIER
        {
            $$ = &ast.Namespace{
                Scope: $2,
                Name: $3,
            }
        }

%%
