package subject_test

import (
	"task.local/shared-counter-update-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	n := 0
	subject.AddCounter(&n, 1)
	if n != 1 {
		t.Errorf("zero initial got %d want 1", n)
	}
}
