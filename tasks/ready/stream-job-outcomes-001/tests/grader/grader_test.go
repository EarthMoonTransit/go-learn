package subject_test

import (
	"errors"
	"testing"
)

func TestGraderCountsAndPanicValues(t *testing.T) {
	for _, pair := range [][2]int{{1, 8}, {2, 8}, {3, 8}, {1, 0}, {3, 1}} {
		heldJobs(t, pair[0], pair[1])
	}
	panicSequence(t, []any{-9, "different", errors.New("second object"), 28})
}
