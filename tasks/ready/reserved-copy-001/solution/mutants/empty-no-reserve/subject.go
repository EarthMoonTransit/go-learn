package subject

func ReservedCopy(in []int, extra int) []int {
	if len(in) == 0 {
		return nil
	}
	out := make([]int, len(in), len(in)+extra)
	copy(out, in)
	return out
}
