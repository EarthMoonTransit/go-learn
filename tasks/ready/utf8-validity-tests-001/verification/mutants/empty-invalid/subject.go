package subject

import (
	"unicode/utf8"
)

func ValidUTF8(s string) bool { return s != "" && utf8.ValidString(s) }
