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
								Offset: 6,
								Line:   1,
								Column: 7,
							},
						),
						Name: ast.NewIdent(
							&ast.Pos{
								Offset: 7,
								Line:   1,
								Column: 8,
							}, &ast.Pos{
								Offset: 12,
								Line:   1,
								Column: 13,
							},
							`Tweet`,
						),
						LBrace: ast.NewLBrace(
							&ast.Pos{
								Offset: 13,
								Line:   1,
								Column: 14,
							}, &ast.Pos{
								Offset: 14,
								Line:   1,
								Column: 15,
							},
						),
						Fields: []*ast.Field{
							&ast.Field{
								ID: ast.NewInt(
									&ast.Pos{
										Offset: 17,
										Line:   2,
										Column: 3,
									},
									&ast.Pos{
										Offset: 18,
										Line:   2,
										Column: 4,
									},
									"4",
								),
								Colon: ast.NewColon(
									&ast.Pos{
										Offset: 18,
										Line:   2,
										Column: 4,
									},
									&ast.Pos{
										Offset: 19,
										Line:   2,
										Column: 5,
									},
								),
								Req: ast.NewOptional(
									&ast.Pos{
										Offset: 20,
										Line:   2,
										Column: 6,
									},
									&ast.Pos{
										Offset: 28,
										Line:   2,
										Column: 14,
									},
								),
								Type: ast.NewIdent(
									&ast.Pos{
										Offset: 29,
										Line:   2,
										Column: 15,
									}, &ast.Pos{
										Offset: 37,
										Line:   2,
										Column: 23,
									},
									`Location`,
								),
								Name: ast.NewIdent(
									&ast.Pos{
										Offset: 38,
										Line:   2,
										Column: 24,
									}, &ast.Pos{
										Offset: 41,
										Line:   2,
										Column: 27,
									},
									`loc`,
								),
							},
						},
						RBrace: ast.NewRBrace(
							&ast.Pos{
								Offset: 42,
								Line:   3,
								Column: 1,
							}, &ast.Pos{
								Offset: 43,
								Line:   3,
								Column: 2,
							},
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
