package subject

func RuneBoundary(s string, pos int) (int, bool) {
	if pos < 0 || pos > len(s) {
		return 0, false
	}
	return pos, true
}
