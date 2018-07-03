package serializer

import (
	"fmt"
	"text/scanner"
)

type Ranger interface {
	Start() scanner.Position
	End() scanner.Position
}

func Range(r Ranger) string {
	return fmt.Sprintf("%s-%s", Pos(r.Start()), Pos(r.End()))
}

func Pos(p scanner.Position) string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}
