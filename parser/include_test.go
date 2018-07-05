package parser_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/minodisk/thriftast/ast"
	"github.com/minodisk/thriftast/parser"
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
								Offset: 8,
								Line:   1,
								Column: 9,
							},
						),
						Path: ast.NewString(
							&ast.Pos{
								Offset: 9,
								Line:   1,
								Column: 10,
							},
							&ast.Pos{
								Offset: 23,
								Line:   1,
								Column: 24,
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
								Offset: 8,
								Line:   1,
								Column: 9,
							},
						),
						Name: ast.NewIdent(
							&ast.Pos{
								Offset: 9,
								Line:   1,
								Column: 10,
							},
							&ast.Pos{
								Offset: 14,
								Line:   1,
								Column: 15,
							},
							`tweet`,
						),
						Path: ast.NewString(
							&ast.Pos{
								Offset: 15,
								Line:   1,
								Column: 16,
							},
							&ast.Pos{
								Offset: 29,
								Line:   1,
								Column: 30,
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
			got := parser.Parse(bytes.NewBuffer([]byte(c.in)))
			if !reflect.DeepEqual(*got, *c.out) {
				g, _ := json.Marshal(got)
				w, _ := json.Marshal(c.out)
				t.Errorf("\ngot:\n%s\nwant:\n%s", g, w)
			}
		})
	}
}
