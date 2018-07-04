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

func (n *Namespace) Type() string {
	return "Namespace"
}

func (n *Namespace) MarshalJSON() ([]byte, error) {
	typed := struct {
		Type string `json:"type"`
		Namespace
	}{
		n.Type(),
		*n,
	}
	return json.Marshal(typed)
}
