package subject

import (
	"strings"
)

func AddBorder(text string, width int) string {
	if width == 0 {
		return text
	}
	line := strings.Repeat("*", width)
	return line + "\n" + text + "\n" + line
}
