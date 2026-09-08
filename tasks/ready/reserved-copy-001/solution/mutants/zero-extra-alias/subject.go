package subject

func ReservedCopy(in []int, extra int) []int {
	if extra == 0 {
		return in
	}
	out := make([]int, len(in), len(in)+extra)
	copy(out, in)
	return out
}
