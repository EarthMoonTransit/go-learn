package subject_test

import (
	"task.local/wide-sum-001/subject"
	"testing"
)

func TestWideResults(t *testing.T) {
	for _, c := range []struct {
		a, b int8
		want int
	}{{100, 100, 200}, {-100, -100, -200}, {-8, 3, -5}} {
		if got := subject.WideSum(c.a, c.b); got != c.want {
			t.Errorf("WideSum(%d,%d)=%d, want %d", c.a, c.b, got, c.want)
		}
	}
}
