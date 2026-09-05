package subject

func ReplaceByte(s string, index int, value byte) (string, bool) {
	if index < 0 || index >= len(s)-1 {
		return s, false
	}
	data := []byte(s)
	data[index] = value
	return string(data), true
}
