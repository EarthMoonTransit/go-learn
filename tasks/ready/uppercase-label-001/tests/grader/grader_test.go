package subject_test

import (
	"task.local/uppercase-label-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.UpperLabel("ab\tcd"); got != "INPUT: AB\tCD" {
		t.Errorf("result got %q want %q", got, "INPUT: AB\tCD")
	}
}
