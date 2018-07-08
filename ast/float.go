package ast

import (
	"encoding/json"
)

type Float struct {
	start *Pos
	end   *Pos
	name  string
}

func NewFloat(start, end *Pos, name string) *Float {
	return &Float{
		start,
		end,
		name,
	}
}

func (s *Float) Type() string {
	return "Float"
}

func (s *Float) Start() *Pos {
	return s.start
}

func (s *Float) End() *Pos {
	return s.end
}

func (s *Float) Name() string {
	return s.name
}

func (s *Float) MarshalJSON() ([]byte, error) {
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
