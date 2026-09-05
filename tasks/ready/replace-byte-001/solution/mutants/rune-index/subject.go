package subject

func ReplaceByte(s string, index int, value byte) (string, bool) {
	data := []rune(s)
	if index < 0 || index >= len(data) {
		return s, false
	}
	data[index] = rune(value)
	return string(data), true
}
