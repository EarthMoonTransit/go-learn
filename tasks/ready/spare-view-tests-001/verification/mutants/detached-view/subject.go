package subject

func SpareView(in []int, low, high, spare int) []int {
	out := make([]int, high-low, high-low+spare)
	copy(out, in[low:high])
	return out
}
