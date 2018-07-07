package tests

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/minodisk/thriftast/ast"
	"github.com/minodisk/thriftast/parser"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestDiff(t *testing.T, in string, want *ast.Program) {
	got := parser.Parse(bytes.NewBuffer([]byte(in)))

	if reflect.DeepEqual(*got, *want) {
		return
	}

	g, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	w, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(g), string(w), false)
	t.Error(dmp.DiffPrettyText(diffs))
}
