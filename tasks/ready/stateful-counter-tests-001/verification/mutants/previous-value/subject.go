package subject

func NewCounter(start int) func(int) int {
	value := start
	return func(delta int) int { old := value; value += delta; return old }
}
