package parser_test

import (
	"testing"

	"github.com/minodisk/thriftast/ast"
	"github.com/minodisk/thriftast/tests"
)

func TestParse_Struct(t *testing.T) {
	for _, c := range []struct {
		in  string
		out *ast.Program
	}{
		{
			`struct Tweet {
  4: optional Location loc
}`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Struct{
						Keyword: ast.NewKeyword(
							"Struct",
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
							`loc`,
						),
						Fields: []*ast.Field{
							&ast.Field{
								ID: ast.NewInt(
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
									"4",
								),
								Colon: ast.NewColon(
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
								Req: ast.NewOptional(
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
									`Location`,
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
									`loc`,
								),
							},
						},
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
