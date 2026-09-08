package subject

import (
	"strings"
)

func AddBorder(text string, width int) string {
	if text == "" {
		return ""
	}
	line := strings.Repeat("*", width)
	return line + "\n" + text + "\n" + line
}
