package subject

type Sample struct {
	Label          string
	Attempts, Hits int
}

func (s Sample) Average() float64 { return 0 }
