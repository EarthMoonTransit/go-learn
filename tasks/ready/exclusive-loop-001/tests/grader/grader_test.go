package subject_test

import (
	"task.local/exclusive-loop-001/subject"
	"testing"
)

func TestBoundariesAndProgress(t *testing.T) {
	for limit := 0; limit <= 10000; limit++ {
		if got := subject.CountEven(limit); got != (limit+1)/2 {
			t.Fatalf("wrong count for limit %d", limit)
		}
	}
}
