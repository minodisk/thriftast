package parser_test

import (
	"testing"

	"github.com/minodisk/thriftast/ast"
	"github.com/minodisk/thriftast/tests"
)

func TestParse_Include(t *testing.T) {
	for _, c := range []struct {
		in  string
		out *ast.Program
	}{
		{
			`include "tweet.thrift"`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Include{
						Keyword: ast.NewKeyword(
							"Include",
							&ast.Pos{
								Offset: 0,
								Line:   1,
								Column: 1,
							},
							&ast.Pos{
								Offset: 7,
								Line:   1,
								Column: 8,
							},
						),
						Path: ast.NewString(
							&ast.Pos{
								Offset: 8,
								Line:   1,
								Column: 9,
							}, &ast.Pos{
								Offset: 22,
								Line:   1,
								Column: 23,
							},
							`"tweet.thrift"`,
						),
					},
				},
			},
		},
		{
			`include tweet "tweet.thrift"`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Include{
						Keyword: ast.NewKeyword(
							"Include",
							&ast.Pos{
								Offset: 0,
								Line:   1,
								Column: 1,
							},
							&ast.Pos{
								Offset: 7,
								Line:   1,
								Column: 8,
							},
						),
						Name: ast.NewIdent(
							&ast.Pos{
								Offset: 8,
								Line:   1,
								Column: 9,
							},
							&ast.Pos{
								Offset: 13,
								Line:   1,
								Column: 14,
							},
							`tweet`,
						),
						Path: ast.NewString(
							&ast.Pos{
								Offset: 14,
								Line:   1,
								Column: 15,
							}, &ast.Pos{
								Offset: 28,
								Line:   1,
								Column: 29,
							},
							`"tweet.thrift"`,
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
