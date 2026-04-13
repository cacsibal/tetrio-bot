package internal

type Piece int

const (
	I = iota
	O
	T
	L
	J
	S
	Z
)

func (p Piece) String() string {
	return []string{"I", "O", "T", "L", "J", "S", "Z"}[p]
}
