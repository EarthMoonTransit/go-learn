package subject_test

import (
	"task.local/generic-find-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	v, ok := subject.Find([]int{1, 2, 4}, func(x int) bool { return x%2 == 0 })
	if v != 2 || !ok {
		t.Errorf("first even got%d,%v", v, ok)
	}
}

// Add cases for the remaining requirements in TASK.md.
