package subject

func Window(input []int, low, high int) ([]int, bool) {
	if low < 0 || high < low || high > len(input) {
		clear(input)
		return nil, false
	}
	return input[low:high], true
}
