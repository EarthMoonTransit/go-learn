package subject

import "errors"

type Delta struct {
	Key    string
	Amount int
}

var ErrCount = errors.New("invalid counts")

func ApplyCounts(counts map[string]int, deltas []Delta) error {
	working := make(map[string]int, len(counts))
	for k, v := range counts {
		working[k] = v
	}
	for _, d := range deltas {
		next := working[d.Key] + d.Amount
		if next < 0 {
			return ErrCount
		}
		working[d.Key] = next
	}
	for k, v := range working {
		if v == 0 {
			delete(working, k)
		}
	}
	clear(counts)
	for k, v := range working {
		counts[k] = v
	}
	return nil
}
