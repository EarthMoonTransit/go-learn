package subject

import "fmt"

func FormatStatus(label string, count int, enabled bool) string {
	flag := 0
	if enabled {
		flag = 1
	}
	return fmt.Sprintf("label=%s; count=%d; enabled=%d", label, count, flag)
}
