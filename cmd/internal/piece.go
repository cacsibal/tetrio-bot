package internal

type Piece int

const (
	I Piece = iota
	O
	T
	L
	J
	S
	Z
)

var AllPieces = []Piece{I, O, T, L, J, S, Z}

func (p Piece) String() string {
	if int(p) < 0 || int(p) >= len(AllPieces) {
		return "Unknown"
	}

	return []string{"I", "O", "T", "L", "J", "S", "Z"}[p]
}
