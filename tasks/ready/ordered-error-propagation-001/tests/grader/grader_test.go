package subject_test

import (
	"math"
	"task.local/ordered-error-propagation-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	c := &calculator{amount: -8, factor: .5}
	n, e := subject.Divide(c, 2)
	if e != nil || math.IsNaN(n) || math.Abs(n+2) > 1e-9 {
		t.Errorf("signed result got %g,%v want-2,nil", n, e)
	}
}
