package subject

func ByteWindow(s string, start, end int) (string, bool) {
	if start < 0 {
		start = 0
	}
	if start < 0 || end < start || end > len(s) {
		return "", false
	}
	return s[start:end], true
}
