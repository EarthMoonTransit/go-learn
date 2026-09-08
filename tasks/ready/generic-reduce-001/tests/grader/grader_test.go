package subject_test

import (
	"task.local/generic-reduce-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.Reduce([]int{}, func(a, b int) int { return a + b }, 7); got != 7 {
		t.Errorf("allocated empty got%d", got)
	}
}
