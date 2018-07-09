package ast

import "encoding/json"

type Colon struct {
	start *Pos
	end   *Pos
}

func NewColon(start, end *Pos) *Colon {
	return &Colon{
		start,
		end,
	}
}

func (d *Colon) Start() *Pos {
	return d.start
}

func (d *Colon) End() *Pos {
	return d.end
}

func (d *Colon) Name() string {
	return ":"
}

func (d *Colon) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExpType string `json:"__type__"`
		Start   *Pos   `json:"start"`
		End     *Pos   `json:"end"`
		Name    string `json:"name"`
	}{
		"Colon",
		d.Start(),
		d.End(),
		d.Name(),
	}
	return json.Marshal(typed)
}
