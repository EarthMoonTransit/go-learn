package subject_test

import (
	"testing"

	"task.local/clamp-tests-001/subject"
)

func TestClampInsideRange(t *testing.T) {
	if got := subject.Clamp(5, 0, 10); got != 5 {
		t.Fatalf("Clamp returned %d", got)
	}
}
