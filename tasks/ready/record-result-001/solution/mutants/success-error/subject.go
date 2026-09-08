package subject

import (
	"errors"
)

type Record struct {
	Label string
	Count int
}

func NewRecord(label string, count int) (*Record, error) {
	return &Record{Label: label, Count: count}, errors.New("unexpected")
}
