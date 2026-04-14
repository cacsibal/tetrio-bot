package internal

type PieceCoordinates [3]Coordinate

func MakePiece(p Piece) PieceCoordinates {
	switch p {
	case I:
		return PieceCoordinates{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}
	case O:
		return PieceCoordinates{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}}
	case T:
		return PieceCoordinates{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}}
	case L:
		return PieceCoordinates{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}}
	case J:
		return PieceCoordinates{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: -1, Y: 1}}
	case S:
		return PieceCoordinates{{X: -1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}}
	case Z:
		return PieceCoordinates{{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 0}}
	default:
		return PieceCoordinates{}
	}
}

func PieceTable(p Piece, r Rotation) PieceCoordinates {
	cells := MakePiece(p)
	return PieceCoordinates{
		cells[0].Rotate(r),
		cells[1].Rotate(r),
		cells[2].Rotate(r),
	}
}
