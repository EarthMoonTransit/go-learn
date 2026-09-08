package subject

import (
	"strings"
	"unicode/utf8"
)

func ValidUTF8(s string) bool { return utf8.ValidString(s) && !strings.ContainsRune(s, utf8.RuneError) }
