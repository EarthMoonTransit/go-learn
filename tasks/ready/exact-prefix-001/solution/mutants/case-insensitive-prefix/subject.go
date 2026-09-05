package subject

import "strings"

func RemovePrefix(s, prefix string) (string, bool) {
	if len(s) >= len(prefix) && strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix)) {
		return s[len(prefix):], true
	}
	return s, false
}
