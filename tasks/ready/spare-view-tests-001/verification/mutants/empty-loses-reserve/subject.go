package subject

func SpareView(in []int, low, high, spare int) []int {
	if low == high {
		return nil
	}
	return in[low : high : high+spare]
}
