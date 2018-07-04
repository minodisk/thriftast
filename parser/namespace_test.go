package parser_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/minodisk/thriftast/ast"
	"github.com/minodisk/thriftast/parser"
)

func TestParse_Namespace(t *testing.T) {
	for _, c := range []struct {
		in  string
		out *ast.Program
	}{
		{
			`namespace cpp com.example.project`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Namespace{
						Keyword: ast.NewKeyword(
							"Namespace",
							&ast.Pos{
								Offset: 0,
								Line:   1,
								Column: 1,
							},
							&ast.Pos{
								Offset: 9,
								Line:   1,
								Column: 10,
							},
						),
						Scope: ast.NewIdent(
							&ast.Pos{
								Offset: 10,
								Line:   1,
								Column: 11,
							},
							&ast.Pos{
								Offset: 13,
								Line:   1,
								Column: 14,
							},
							"cpp",
						),
						Name: ast.NewIdent(
							&ast.Pos{
								Offset: 14,
								Line:   1,
								Column: 15,
							},
							&ast.Pos{
								Offset: 33,
								Line:   1,
								Column: 34,
							},
							"com.example.project",
						),
					},
				},
			},
		},
		{
			`namespace * com.example.project`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Namespace{
						Keyword: ast.NewKeyword(
							"Namespace",
							&ast.Pos{
								Offset: 0,
								Line:   1,
								Column: 1,
							},
							&ast.Pos{
								Offset: 9,
								Line:   1,
								Column: 10,
							},
						),
						Scope: ast.NewIdent(
							&ast.Pos{
								Offset: 10,
								Line:   1,
								Column: 11,
							},
							&ast.Pos{
								Offset: 11,
								Line:   1,
								Column: 12,
							},
							"*",
						),
						Name: ast.NewIdent(
							&ast.Pos{
								Offset: 12,
								Line:   1,
								Column: 13,
							},
							&ast.Pos{
								Offset: 31,
								Line:   1,
								Column: 32,
							},
							"com.example.project",
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
