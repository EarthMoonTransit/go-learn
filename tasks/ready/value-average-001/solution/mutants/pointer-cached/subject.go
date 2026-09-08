package subject

type Sample struct {
	Label          string
	Attempts, Hits int
}

func (s Sample) Average() float64 {
	if s.Attempts == 0 {
		return 0
	}
	return float64(s.Hits) / float64(s.Attempts)
}

func MethodFunctions() (func(Sample) float64, func(*Sample) float64) {
	cached := 0.0
	seen := false
	return Sample.Average, func(s *Sample) float64 {
		if !seen {
			cached = s.Average()
			seen = true
		}
		return cached
	}
}
