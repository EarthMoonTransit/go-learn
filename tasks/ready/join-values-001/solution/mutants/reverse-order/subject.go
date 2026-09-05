package subject

import (
	"fmt"
	"strings"
)

func FormatValues(prefix string, values []int) string {
	var result strings.Builder
	result.WriteString(prefix)
	result.WriteString("[")
	for i := range values {
		value := values[len(values)-1-i]
		if i > 0 {
			result.WriteString(",")
		}
		result.WriteString(fmt.Sprintf("%d", value))
	}
	result.WriteString("]")
	return result.String()
}
