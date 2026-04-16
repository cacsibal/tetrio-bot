package internal

type FatalityState int

const (
	Safe FatalityState = iota
	Critical
	Fatal
)

type ObligationState int

const (
	ObligationNone ObligationState = iota
	ObligationMustDownstack
	ObligationMustCancel
)

type SurgeState int

const (
	Dormant SurgeState = iota
	Building
	Active
)

type PhaseState int

const (
	Opener PhaseState = iota
	Midgame
	Endgame
)

type ClearType int

const (
	ClearTypeNone ClearType = iota
	ClearTypeSingle
	ClearTypeDouble
	ClearTypeTriple
	ClearTypeQuad
)

func FromLines(lines uint8) ClearType {
	switch lines {
	case 0:
		return ClearTypeNone
	case 1:
		return ClearTypeSingle
	case 2:
		return ClearTypeDouble
	case 3:
		return ClearTypeTriple
	case 4:
		return ClearTypeQuad
	default:
		return ClearTypeNone
	}
}

func (c ClearType) String() string {
	return []string{"None", "Single", "Double", "Triple", "Quad"}[c]
}

type ClearEvent struct {
	ClearType      ClearType
	SpinType       SpinType
	LinesCleared   uint8
	AttackSent     float32
	B2BBefore      uint8
	B2BAfter       uint8
	ComboBefore    uint32
	ComboAfter     uint32
	IsSurgeRelease bool
	IsGarbageClear bool
	IsPerfectClear bool
	Piece          Piece
}

type TransitionObservation struct {
	ResultingHeight      uint32
	ResultingB2B         uint8
	ResultingCombo       uint32
	LinesCleared         uint8
	HoldUsed             bool
	PendingGarbage       bool
	PendingPerfect       uint8
	ImminentGarbage      uint8
	SpawnEnvelopeBlocked bool
}

type GameState struct {
}
