package subject_test

import (
	"task.local/clamp-tests-001/subject"
	"testing"
)

func TestClampInsideRange(t *testing.T) {
	if got := subject.Clamp(5, 0, 10); got != 5 {
		t.Fatalf("Clamp(5,0,10)=%d, want 5", got)
	}
}
