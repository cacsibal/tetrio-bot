package internal

import "fmt"

type Coordinate struct {
	X, Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c Coordinate) Add(x, y int) Coordinate {
	return Coordinate{c.X + x, c.Y + y}
}

func (c Coordinate) Sub(x, y int) Coordinate {
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
