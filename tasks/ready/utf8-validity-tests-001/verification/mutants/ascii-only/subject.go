package subject

func ValidUTF8(s string) bool {
	for _, b := range []byte(s) {
		if b > 127 {
			return false
		}
	}
	return true
}
