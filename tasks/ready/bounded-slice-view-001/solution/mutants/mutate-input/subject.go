package subject

func CappedView(in []int, start, end int) []int {
	if start < end {
		in[start] = 0
	}
	return in[start:end:end]
}
