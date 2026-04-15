package internal

type BagTracker struct {
	seen  [7]bool
	count uint8
}

func (b *BagTracker) Reset() {
	b.seen = [7]bool{}
	b.count = 0
}

func (b *BagTracker) New() BagTracker {
	return BagTracker{}
}

func (b *BagTracker) Consume(p Piece) {
	if b.count >= 7 {
		b.Reset()
	}

	idx := p
	if b.seen[idx] {
		b.Reset()
	}

	b.seen[idx] = true
	b.count++
}

func (b *BagTracker) Remaining() []Piece {
	remaining := make([]Piece, 0, 7-b.count)

	for _, p := range AllPieces {
		if !b.seen[p] {
			remaining = append(remaining, p)
		}
	}

	return remaining
}

func (b *BagTracker) PredictNext(queue []Piece) []Piece {
	for _, piece := range queue {
		b.Consume(piece)
	}

	return b.Remaining()
}

func (b *BagTracker) Count() uint8 {
	return b.count
}

func (b *BagTracker) ExtendQueue(queue []Piece, current Piece, hold *Piece) []Piece {
	if hold != nil {
		b.Consume(*hold)
	}

	b.Consume(current)
	for _, p := range queue {
		b.Consume(p)
	}

	remaining := b.Remaining()
	extended := make([]Piece, len(queue))
	copy(extended, queue)

	if len(remaining) <= 2 && len(remaining) > 0 {
		extended = append(extended, remaining...)
	}

	return extended
}
