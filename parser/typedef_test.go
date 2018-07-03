package parser_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"text/scanner"

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
						DefinitionType: &ast.Identifier{
							Start: scanner.Position{
								Filename: "",
								Offset:   8,
								Line:     1,
								Column:   9,
							},
							End: scanner.Position{
								Filename: "",
								Offset:   11,
								Line:     1,
								Column:   12,
							},
							Name: "i32",
						},
						Identifier: &ast.Identifier{
							Start: scanner.Position{
								Filename: "",
								Offset:   12,
								Line:     1,
								Column:   13,
							},
							End: scanner.Position{
								Filename: "",
								Offset:   21,
								Line:     1,
								Column:   22,
							},
							Name: "MyInteger",
						},
					},
				},
			},
		},
		{
			`typedef Tweet ReTweet`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Typedef{
						DefinitionType: &ast.Identifier{
							Start: scanner.Position{
								Filename: "",
								Offset:   8,
								Line:     1,
								Column:   9,
							},
							End: scanner.Position{
								Filename: "",
								Offset:   13,
								Line:     1,
								Column:   14,
							},
							Name: "Tweet",
						},
						Identifier: &ast.Identifier{
							Start: scanner.Position{
								Filename: "",
								Offset:   14,
								Line:     1,
								Column:   15,
							},
							End: scanner.Position{
								Filename: "",
								Offset:   21,
								Line:     1,
								Column:   22,
							},
							Name: "ReTweet",
						},
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
