package subject_test

import (
	subject "task.local/recursive-total-001/subject"
	"testing"
)

func TestAllAdmittedLengths(t *testing.T) {
	for n := 0; n <= 32; n++ {
		values := make([]int, n)
		want := 0
		for i := range values {
			values[i] = (i%5 - 2) * 100
			want += values[i]
		}
		before := append([]int(nil), values...)
		if got := subject.Total(values); got != want {
			t.Errorf("len=%d total=%d; want %d", n, got, want)
		}
		for i, v := range before {
			if values[i] != v {
				t.Errorf("len=%d input[%d]=%d; want %d", n, i, values[i], v)
			}
		}
	}
}
