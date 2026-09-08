package subject

func Replace(log string, oldRune, newRune rune) string {
	r := []rune(log)
	for i, v := range r {
		if v == oldRune {
			r[i] = rune(byte(newRune))
		}
	}
	return string(r)
}
