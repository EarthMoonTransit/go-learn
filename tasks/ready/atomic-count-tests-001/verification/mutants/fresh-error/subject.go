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
			return errors.New("invalid counts")
		}
		working[d.Key] = next
	}
	clear(counts)
	for k, v := range working {
		counts[k] = v
	}
	return nil
}
