package ast

import (
	"encoding/json"
)

type Namespace struct {
	Keyword *Keyword `json:"keyword"`
	Scope   *Ident   `json:"scope"`
	Name    *Ident   `json:"name"`
}

func NewNamespace(start, end *Pos) *Namespace {
	return &Namespace{
		Keyword: NewKeyword("Namespace", start, end),
	}
}

func (n *Namespace) Start() *Pos {
	return n.Keyword.Start
}

func (n *Namespace) End() *Pos {
	return n.Name.end
}

func (n *Namespace) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExtType string `json:"type"`
		Namespace
	}{
		"Namespace",
		*n,
	}
	return json.Marshal(typed)
}
