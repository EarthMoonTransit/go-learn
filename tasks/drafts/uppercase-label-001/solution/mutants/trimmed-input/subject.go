package subject

import (
	"strings"
)

func UpperLabel(text string) string { return "INPUT: " + strings.ToUpper(strings.TrimSpace(text)) }
