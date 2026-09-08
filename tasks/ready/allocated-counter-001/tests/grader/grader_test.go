package subject_test

import (
	"task.local/allocated-counter-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	p := subject.NewCounter(98)
	if p == nil || *p != 98 {
		t.Errorf("expected pointer containing 98, got %v", p)
	}
}
