package subject_test

import (
	"task.local/month-enumeration-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.GetMonth(subject.Dec); got != 12 {
		t.Errorf("last month got%d want12", got)
	}
}
