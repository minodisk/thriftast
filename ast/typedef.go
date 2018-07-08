package ast

import (
	"encoding/json"
)

type Typedef struct {
	Keyword        *Keyword `json:"keyword"`
	DefinitionType *Ident   `json:"definition_type"`
	Identifier     *Ident   `json:"identifier"`
}

func NewTypedef(start, end *Pos) *Typedef {
	return &Typedef{
		Keyword: NewKeyword("Typedef", start, end),
	}
}

func (t *Typedef) Start() *Pos {
	return t.Keyword.Start
}

func (t *Typedef) End() *Pos {
	return t.Identifier.end
}

func (t *Typedef) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExtType string `json:"type"`
		Typedef
	}{
		"Typedef",
		*t,
	}
	return json.Marshal(typed)
}
