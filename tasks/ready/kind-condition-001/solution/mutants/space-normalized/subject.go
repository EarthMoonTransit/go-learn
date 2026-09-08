package subject

import (
	"strings"
)

func RequiresPermit(kind string) bool {
	kind = strings.TrimSpace(kind)
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}
