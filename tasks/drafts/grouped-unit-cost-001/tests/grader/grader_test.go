package subject_test

import (
	"task.local/grouped-unit-cost-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.BatchCost(148); got != uint(1410000) {
		t.Errorf("BatchCost: got %v, want %v", got, uint(1410000))
	}
}
