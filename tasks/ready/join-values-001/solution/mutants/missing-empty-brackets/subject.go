package subject

import (
	"fmt"
	"strings"
)

func FormatValues(prefix string, values []int) string {
	if len(values) == 0 {
		return prefix
	}
	var result strings.Builder
	result.WriteString(prefix)
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
