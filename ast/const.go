package ast

import (
	"encoding/json"
)

type Const struct {
	Keyword *Keyword `json:"keyword"`
	Type    *Ident   `json:"type"`
	Name    *Ident   `json:"name"`
	Equal   *Equal   `json:"equal"`
	Value   Value    `json:"value"`
}

func NewConst(start, end *Pos) *Const {
	return &Const{
		Keyword: NewKeyword("Const", start, end),
	}
}

func (c *Const) Start() *Pos {
	return c.Keyword.Start
}

func (c *Const) End() *Pos {
	return c.Value.End()
}

func (n *Const) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExpType string `json:"__type__"`
		Const
	}{
		"Const",
		*n,
	}
	return json.Marshal(typed)
}
