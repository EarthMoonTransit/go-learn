package subject_test

import (
	"math"
	"task.local/shape-contract-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	var shape subject.Shape = subject.Rectangle{Width: 12, Height: 6}
	got := shape.Area()
	if math.IsNaN(got) || math.Abs(got-72) > 1e-9 {
		t.Errorf("rectangle got %g want72", got)
	}
}

// Add cases for the remaining requirements in TASK.md.
