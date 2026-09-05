package subject

func ReplaceByte(s string, index int, value byte) (string, bool) {
	if index < 0 || index >= len(s) {
		return s, false
	}
	return s[:index] + string(value) + s[index+1:], true
}
