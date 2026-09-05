package subject

import "fmt"

func FormatStatus(label string, count int, enabled bool) string {
	return fmt.Sprintf("label="+label+"; count=%d; enabled=%t", count, enabled)
}
