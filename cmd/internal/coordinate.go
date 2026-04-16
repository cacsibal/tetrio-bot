package internal

import "fmt"

type Coordinate struct {
	X, Y int8
}

func NewCoordinate(x, y int32) Coordinate {
	return Coordinate{X: int8(x), Y: int8(y)}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c Coordinate) Add(x, y int8) Coordinate {
	return Coordinate{c.X + x, c.Y + y}
}

func (c Coordinate) Sub(x, y int8) Coordinate {
	return Coordinate{c.X - x, c.Y - y}
}

func (c Coordinate) Rotate(r Rotation) Coordinate {
	switch r {
	case East:
		return Coordinate{X: c.Y, Y: -c.X}
	case South:
		return Coordinate{X: -c.X, Y: -c.Y}
	case West:
		return Coordinate{X: -c.Y, Y: c.X}
	case North:
		return c
	default:
		return c
	}
}

func IsOkX(x int32) bool {
	return x >= 0 && x < int32(ColNumBits)
}

func IsOkY(y int32) bool {
	return y >= 0 && y < int32(RowNumBits)
}
