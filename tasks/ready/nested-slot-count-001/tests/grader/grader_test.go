package subject_test

import (
	"task.local/nested-slot-count-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.CountSlots(map[string][]bool{"A": {true, true, true, true, true, true, true, true}, "B": {false, false, false, false, false, false, false, false}}); got != 16 {
		t.Errorf("CountSlots: got %v, want %v", got, 16)
	}
}
