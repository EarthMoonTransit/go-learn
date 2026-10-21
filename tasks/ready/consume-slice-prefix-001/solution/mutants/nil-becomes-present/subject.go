package subject

func TakePrefix(input *[]int, n int) ([]int, bool) {
	old := *input
	if old == nil && n == 0 {
		*input = make([]int, 0)
		return make([]int, 0), true
	}
	if n < 0 || n > len(old) {
		return nil, false
	}
	prefix := old[:n]
	*input = old[n:]
	return prefix, true
}
