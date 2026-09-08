package subject_test

import (
	"task.local/variadic-aliasing-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	v := []int{1}
	if !subject.ReplaceFirst(0, v...) || v[0] != 0 {
		t.Errorf("single zero got%v", v)
	}
}
