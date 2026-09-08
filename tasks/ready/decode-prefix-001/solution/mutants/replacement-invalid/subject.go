package subject

import (
	"unicode/utf8"
)

func DecodePrefix(s string) (rune, int, bool) {
	if s == "" {
		return 0, 0, false
	}
	r, n := utf8.DecodeRuneInString(s)
	return r, n, r != utf8.RuneError
}
