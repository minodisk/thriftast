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

func (i *Include) Start() *Pos {
	return i.Keyword.Start
}

func (i *Include) End() *Pos {
	return i.Path.end
}

func (n *Include) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExtType string `json:"__type__"`
		Include
	}{
		"Include",
		*n,
	}
	return json.Marshal(typed)
}
