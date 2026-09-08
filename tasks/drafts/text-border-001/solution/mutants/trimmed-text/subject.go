package subject

import (
	"strings"
)

func AddBorder(text string, width int) string {
	line := strings.Repeat("*", width)
	return line + "\n" + strings.TrimSpace(text) + "\n" + line
}
