package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveRoundtrip(t *testing.T) {
	m := NewMove(T, East, 5, 10, false)
	assert.Equal(t, T, m.Piece())
	assert.Equal(t, East, m.Rotation())
	assert.Equal(t, 5, m.X())
	assert.Equal(t, 10, m.Y())
	assert.Equal(t, NoSpin, m.Spin())
}

func TestMoveTspin(t *testing.T) {
	m := NewTSpinMove(South, 3, 5, true)
	assert.Equal(t, T, m.Piece())
	assert.Equal(t, Full, m.Spin())
}

func TestMoveTspinMini(t *testing.T) {
	m := NewTSpinMove(North, 4, 0, false)
	assert.Equal(t, T, m.Piece())
	assert.Equal(t, Mini, m.Spin())
}
