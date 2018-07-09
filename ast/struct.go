package ast

import (
	"encoding/json"
)

type Struct struct {
	Keyword *Keyword `json:"keyword"`
	Name    *Ident   `json:"name"`
	LBrace  *Brace   `json:"lbrace"`
	Fields  []*Field `json:"fields"`
	RBrace  *Brace   `json:"rbrace"`
}

func NewStruct(start, end *Pos) *Struct {
	return &Struct{
		Keyword: NewKeyword("Struct", start, end),
	}
}

func (n *Struct) Start() *Pos {
	return n.Keyword.Start
}

func (n *Struct) End() *Pos {
	return n.Name.end
}

func (n *Struct) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExtType string `json:"type"`
		Struct
	}{
		"Struct",
		*n,
	}
	return json.Marshal(typed)
}
