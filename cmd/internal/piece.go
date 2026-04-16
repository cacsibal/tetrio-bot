package internal

type Piece int

const NumPieces = 7

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

func IsOkPiece(p Piece) bool {
	return uint8(p) < NumPieces
}
