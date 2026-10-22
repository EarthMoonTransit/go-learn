package subject

func ReplaceRuneRange(s string, lo, hi int, replacement string) (string, bool) {
	values := []byte(s)
	if lo < 0 || hi < lo || hi > len(values) {
		return s, false
	}
	out := append([]byte(nil), values[:lo]...)
	out = append(out, []byte(replacement)...)
	out = append(out, values[hi:]...)
	return string(out), true
}
