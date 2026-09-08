package subject_test

import (
	"task.local/capture-panic-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	err, p := subject.Capture(func() error { return nil })
	if err != nil || p {
		t.Errorf("fresh normal got%v,%v", err, p)
	}
}
