package subject

type Counters struct{ Accepted, Rejected, Kept int }
type Layer struct {
	Counters
	Rejected int
	Tag      string
}

func ApplyCounts(in Layer, accepted, rejected, audit int) Layer { return in }
