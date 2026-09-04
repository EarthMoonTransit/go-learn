package subject_test

import (
	"task.local/exclusive-loop-001/subject"
	"testing"
)

func TestSmallIntervals(t *testing.T) {
	for _, c := range []struct{ limit, want int }{{0, 0}, {1, 1}, {2, 1}, {5, 3}} {
		if got := subject.CountEven(c.limit); got != c.want {
			t.Errorf("CountEven(%d)=%d, want %d", c.limit, got, c.want)
		}
	}
}
