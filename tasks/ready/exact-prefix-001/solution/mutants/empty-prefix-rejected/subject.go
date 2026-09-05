package subject

import "strings"

func RemovePrefix(s, prefix string) (string, bool) {
	if prefix == "" {
		return s, false
	}
	return strings.CutPrefix(s, prefix)
}
