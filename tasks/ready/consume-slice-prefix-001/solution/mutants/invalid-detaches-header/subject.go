package subject

func TakePrefix(input *[]int, n int) ([]int, bool) {
	old := *input
	if n < 0 || n > len(old) {
		out := make([]int, len(old), cap(old))
		copy(out, old)
		*input = out
		return nil, false
	}
	prefix := old[:n]
	*input = old[n:]
	return prefix, true
}
