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
