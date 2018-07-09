package ast

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

func (d *Colon) Type() string {
	return "Colon"
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
