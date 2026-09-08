package subject_test

import (
	"task.local/independent-scaled-copy-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	got := subject.ScaledCopy([]float64{2}, 1)
	if len(got) != 1 || got[0] != 1 {
		t.Errorf("downscale got %v want [1]", got)
	}
}
