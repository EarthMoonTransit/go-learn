package subject

func ByteWindow(s string, start, end int) (string, bool) {
	if start > end {
		start, end = end, start
	}
	if start < 0 || end < start || end > len(s) {
		return "", false
	}
	return s[start:end], true
}
