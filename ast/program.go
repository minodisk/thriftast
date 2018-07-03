package ast

import "encoding/json"

type Program struct {
	Expressions []Expression
}

func (p *Program) Type() string {
	return "Program"
}

func (p *Program) MarshalJSON() ([]byte, error) {
	typed := struct {
		Program
		Type string
	}{
		*p,
		p.Type(),
	}
	return json.Marshal(typed)
}
