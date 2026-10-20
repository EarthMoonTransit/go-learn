package subject

type Counters struct{ Accepted, Rejected, Kept int }
type Layer struct {
	Counters
	Rejected int
	Tag      string
}

func ApplyCounts(in Layer, accepted, rejected, audit int) Layer {
	in.Accepted += accepted
	in.Counters.Rejected += audit
	in.Rejected += rejected
	return in
}
