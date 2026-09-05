package subject

func ReplaceByte(s string, index int, value byte) (string, bool) {
	if index < 0 || index >= len(s) {
		return s, false
	}
	data := []byte(s)
	old := data[index]
	for i, b := range data {
		if b == old {
			data[i] = value
		}
	}
	return string(data), true
}
