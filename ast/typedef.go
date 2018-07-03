package ast

import "encoding/json"

type Typedef struct {
	DefinitionType *Identifier
	Identifier     *Identifier
}

func (t *Typedef) Type() string {
	return "Typedef"
}

func (t *Typedef) MarshalJSON() ([]byte, error) {
	typed := struct {
		Typedef
		Type string
	}{
		*t,
		t.Type(),
	}
	return json.Marshal(typed)
}
