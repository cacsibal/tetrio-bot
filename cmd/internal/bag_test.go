package internal

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTrackerHasAllRemaining(t *testing.T) {
	tracker := BagTracker{}
	remaining := tracker.Remaining()

	assert.Equal(t, 7, len(remaining))
	assert.Equal(t, []Piece{I, O, T, L, J, S, Z}, remaining)
}

func TestConsumeReducesRemaining(t *testing.T) {
	tracker := BagTracker{}
	tracker.Consume(I)
	tracker.Consume(T)
	tracker.Consume(S)

	remaining := tracker.Remaining()

	assert.Equal(t, 4, len(remaining))
	assert.False(t, slices.Contains(remaining, I))
	assert.False(t, slices.Contains(remaining, T))
	assert.False(t, slices.Contains(remaining, S))
	assert.True(t, slices.Contains(remaining, O))
	assert.True(t, slices.Contains(remaining, L))
	assert.True(t, slices.Contains(remaining, J))
	assert.True(t, slices.Contains(remaining, Z))
}

func TestFullBagCycleResets(t *testing.T) {
	tracker := BagTracker{}
	for p := range AllPieces {
		tracker.Consume(Piece(p))
	}
	assert.Equal(t, uint8(7), tracker.Count())
	tracker.Consume(T)
	assert.Equal(t, uint8(1), tracker.Count())
	remaining := tracker.Remaining()
	assert.Equal(t, 6, len(remaining))
	assert.False(t, slices.Contains(remaining, T))
}

func TestPredictNextConsumesAndReturnsRemaining(t *testing.T) {
	tracker := BagTracker{}
	tracker.Consume(I)
	tracker.Consume(O)

	queue := []Piece{T, L, J}
	remaining := tracker.PredictNext(queue)
	assert.Equal(t, 2, len(remaining))
	assert.Equal(t, []Piece{S, Z}, remaining)
}
