package subject

import "errors"

type Delta struct {
	Key    string
	Amount int
}

var ErrCount = errors.New("invalid counts")

func ApplyCounts(counts map[string]int, deltas []Delta) error {
	working := make(map[string]int, len(counts))
	for _, d := range deltas {
		working[d.Key] = counts[d.Key]
	}
	for _, d := range deltas {
		next := working[d.Key] + d.Amount
		if next < 0 {
			return ErrCount
		}
		working[d.Key] = next
	}
	clear(counts)
	for k, v := range working {
		counts[k] = v
	}
	return nil
}
