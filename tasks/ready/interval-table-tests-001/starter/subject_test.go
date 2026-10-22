package subject_test

import (
	subject "task.local/interval-table-tests-001/subject"
	"testing"
)

func TestInterior(t *testing.T) {
	if got := subject.Contains(0, 10, 5); !got {
		t.Errorf("Contains(0,10,5)=%t; want true", got)
	}
}
