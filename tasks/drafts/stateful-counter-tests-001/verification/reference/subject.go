package subject

func NewCounter(start int) func(int) int {
	value := start
	return func(delta int) int { value += delta; return value }
}
