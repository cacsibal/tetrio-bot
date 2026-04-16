package internal

type Rotation int

const NumRotations = 4

const (
	North Rotation = iota
	South
	East
	West
)

func IsOkRotation(r Rotation) bool {
	return uint8(r) < NumRotations
}
