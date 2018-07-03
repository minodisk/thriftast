package ast

import "encoding/json"

type Namespace struct {
	Scope *Identifier
	Name  *Identifier
}

func (n *Namespace) Type() string {
	return "Namespace"
}

func (n *Namespace) MarshalJSON() ([]byte, error) {
	typed := struct {
		Namespace
		Type string
	}{
		*n,
		n.Type(),
	}
	return json.Marshal(typed)
}
