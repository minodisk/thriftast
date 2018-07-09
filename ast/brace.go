package ast

import "encoding/json"

type LR int

const (
	Left LR = iota
	Right
)

type Brace struct {
	start *Pos
	end   *Pos
	lr    LR
}

func NewBrace(start, end *Pos, lr LR) *Brace {
	return &Brace{
		start,
		end,
		lr,
	}
}

func NewLBrace(start, end *Pos) *Brace {
	return NewBrace(start, end, Left)
}

func NewRBrace(start, end *Pos) *Brace {
	return NewBrace(start, end, Right)
}

func (d *Brace) Start() *Pos {
	return d.start
}

func (d *Brace) End() *Pos {
	return d.end
}

func (d *Brace) Name() string {
	switch d.lr {
	default:
		return ""
	case Left:
		return "{"
	case Right:
		return "}"
	}
}

func (d *Brace) MarshalJSON() ([]byte, error) {
	typed := struct {
		ExpType string `json:"__type__"`
		Start   *Pos   `json:"start"`
		End     *Pos   `json:"end"`
		Name    string `json:"name"`
	}{
		"Brace",
		d.Start(),
		d.End(),
		d.Name(),
	}
	return json.Marshal(typed)
}
