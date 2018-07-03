package ast

import (
	"encoding/json"
	"text/scanner"

	"github.com/minodisk/thriftast/serializer"
)

type Ident struct {
	start scanner.Position
	end   scanner.Position
	name  string
}

func NewIdent(start, end scanner.Position, name string) *Ident {
	return &Ident{
		start,
		end,
		name,
	}
}

func (i *Ident) Type() string {
	return "Ident"
}

func (i *Ident) Start() scanner.Position {
	return i.start
}

func (i *Ident) End() scanner.Position {
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
		Range string `json:"range"`
		Name  string `json:"name"`
	}{
		i.Type(),
		serializer.Range(i),
		i.Name(),
	}
	return json.Marshal(typed)
}
