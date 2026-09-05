package subject

import "strings"

func RemovePrefix(s, prefix string) (string, bool) {
	return strings.CutPrefix(s, prefix)
}
