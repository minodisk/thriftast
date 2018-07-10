package ast

import "encoding/json"

type RO int

const (
	Required RO = iota
	Optional
)

type Req struct {
	start *Pos
	end   *Pos
	ro    RO
}

func NewReq(start, end *Pos, ro RO) *Req {
	return &Req{
		start,
		end,
		ro,
	}
}

func NewRequired(start, end *Pos) *Req {
	return NewReq(start, end, Required)
}

func NewOptional(start, end *Pos) *Req {
	return NewReq(start, end, Optional)
}

func (d *Req) Start() *Pos {
	return d.start
}

func (d *Req) End() *Pos {
	return d.end
}

func (d *Req) Name() string {
	switch d.ro {
	default:
		return ""
	case Required:
		return "required"
	case Optional:
		return "optional"
	}
}

func (d *Req) MarshalJSON() ([]byte, error) {
	typed := struct {
		Type  string `json:"type"`
		Start *Pos   `json:"start"`
		End   *Pos   `json:"end"`
		Name  string `json:"name"`
	}{
		"Req",
		d.Start(),
		d.End(),
		d.Name(),
	}
	return json.Marshal(typed)
}
