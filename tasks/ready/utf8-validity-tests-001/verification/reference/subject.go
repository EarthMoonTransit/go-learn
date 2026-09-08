package subject

import (
	"unicode/utf8"
)

func ValidUTF8(s string) bool { return utf8.ValidString(s) }
