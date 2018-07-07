package ast

import (
	"encoding/json"
)

type String struct {
	start *Pos
	end   *Pos
	name  string
}

func NewString(start, end *Pos, name string) *String {
	return &String{
		start,
		end,
		name,
	}
}

func (s *String) Type() string {
	return "String"
}

func (s *String) Start() *Pos {
	return s.start
}

func (s *String) End() *Pos {
	return s.end
}

func (s *String) Name() string {
	return s.name
}

func (s *String) MarshalJSON() ([]byte, error) {
	typed := struct {
		Type  string `json:"type"`
		Start *Pos   `json:"start"`
		End   *Pos   `json:"end"`
		Name  string `json:"name"`
	}{
		s.Type(),
		s.Start(),
		s.End(),
		s.Name(),
	}
	return json.Marshal(typed)
}
