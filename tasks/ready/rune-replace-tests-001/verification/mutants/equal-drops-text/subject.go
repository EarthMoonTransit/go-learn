package subject

func Replace(log string, oldRune, newRune rune) string {
	if oldRune == newRune {
		return ""
	}
	r := []rune(log)
	for i, v := range r {
		if v == oldRune {
			r[i] = newRune
		}
	}
	return string(r)
}
