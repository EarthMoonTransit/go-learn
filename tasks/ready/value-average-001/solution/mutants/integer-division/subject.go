package subject

type Sample struct {
	Label          string
	Attempts, Hits int
}

func (s Sample) Average() float64 {
	if s.Attempts == 0 {
		return 0
	}
	return float64(s.Hits / s.Attempts)
}

func MethodFunctions() (func(Sample) float64, func(*Sample) float64) {
	return Sample.Average, (*Sample).Average
}
