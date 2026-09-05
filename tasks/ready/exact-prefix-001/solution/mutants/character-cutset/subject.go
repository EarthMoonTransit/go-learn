package subject

import "strings"

func RemovePrefix(s, prefix string) (string, bool) {
	result := strings.TrimLeft(s, prefix)
	return result, result != s
}
