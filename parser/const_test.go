package parser_test

import (
	"testing"

	"github.com/minodisk/thriftast/ast"
	"github.com/minodisk/thriftast/tests"
)

func TestParse_Const(t *testing.T) {
	for _, c := range []struct {
		in  string
		out *ast.Program
	}{
		{
			`const i32 INT_CONST = 1234`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Const{
						Keyword: ast.NewKeyword(
							"Const",
							&ast.Pos{
								Offset: 0,
								Line:   1,
								Column: 1,
							},
							&ast.Pos{
								Offset: 5,
								Line:   1,
								Column: 6,
							},
						),
						Type: ast.NewIdent(
							&ast.Pos{
								Offset: 6,
								Line:   1,
								Column: 7,
							}, &ast.Pos{
								Offset: 9,
								Line:   1,
								Column: 10,
							},
							`i32`,
						),
						Name: ast.NewIdent(
							&ast.Pos{
								Offset: 10,
								Line:   1,
								Column: 11,
							}, &ast.Pos{
								Offset: 19,
								Line:   1,
								Column: 20,
							},
							`INT_CONST`,
						),
						Equal: ast.NewEqual(
							&ast.Pos{
								Offset: 20,
								Line:   1,
								Column: 21,
							}, &ast.Pos{
								Offset: 21,
								Line:   1,
								Column: 22,
							},
						),
						Value: ast.NewInt(
							&ast.Pos{
								Offset: 22,
								Line:   1,
								Column: 23,
							}, &ast.Pos{
								Offset: 26,
								Line:   1,
								Column: 27,
							},
							`1234`,
						),
					},
				},
			},
		},
	} {
		c := c
		t.Run(c.in, func(t *testing.T) {
			tests.TestDiff(t, c.in, c.out)
		})
	}
}
