package ast

import "encoding/json"

type Dot struct {
	start *Pos
	end   *Pos
}

func NewDot(start, end *Pos) *Dot {
	return &Dot{
		start,
		end,
	}
}

func (d *Dot) Start() *Pos {
	return d.start
}

func (d *Dot) End() *Pos {
	return d.end
}

func (d *Dot) Name() string {
	return "."
}

func (d *Dot) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExpType string `json:"__type__"`
		Start   *Pos   `json:"start"`
		End     *Pos   `json:"end"`
		Name    string `json:"name"`
	}{
		"Dot",
		d.Start(),
		d.End(),
		d.Name(),
	}
	return json.Marshal(typed)
}
