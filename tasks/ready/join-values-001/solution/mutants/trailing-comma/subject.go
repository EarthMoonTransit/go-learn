package subject

import (
	"fmt"
	"strings"
)

func FormatValues(prefix string, values []int) string {
	var result strings.Builder
	result.WriteString(prefix)
	result.WriteString("[")
	for i, value := range values {
		_ = i
		result.WriteString(fmt.Sprintf("%d", value))
		result.WriteString(",")
	}
	result.WriteString("]")
	return result.String()
}
