package subject

func Contains(lo, hi, x int) bool {
	if lo == hi {
		return x == lo
	}
	return lo <= x && x < hi
}
