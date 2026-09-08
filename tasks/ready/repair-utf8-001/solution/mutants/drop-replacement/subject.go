package subject

import (
	"unicode/utf8"
)

func RepairUTF8(s string) string {
	out := []rune{}
	for _, r := range s {
		if r != utf8.RuneError {
			out = append(out, r)
		}
	}
	return string(out)
}
