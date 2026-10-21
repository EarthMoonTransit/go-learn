package subject

func SpareView(in []int, low, high, spare int) []int {
	out := in[low : high : high+spare]
	for i := range out {
		out[i] = 0
	}
	return out
}
