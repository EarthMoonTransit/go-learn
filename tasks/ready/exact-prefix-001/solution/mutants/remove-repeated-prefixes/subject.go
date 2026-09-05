package subject

import "strings"

func RemovePrefix(s, prefix string) (string, bool) {
	found := strings.HasPrefix(s, prefix)
	for prefix != "" && strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
	}
	return s, found
}
