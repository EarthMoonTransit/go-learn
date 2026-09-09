package subject_test

import (
	"task.local/diagnostic-helper-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	r := &recordReporter{}
	subject.AssertInt(r, 0, -1)
	if len(r.messages) != 1 || r.messages[0] != "got 0, want -1" {
		t.Errorf("zero mismatch got%v", r.messages)
	}
}
