package internal

const (
	RowNumBits = 40
	ColNumBits = 10
	FullRow    = (1 << ColNumBits) - 1
)

type Board struct {
	Rows [RowNumBits]uint16
	Cols [ColNumBits]Bitboard
}

func NewBoard() Board {
	return Board{}
}

func (b *Board) Occupied(x, y int32) bool {
	if y < 0 || y >= RowNumBits {
		return false
	}

	return b.Rows[y]&(1<<x) != 0
}

func (b *Board) OccupiedCoord(c Coordinate) bool {
	return b.Occupied(int32(c.X), int32(c.Y))
}

func (b *Board) Obstructed(x, y int32) bool {
	return !IsOkX(x) || !IsOkY(y) || b.Occupied(x, y)
}

func (b *Board) ObstructedCoord(c Coordinate) bool {
	return b.Obstructed(int32(c.X), int32(c.Y))
}

func (b *Board) ObstructedMove(m Move) bool {
	pc := m.Cells()
	off := NewCoordinate(m.X(), m.Y())

	return b.ObstructedCoord(off) ||
		b.ObstructedCoord(pc[0].Add(off.X, off.Y)) ||
		b.ObstructedCoord(pc[1].Add(off.X, off.Y)) ||
		b.ObstructedCoord(pc[2].Add(off.X, off.Y))
}

func (b *Board) LegalLockPlacement(m Move) bool {
	if !m.IsOkMove() || b.ObstructedMove(m) {
		return false
	}

	pc := m.Cells()
	off := NewCoordinate(m.X(), m.Y())

	cells := [4]Coordinate{
		off,
		pc[0].Add(off.X, off.Y),
		pc[1].Add(off.X, off.Y),
		pc[2].Add(off.X, off.Y),
	}

	for _, c := range cells {
		if b.Obstructed(int32(c.X), int32(c.Y)-1) {
			return true
		}
	}

	return false
}

func (b *Board) Col(x int) Bitboard {
	mask := uint16(1) << x
	var result Bitboard

	for y := 0; y < RowNumBits; y++ {
		if b.Rows[y]&mask != 0 {
			result |= Bitboard(1) << y
		}
	}

	return result
}
