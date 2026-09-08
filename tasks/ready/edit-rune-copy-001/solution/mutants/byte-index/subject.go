package subject

func EditRune(s string, index int, replacement rune) (string, bool) {
	b := []byte(s)
	if index < 0 || index >= len(b) {
		return s, false
	}
	b[index] = byte(replacement)
	return string(b), true
}
