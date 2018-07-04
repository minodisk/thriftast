package ast

import (
	"encoding/json"
)

type Ident struct {
	start *Pos
	end   *Pos
	name  string
}

func NewIdent(start, end *Pos, name string) *Ident {
	return &Ident{
		start,
		end,
		name,
	}
}

func (i *Ident) Type() string {
	return "Ident"
}

func (i *Ident) Start() *Pos {
	return i.start
}

func (i *Ident) End() *Pos {
	return i.end
}

func (i *Ident) Name() string {
	return i.name
}

func (i *Ident) Append(j Tokener) *Ident {
	return &Ident{
		start: i.Start(),
		end:   j.End(),
		name:  i.Name() + j.Name(),
	}
}

func (i *Ident) MarshalJSON() ([]byte, error) {
	typed := struct {
		Type  string `json:"type"`
		Start *Pos   `json:"start"`
		End   *Pos   `json:"end"`
		Name  string `json:"name"`
	}{
		i.Type(),
		i.Start(),
		i.End(),
		i.Name(),
	}
	return json.Marshal(typed)
}
