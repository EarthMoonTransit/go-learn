package subject_test

import (
	"math"
	"task.local/two-quantity-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	gotA, gotB := subject.Quantities([]string{"solid", "liquid", "solid"})
	if gotA != 100 || math.IsNaN(gotB) || math.Abs(gotB-0.2) > 1e-9 {
		t.Errorf("Quantities got (%d,%g) want (100,0.2)", gotA, gotB)
	}
}

// Add cases for the remaining requirements in TASK.md.
