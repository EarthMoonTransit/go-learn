package subject_test

import (
	"math"
	"task.local/independent-scaled-copy-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	got := subject.ScaledCopy([]float64{2}, 1)
	if len(got) != 1 || math.IsNaN(got[0]) || math.Abs(got[0]-1) > 1e-9 {
		t.Errorf("downscale got %v want [1]", got)
	}
}
