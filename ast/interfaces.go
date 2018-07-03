package ast

import "text/scanner"

type Expression interface {
	Type() string
}

type Tokener interface {
	Type() string
	Start() scanner.Position
	End() scanner.Position
	Name() string
}
