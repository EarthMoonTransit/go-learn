package subject_test

import (
	"task.local/wide-sum-001/subject"
	"testing"
)

func TestCompleteInt8Domain(t *testing.T) {
	for a := -128; a <= 127; a++ {
		for b := -128; b <= 127; b++ {
			if got := subject.WideSum(int8(a), int8(b)); got != a+b {
				t.Fatalf("WideSum(%d,%d)=%d, want %d", a, b, got, a+b)
			}
		}
	}
}
