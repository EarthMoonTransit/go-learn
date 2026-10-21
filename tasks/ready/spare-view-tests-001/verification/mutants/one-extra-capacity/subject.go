package subject

func SpareView(in []int, low, high, spare int) []int {
	end := high + spare
	if end < cap(in) {
		end++
	}
	return in[low:high:end]
}
