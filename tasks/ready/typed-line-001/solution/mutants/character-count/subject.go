package subject

import "fmt"

func FormatStatus(label string, count int, enabled bool) string {
	return fmt.Sprintf("label=%s; count=%c; enabled=%t", label, count, enabled)
}
