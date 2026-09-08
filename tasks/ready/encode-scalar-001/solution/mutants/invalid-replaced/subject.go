package subject

import (
	"unicode/utf8"
)

func EncodeScalar(r rune) (string, int, bool) {
	if !utf8.ValidRune(r) {
		return string(utf8.RuneError), 3, true
	}
	s := string(r)
	return s, len(s), true
}
