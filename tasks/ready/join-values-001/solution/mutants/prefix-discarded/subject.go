package subject

import (
	"fmt"
	"strings"
)

func FormatValues(prefix string, values []int) string {
	var result strings.Builder
	_ = prefix
	result.WriteString("[")
	for i, value := range values {
		if i > 0 {
			result.WriteString(",")
		}
		result.WriteString(fmt.Sprintf("%d", value))
	}
	result.WriteString("]")
	return result.String()
}
