package ast

import (
	"encoding/json"
)

type Field struct {
	ID           *Int   `json:"id"`
	Colon        *Colon `json:"colon"`
	Req          *Req   `json:"req"`
	Type         *Ident `json:"type"`
	Name         *Ident `json:"name"`
	Equal        *Equal `json:"equal"`
	DefaultValue Value  `json:"default_value"`
}

func NewField() *Field {
	return &Field{}
}

func (c *Field) Start() *Pos {
	return c.ID.Start()
}

func (c *Field) End() *Pos {
	if c.DefaultValue != nil {
		return c.DefaultValue.End()
	}
	return c.Name.End()
}

func (n *Field) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExpType string `json:"__type__"`
		Field
	}{
		"Field",
		*n,
	}
	return json.Marshal(typed)
}
