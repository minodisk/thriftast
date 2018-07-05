package ast

import (
	"encoding/json"
)

type Include struct {
	Keyword *Keyword `json:"keyword"`
	Name    *Ident   `json:"name"`
	Path    *String  `json:"path"`
}

func NewInclude(start, end *Pos) *Include {
	return &Include{
		Keyword: NewKeyword("Include", start, end),
	}
}

func (n *Include) Type() string {
	return "Include"
}

func (n *Include) MarshalJSON() ([]byte, error) {
	typed := struct {
		Type string `json:"type"`
		Include
	}{
		n.Type(),
		*n,
	}
	return json.Marshal(typed)
}
