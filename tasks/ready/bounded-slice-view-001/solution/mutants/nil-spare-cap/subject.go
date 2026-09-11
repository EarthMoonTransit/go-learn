package subject

func CappedView(in []int, start, end int) []int {
	if in == nil {
		return make([]int, 0, 1)
	}
	return in[start:end:end]
}
