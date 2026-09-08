package subject

func RuneBoundary(s string, pos int) (int, bool) {
	if pos < 0 {
		return 0, false
	}
	n := 0
	for i := range s {
		if n == pos {
			return i, true
		}
		n++
	}
	if n == pos {
		return 0, false
	}
	return 0, false
}
