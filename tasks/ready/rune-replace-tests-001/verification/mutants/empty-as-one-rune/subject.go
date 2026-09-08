package subject

func Replace(log string, oldRune, newRune rune) string {
	if log == "" {
		return string(newRune)
	}
	r := []rune(log)
	for i, v := range r {
		if v == oldRune {
			r[i] = newRune
		}
	}
	return string(r)
}
