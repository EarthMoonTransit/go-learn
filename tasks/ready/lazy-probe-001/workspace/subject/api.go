package subject

var probeCalls int
var probeValue bool

// Probe is supplied. Its return value is configured by ResetProbe.
func Probe() bool {
	probeCalls++
	return probeValue
}

// ResetProbe starts an independent, sequential observation.
func ResetProbe(value bool) { probeValue = value; probeCalls = 0 }

// ProbeCalls reports calls since the last ResetProbe.
func ProbeCalls() int { return probeCalls }
