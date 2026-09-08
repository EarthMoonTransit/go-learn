package subject

func ReservedCopy(in []int, extra int) []int {
	out := make([]int, len(in)+extra)
	copy(out, in)
	return out
}
