package subject_test

import (
	"task.local/kind-condition-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.RequiresPermit("stroller"); got != false {
		t.Errorf("RequiresPermit: got %v, want %v", got, false)
	}
}
