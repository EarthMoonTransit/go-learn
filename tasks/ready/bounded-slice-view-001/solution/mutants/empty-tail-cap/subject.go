package subject

func CappedView(in []int, start, end int) []int {
	if start == end {
		return in[start:start]
	}
	return in[start:end:end]
}
