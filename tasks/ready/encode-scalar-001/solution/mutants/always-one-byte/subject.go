package subject

import (
	"unicode/utf8"
)

func EncodeScalar(r rune) (string, int, bool) {
	if !utf8.ValidRune(r) {
		return "", 0, false
	}
	s := string(r)
	return s, 1, true
}
