package internal

const (
	AttackSingle uint8 = 0
	AttackDouble uint8 = 1
	AttackTriple uint8 = 2
	AttackQuad   uint8 = 4

	SpinMini       uint8 = 0
	Spin           uint8 = 0
	SpinMiniSingle uint8 = 0
	SpinSingle     uint8 = 2
	SpinMiniDouble uint8 = 1
	SpinDouble     uint8 = 4
	SpinMiniTriple uint8 = 2
	SpinTriple     uint8 = 6
	SpinQuad       uint8 = 10

	B2BBonus        uint8   = 1
	B2BChainingLog  float32 = 0.8
	ComboBonus      float32 = 0.25
	ComboFloorScale float32 = 1.25
)

var ModernComboTable = [13]uint8{0, 1, 1, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4}
