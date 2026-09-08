package subject

import (
	"strings"
)

func ChooseLabel(a, b string) string {
	if strings.ToLower(b) < strings.ToLower(a) {
		a = b
	}
	return "selected: " + a
}
