package parser_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/minodisk/thriftast/ast"
	"github.com/minodisk/thriftast/parser"
)

func TestParse_Typedef(t *testing.T) {
	for _, c := range []struct {
		in  string
		out *ast.Program
	}{
		{
			`typedef i32 MyInteger`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Typedef{
						Keyword: ast.NewKeyword(
							"Typedef",
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
						DefinitionType: ast.NewIdent(
							&ast.Pos{
								Offset: 8,
								Line:   1,
								Column: 9,
							},
							&ast.Pos{
								Offset: 11,
								Line:   1,
								Column: 12,
							},
							"i32",
						),
						Identifier: ast.NewIdent(
							&ast.Pos{
								Offset: 12,
								Line:   1,
								Column: 13,
							},
							&ast.Pos{
								Offset: 21,
								Line:   1,
								Column: 22,
							},
							"MyInteger",
						),
					},
				},
			},
		},
		{
			`typedef Tweet ReTweet`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Typedef{
						Keyword: ast.NewKeyword(
							"Typedef",
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
						DefinitionType: ast.NewIdent(
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
							"Tweet",
						),
						Identifier: ast.NewIdent(
							&ast.Pos{
								Offset: 14,
								Line:   1,
								Column: 15,
							},
							&ast.Pos{
								Offset: 21,
								Line:   1,
								Column: 22,
							},
							"ReTweet",
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
