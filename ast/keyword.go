package ast

import "text/scanner"

type Dot struct {
	start scanner.Position
	end   scanner.Position
}

func NewDot(start, end scanner.Position) *Dot {
	return &Dot{
		start,
		end,
	}
}

func (d *Dot) Type() string {
	return "Dot"
}

func (d *Dot) Start() scanner.Position {
	return d.start
}

func (d *Dot) End() scanner.Position {
	return d.end
}

func (d *Dot) Name() string {
	return "."
}
