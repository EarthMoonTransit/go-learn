package subject

import (
	"unicode/utf8"
)

func WithinLimit(log string, limit int) bool {
	if limit < 0 {
		return true
	}
	return utf8.RuneCountInString(log) <= limit
}
