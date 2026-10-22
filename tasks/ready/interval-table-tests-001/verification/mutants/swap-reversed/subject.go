package subject

func Contains(lo, hi, x int) bool {
	if lo > hi {
		lo, hi = hi, lo
	}
	return lo <= x && x < hi
}
