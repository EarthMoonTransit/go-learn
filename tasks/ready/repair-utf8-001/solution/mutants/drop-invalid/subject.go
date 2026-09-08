package subject

import (
	"strings"
)

func RepairUTF8(s string) string { return strings.ToValidUTF8(s, "") }
