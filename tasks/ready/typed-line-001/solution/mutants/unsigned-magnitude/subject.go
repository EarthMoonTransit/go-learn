package subject

import "fmt"

func FormatStatus(label string, count int, enabled bool) string {
	if count < 0 {
		count = -count
	}
	return fmt.Sprintf("label=%s; count=%d; enabled=%t", label, count, enabled)
}
