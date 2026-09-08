package subject_test

import (
	"task.local/remaining-constant-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.RemainingMinutes(39); got != 1 {
		t.Errorf("RemainingMinutes: got %v, want %v", got, 1)
	}
}
