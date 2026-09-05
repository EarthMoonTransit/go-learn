package subject

import "strings"

func RemovePrefix(s, prefix string) (string, bool) {
	if !strings.HasPrefix(s, prefix) {
		return "", false
	}
	return strings.CutPrefix(s, prefix)
}
