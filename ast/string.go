package ast

import (
	"encoding/json"
)

type String struct {
	Keyword *Keyword `json:"keyword"`
	Name    *Ident   `json:"name"`
}

func NewString(start, end *Pos) *String {
	return &String{
		Keyword: NewKeyword("String", start, end),
	}
}

func (n *String) Type() string {
	return "String"
}

func (n *String) MarshalJSON() ([]byte, error) {
	typed := struct {
		Type string `json:"type"`
		String
	}{
		n.Type(),
		*n,
	}
	return json.Marshal(typed)
}
