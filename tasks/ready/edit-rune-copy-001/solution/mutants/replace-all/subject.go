package subject

func EditRune(s string, index int, replacement rune) (string, bool) {
	r := []rune(s)
	if index < 0 || index >= len(r) {
		return s, false
	}
	old := r[index]
	for i, v := range r {
		if v == old {
			r[i] = replacement
		}
	}
	return string(r), true
}
