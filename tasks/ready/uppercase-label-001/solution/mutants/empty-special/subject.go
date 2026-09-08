package subject

import (
	"strings"
)

func UpperLabel(text string) string {
	if text == "" {
		return ""
	}
	return "INPUT: " + strings.ToUpper(text)
}
