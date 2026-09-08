package subject

import (
	"unicode/utf8"
)

func WithinLimit(log string, limit int) bool { return utf8.RuneCountInString(log) < limit }
