package subject

func ReplaceRuneRange(s string, lo, hi int, replacement string) (string, bool) {
	values := []rune(s)
	if lo < 0 || hi < lo || hi > len(values) || lo == len(values) {
		return s, false
	}
	out := append([]rune(nil), values[:lo]...)
	out = append(out, []rune(replacement)...)
	out = append(out, values[hi:]...)
	return string(out), true
}
