package subject

func TakePrefix(input *[]int, n int) ([]int, bool) {
	old := *input
	if n < 0 || n > len(old) {
		return nil, false
	}
	prefix := make([]int, n, cap(old))
	copy(prefix, old[:n])
	*input = old[n:]
	return prefix, true
}
