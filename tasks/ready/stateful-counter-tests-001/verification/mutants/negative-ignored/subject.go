package subject

func NewCounter(start int) func(int) int {
	value := start
	return func(delta int) int {
		if delta > 0 {
			value += delta
		}
		return value
	}
}
