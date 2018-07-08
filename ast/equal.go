package ast

import "encoding/json"

type Equal struct {
	start *Pos
	end   *Pos
}

func NewEqual(start, end *Pos) *Equal {
	return &Equal{
		start,
		end,
	}
}

func (e *Equal) Start() *Pos {
	return e.start
}

func (e *Equal) End() *Pos {
	return e.end
}

func (e *Equal) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExpType string `json:"__type__"`
		Start   *Pos   `json:"start"`
		End     *Pos   `json:"end"`
	}{
		"Equal",
		e.Start(),
		e.End(),
	}
	return json.Marshal(typed)
}
