package ast

type Keyword struct {
	Type  string `json:"type"`
	Start *Pos   `json:"start"`
	End   *Pos   `json:"end"`
}

func NewKeyword(t string, start, end *Pos) *Keyword {
	return &Keyword{
		t,
		start,
		end,
	}
}
