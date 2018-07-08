package ast

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

func (d *Dot) Type() string {
	return "Dot"
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
