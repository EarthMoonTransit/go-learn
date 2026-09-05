package subject

import "fmt"

func FormatStatus(label string, count int, enabled bool) string {
	return fmt.Sprintf("label=%q; count=%c; enabled=%t\n", label, count, enabled)
}
