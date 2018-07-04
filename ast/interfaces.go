package ast

type Expression interface {
	Type() string
}

type Tokener interface {
	Type() string
	Start() *Pos
	End() *Pos
	Name() string
}
