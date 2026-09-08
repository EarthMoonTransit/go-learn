package subject

import (
	"strings"
)

func ChooseLabel(a, b string) string {
	if b < a {
		a = b
	}
	return "selected: " + strings.TrimSpace(a)
}
