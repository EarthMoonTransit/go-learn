package subject

func Bounded(value, low, high int) (int, bool) {
	if low > high {
		low, high = high, low
	}
	if value < low {
		return low, true
	}
	if value > high {
		return high, true
	}
	return value, true
}
