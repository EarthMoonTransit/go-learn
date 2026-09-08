package subject

func Replace(log string, oldRune, newRune rune) string {
	out := make([]byte, len(log))
	n := 0
	for _, r := range log {
		if r == oldRune {
			r = newRune
		}
		n += copy(out[n:], string(r))
		if n == len(out) {
			break
		}
	}
	return string(out[:n])
}
