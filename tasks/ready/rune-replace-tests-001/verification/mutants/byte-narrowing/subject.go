package subject

func Replace(log string, oldRune, newRune rune) string {
	b := []byte(log)
	for i, v := range b {
		if rune(v) == oldRune {
			b[i] = byte(newRune)
		}
	}
	return string(b)
}
