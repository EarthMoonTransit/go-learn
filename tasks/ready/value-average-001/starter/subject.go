package subject

type Sample struct {
	Label          string
	Attempts, Hits int
}

func (s Sample) Average() float64 { return 0 }

func MethodFunctions() (func(Sample) float64, func(*Sample) float64) { return nil, nil }
