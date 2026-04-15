package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceTableINorth(t *testing.T) {
	pc := PieceTable(I, North)

	assert.Equal(t, -1, pc[0].X)
	assert.Equal(t, 0, pc[0].Y)

	assert.Equal(t, 1, pc[1].X)
	assert.Equal(t, 0, pc[1].Y)

	assert.Equal(t, 2, pc[2].X)
	assert.Equal(t, 0, pc[2].Y)
}

func TestPieceTableIEast(t *testing.T) {
	pc := PieceTable(I, East)

	assert.Equal(t, 0, pc[0].X)
	assert.Equal(t, 1, pc[0].Y)

	assert.Equal(t, 0, pc[1].X)
	assert.Equal(t, -1, pc[1].Y)

	assert.Equal(t, 0, pc[2].X)
	assert.Equal(t, -2, pc[2].Y)
}
