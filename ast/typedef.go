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

func (t *Typedef) Type() string {
	return "Typedef"
}

func (t *Typedef) MarshalJSON() ([]byte, error) {
	typed := struct {
		Type string `json:"type"`
		Typedef
	}{
		t.Type(),
		*t,
	}
	return json.Marshal(typed)
}
