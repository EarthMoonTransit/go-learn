package subject

import "unicode/utf8"

func RepairUTF8(s string) string {
	if s == "" {
		return string(utf8.RuneError)
	}
	return string([]rune(s))
}
