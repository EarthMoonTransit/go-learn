package subject

func EditRune(s string, index int, replacement rune) (string, bool) {
	r := []rune(s)
	if index < 0 || index >= len(r) {
		return s, false
	}
	r[index] = replacement
	return string(r), true
}
