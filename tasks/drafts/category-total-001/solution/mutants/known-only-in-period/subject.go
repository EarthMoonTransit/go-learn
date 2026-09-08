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

func CategoryTotal(in []Record, p Period, category string) (float64, error) {
	known := false
	total := 0.0
	for _, r := range in {
		if r.Category == category {
			if r.Day >= p.From && r.Day <= p.To {
				known = true
				total += r.Amount
			}
		}
	}
	if !known {
		return 0, ErrCategory
	}
	return total, nil
}
