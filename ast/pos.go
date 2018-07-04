package ast

import (
	"fmt"
	"text/scanner"
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func NewPos(p scanner.Position) *Pos {
	return &Pos{
		p.Offset,
		p.Line,
		p.Column,
	}
}

func (p *Pos) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

func (p *Pos) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", p)), nil
}
