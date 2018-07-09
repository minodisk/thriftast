package ast

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

func (d *Req) Type() string {
	return "Req"
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
		return "Required"
	case Optional:
		return "Optional"
	}
}
