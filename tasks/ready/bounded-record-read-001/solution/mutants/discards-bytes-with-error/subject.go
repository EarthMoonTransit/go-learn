package subject

import (
	"errors"
	"fmt"
	"io"
)

type Record struct{ Value uint16 }

var ErrVersion = errors.New("unsupported record version")

type RecordFailure interface {
	error
	RecordIndex() int
	Unwrap() error
}

type recordFailure struct {
	at    int
	cause error
}

func (e *recordFailure) Error() string {
	return fmt.Sprintf("record %d: %s", e.at, e.cause.Error())
}
func (e *recordFailure) RecordIndex() int { return e.at }
func (e *recordFailure) Unwrap() error    { return e.cause }

func failedRecords(records []Record, cause error) ([]Record, error) {
	return records, &recordFailure{at: len(records), cause: cause}
}

func ReadRecords(src io.Reader, count int) ([]Record, error) {
	var records []Record
	for len(records) < count {
		var frame [3]byte
		used := 0
		for used < len(frame) {
			n, err := src.Read(frame[used:])
			if err != nil {
				return failedRecords(records, err)
			}
			used += n
			if used == len(frame) && frame[0] == 1 {
				records = append(records, Record{Value: uint16(frame[1])<<8 | uint16(frame[2])})
			}
			if err != nil && err != io.EOF {
				return failedRecords(records, err)
			}
			if used == len(frame) && frame[0] != 1 {
				return failedRecords(records, ErrVersion)
			}
			if err == io.EOF {
				if len(records) == count {
					return records, nil
				}
				return failedRecords(records, io.ErrUnexpectedEOF)
			}
		}
	}
	return records, nil
}
