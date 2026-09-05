package subject

import "strings"

func RemovePrefix(s, prefix string) (string, bool) {
	index := strings.Index(s, prefix)
	if index < 0 {
		return s, false
	}
	return s[:index] + s[index+len(prefix):], true
}
