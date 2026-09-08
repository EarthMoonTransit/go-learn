package subject_test

import (
	"task.local/whole-minute-output-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.SuccessfulPerMinute(1105, 90); got != 16 {
		t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 16)
	}
}
