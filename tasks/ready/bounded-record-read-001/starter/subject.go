package subject

import (
	"errors"
	"io"
)

type Record struct{ Value uint16 }

var ErrVersion = errors.New("unsupported record version")

type RecordFailure interface {
	error
	RecordIndex() int
	Unwrap() error
}

// Add private recovery state and construct independent failures.
type recordFailure struct{}

func (e *recordFailure) Error() string    { return "" }
func (e *recordFailure) RecordIndex() int { return 0 }
func (e *recordFailure) Unwrap() error    { return nil }

func ReadRecords(src io.Reader, count int) ([]Record, error) {
	return nil, nil
}
