package ast

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

func (d *Brace) Type() string {
	return "Brace"
}

func (d *Brace) Start() *Pos {
	return d.start
}

func (d *Brace) End() *Pos {
	return d.end
}

func (d *Brace) Name() string {
	return "."
}
