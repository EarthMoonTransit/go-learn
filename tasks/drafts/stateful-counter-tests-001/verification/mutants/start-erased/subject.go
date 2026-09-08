package subject

func NewCounter(start int) func(int) int {
	value := 0
	return func(delta int) int { value += delta; return value }
}
