package ast

import "encoding/json"

type Program struct {
	Expressions []Expression `json:"expressions"`
}

func (p *Program) Type() string {
	return "Program"
}

func (p *Program) MarshalJSON() ([]byte, error) {
	typed := struct {
		Type string `json:"type"`
		Program
	}{
		p.Type(),
		*p,
	}
	return json.Marshal(typed)
}
