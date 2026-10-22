package subject

import (
	"sort"
	"strconv"
	"strings"
)

func Summary(values map[string]int) string {
	if len(values) == 0 {
		return "(empty)"
	}
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		v := values[k]
		parts = append(parts, "="+strconv.Itoa(v))
	}
	return strings.Join(parts, ", ")
}
