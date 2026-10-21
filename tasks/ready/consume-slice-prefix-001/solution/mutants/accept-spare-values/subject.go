package subject

func TakePrefix(input *[]int, n int) ([]int, bool) {
	old := *input
	if n < 0 || n > cap(old) {
		return nil, false
	}
	prefix := old[:n]
	if n > len(old) {
		*input = old[n:n]
	} else {
		*input = old[n:]
	}
	return prefix, true
}
