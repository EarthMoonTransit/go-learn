package subject

import (
	"strings"
)

type Record struct {
	Day      int
	Amount   float64
	Category string
}

func InCategory(category string) func(Record) bool {
	return func(r Record) bool { return strings.TrimSpace(r.Category) == strings.TrimSpace(category) }
}
