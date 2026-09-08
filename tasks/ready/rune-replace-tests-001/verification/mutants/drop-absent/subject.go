package subject

import (
	"strings"
)

func Replace(log string, oldRune, newRune rune) string {
	if !strings.ContainsRune(log, oldRune) {
		return ""
	}
	r := []rune(log)
	for i, v := range r {
		if v == oldRune {
			r[i] = newRune
		}
	}
	return string(r)
}
