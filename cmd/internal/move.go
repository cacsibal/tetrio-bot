package internal

const TSPIN uint16 = 7

// Move data is packed into a single 16-bit integer (uint16) to save memory.
//
//	15    14 13    12 11 10     9 8 7 6      5 4 3 2 1 0
//
// +----+---------+------------+----------+-----------------+
// | S  |    R    |     P      |    X     |        Y        |
// +----+---------+------------+----------+-----------------+
//
// Field breakdown:
// - S (1 bit,  bit 15)    : Spin flag (marks if the move resulted in a spin).
// - R (2 bits, bits 13-14): Rotation state (0-3, representing the four directions).
// - P (3 bits, bits 10-12): Piece type (0-7, covers 7 Tetrominoes + 1 sentinel).
// - X (4 bits, bits 6-9)  : Column position (0-15, easily covers a 10-column board).
// - Y (6 bits, bits 0-5)  : Row position (0-63, perfectly covers a 40-row board).
//
// Total: 1 + 2 + 3 + 4 + 6 = 16 bits.
type Move struct {
	data uint16
}

func NewMove(p Piece, r Rotation, x, y int, fullspin bool) Move {
	var pieceVal uint16
	if fullspin {
		pieceVal = TSPIN
	} else {
		pieceVal = uint16(p)
	}

	var spinBit uint16
	if fullspin {
		spinBit = 1
	}

	data := (uint16(y) & 0x3F) |
		((uint16(x) & 0xF) << 6) |
		((pieceVal & 0x7) << 10) |
		((uint16(r) & 0x3) << 13) |
		(spinBit << 15)

	return Move{data: data}
}

func NewTSpinMove(r Rotation, x, y int, fullspin bool) Move {
	var spinBit uint16
	if fullspin {
		spinBit = 1
	}

	data := (uint16(y) & 0x3F) |
		((uint16(x) & 0xF) << 6) |
		((TSPIN & 0x7) << 10) |
		((uint16(r) & 0x3) << 13) |
		(spinBit << 15)

	return Move{data: data}
}

func NoneMove() Move {
	return Move{data: 0}
}

func (m Move) X() int32 {
	return int32((m.data >> 6) & 0xF)
}

func (m Move) Y() int32 {
	return int32(m.data & 0x3F)
}

func (m Move) Piece() Piece {
	raw := (m.data >> 10) & 0x7
	if raw == TSPIN {
		return T
	}
	return Piece(raw)
}

func (m Move) Rotation() Rotation {
	return Rotation((m.data >> 13) & 0x3)
}

func (m Move) Spin() SpinType {
	pieceRaw := (m.data >> 10) & 0x7
	spinBit := (m.data >> 15) & 0x1

	var val uint8
	if pieceRaw == TSPIN {
		val = 1
	}
	val += uint8(spinBit)

	return SpinType(val)
}

func (m Move) Raw() uint16 {
	return m.data
}

func (m Move) Cells() PieceCoordinates {
	return PieceTable(m.Piece(), m.Rotation())
}

func (m Move) IsOkMove() bool {
	return IsOkX(m.X()) && IsOkY(m.Y())
}
