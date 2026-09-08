package subject

type Sample struct {
	Label          string
	Attempts, Hits int
}

func (s Sample) Average() float64 {
	if s.Attempts == 0 {
		return 0
	}
	return float64(s.Attempts) / float64(s.Hits)
}
