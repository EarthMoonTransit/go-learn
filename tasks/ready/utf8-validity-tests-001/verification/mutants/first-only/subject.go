package subject

import (
	"unicode/utf8"
)

func ValidUTF8(s string) bool {
	if s == "" {
		return true
	}
	r, n := utf8.DecodeRuneInString(s)
	return r != utf8.RuneError || n > 1
}
