package subject

import (
	"unicode/utf8"
)

func RuneBoundary(s string, pos int) (int, bool) {
	if pos < 0 || pos > utf8.RuneCountInString(s) {
		return 0, false
	}
	return pos, true
}
