package ast

import (
	"encoding/json"
)

type Int struct {
	start *Pos
	end   *Pos
	name  string
}

func NewInt(start, end *Pos, name string) *Int {
	return &Int{
		start,
		end,
		name,
	}
}

func (s *Int) Type() string {
	return "Int"
}

func (s *Int) Start() *Pos {
	return s.start
}

func (s *Int) End() *Pos {
	return s.end
}

func (s *Int) Name() string {
	return s.name
}

func (s *Int) MarshalJSON() ([]byte, error) {
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
