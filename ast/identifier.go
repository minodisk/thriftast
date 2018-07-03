package ast

import (
	"encoding/json"
	"text/scanner"
)

type Identifier struct {
	Start scanner.Position
	End   scanner.Position
	Name  string
}

func (i *Identifier) Type() string {
	return "Identifier"
}

func (i *Identifier) MarshalJSON() ([]byte, error) {
	typed := struct {
		Identifier
		Type string
	}{
		*i,
		i.Type(),
	}
	return json.Marshal(typed)
}
