package subject_test

import (
	"task.local/bounded-value-001/subject"
	"testing"
)

func TestIntervalExamples(t *testing.T) {
	for _, c := range []struct {
		v, lo, hi, w int
		ok           bool
	}{{-3, 0, 10, 0, true}, {12, 0, 10, 10, true}, {5, 0, 10, 5, true}, {9, 7, 7, 7, true}, {5, 9, 2, 0, false}} {
		got, ok := subject.Bounded(c.v, c.lo, c.hi)
		if got != c.w || ok != c.ok {
			t.Errorf("Bounded(%d,%d,%d)=(%d,%v), want (%d,%v)", c.v, c.lo, c.hi, got, ok, c.w, c.ok)
		}
	}
}
