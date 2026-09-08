package subject_test

import (
	"task.local/floating-categories-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.FloatKind(1e300); got != "finite" {
		t.Errorf("large finite got%q", got)
	}
}
