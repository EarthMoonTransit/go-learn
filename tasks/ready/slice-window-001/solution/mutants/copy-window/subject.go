package subject

func Window(input []int, low, high int) ([]int, bool) {
	if low < 0 || high < low || high > len(input) {
		return nil, false
	}
	out := make([]int, high-low)
	copy(out, input[low:high])
	return out, true
}
