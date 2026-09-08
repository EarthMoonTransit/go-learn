package subject

import (
	"unicode/utf8"
)

func EncodeScalar(r rune) (string, int, bool) {
	if r < 0 || r > utf8.MaxRune {
		return "", 0, false
	}
	s := string(r)
	return s, len(s), true
}
