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
						Scope: &ast.Identifier{
							Start: scanner.Position{
								Filename: "",
								Offset:   10,
								Line:     1,
								Column:   11,
							},
							End: scanner.Position{
								Filename: "",
								Offset:   13,
								Line:     1,
								Column:   14,
							},
							Name: "cpp",
						},
						Name: &ast.Identifier{
							Start: scanner.Position{
								Filename: "",
								Offset:   14,
								Line:     1,
								Column:   15,
							},
							End: scanner.Position{
								Filename: "",
								Offset:   34,
								Line:     1,
								Column:   35,
							},
							Name: "com.example.project",
						},
					},
				},
			},
		},
		{
			`namespace * com.example.project`,
			&ast.Program{
				Expressions: []ast.Expression{
					&ast.Namespace{
						Scope: &ast.Identifier{
							Start: scanner.Position{
								Filename: "",
								Offset:   10,
								Line:     1,
								Column:   11,
							},
							End: scanner.Position{
								Filename: "",
								Offset:   11,
								Line:     1,
								Column:   12,
							},
							Name: "*",
						},
						Name: &ast.Identifier{
							Start: scanner.Position{
								Filename: "",
								Offset:   12,
								Line:     1,
								Column:   13,
							},
							End: scanner.Position{
								Filename: "",
								Offset:   31,
								Line:     1,
								Column:   32,
							},
							Name: "com.example.project",
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
