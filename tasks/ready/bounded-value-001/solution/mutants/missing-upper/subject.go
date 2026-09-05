package subject

func Bounded(value, low, high int) (int, bool) {
	if low > high {
		return 0, false
	}
	if value < low {
		return low, true
	}

	return value, true
}
