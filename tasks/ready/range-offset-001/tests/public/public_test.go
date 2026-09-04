package subject_test

import (
	"task.local/range-offset-001/subject"
	"testing"
)

func TestVisibleUpdates(t *testing.T) {
	values := []int{5, 5, -2}
	subject.OffsetByIndex(values)
	for i, w := range []int{6, 7, 1} {
		if values[i] != w {
			t.Errorf("element %d=%d, want %d", i, values[i], w)
		}
	}
	base := []int{91, 0, 0, 92}
	subject.OffsetByIndex(base[1:3])
	for i, w := range []int{91, 1, 2, 92} {
		if base[i] != w {
			t.Errorf("window element %d=%d, want %d", i, base[i], w)
		}
	}
}
