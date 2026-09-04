package subject_test

import (
	"task.local/division-results-001/subject"
	"testing"
)

func TestResultContract(t *testing.T) {
	for _, n := range []int{-1000, -17, -4, -1, 0, 1, 4, 17, 1000} {
		for _, d := range []int{-1000, -5, -1, 0, 1, 5, 1000} {
			q, r, ok := subject.Divide(n, d)
			if d == 0 {
				if q != 0 || r != 0 || ok {
					t.Errorf("zero-divisor result")
				}
				continue
			}
			if !ok || q != n/d || r != n%d {
				t.Errorf("incorrect result tuple for %d and %d", n, d)
			}
		}
	}
}
