package subject

import "strings"

type Record struct {
	Label string
	Count int
}

func NewRecord(label string, count int) (*Record, error) {
	return &Record{Label: strings.TrimSpace(label), Count: count}, nil
}
