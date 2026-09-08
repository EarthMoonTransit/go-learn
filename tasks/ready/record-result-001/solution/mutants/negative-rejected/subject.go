package subject

import (
	"errors"
)

type Record struct {
	Label string
	Count int
}

func NewRecord(label string, count int) (*Record, error) {
	if count < 0 {
		return nil, errors.New("negative")
	}
	return &Record{Label: label, Count: count}, nil
}
