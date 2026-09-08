package subject_test

import (
	"task.local/column-occupancy-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.CountColumn(map[string][]bool{"A": {true}}, "a"); got != 0 {
		t.Errorf("CountColumn: got %v, want %v", got, 0)
	}
}
