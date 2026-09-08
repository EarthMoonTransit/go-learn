package subject_test

import (
	"task.local/seven-value-window-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.WeekSum([]int{1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3}, 3); got != 21 {
		t.Errorf("WeekSum: got %v, want %v", got, 21)
	}
}
