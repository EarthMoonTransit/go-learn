package subject_test

import (
	"task.local/division-results-001/subject"
	"testing"
)

func TestDivisionExamples(t *testing.T) {
	q, r, ok := subject.Divide(14, 4)
	if q != 3 || r != 2 || !ok {
		t.Fatalf("Divide(14,4)=(%d,%d,%v), want (3,2,true)", q, r, ok)
	}
	q, r, ok = subject.Divide(5, 0)
	if q != 0 || r != 0 || ok {
		t.Fatalf("zero divisor must return (0,0,false)")
	}
}
