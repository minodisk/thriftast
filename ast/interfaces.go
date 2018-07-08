package ast

type Expression interface {
	Start() *Pos
	End() *Pos
}

type Token interface {
	Expression
	Name() string
}

type Value interface {
	Start() *Pos
	End() *Pos
}
