package subject

import (
	"errors"
)

type Record struct {
	Day      int
	Amount   float64
	Category string
}
type Period struct{ From, To int }

var ErrCategory = errors.New("unknown category")

func CategoryTotal(in []Record, p Period, category string) (float64, error) { return 0, nil }
