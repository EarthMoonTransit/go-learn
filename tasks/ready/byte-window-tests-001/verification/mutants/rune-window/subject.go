package subject

func ByteWindow(s string, start, end int) (string, bool) {
	data := []rune(s)
	if start < 0 || end < start || end > len(data) {
		return "", false
	}
	return string(data[start:end]), true
}
