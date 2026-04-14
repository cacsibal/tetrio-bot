package internal

type SpinType int

const (
	NoSpin SpinType = iota
	Mini
	Full
)

var AllSpinTypes = []SpinType{NoSpin, Mini, Full}

func (s SpinType) String() string {
	if int(s) < 0 || int(s) >= len(AllSpinTypes) {
		return "Unknown"
	}

	return []string{"NoSpin", "Mini", "Full"}[s]
}
