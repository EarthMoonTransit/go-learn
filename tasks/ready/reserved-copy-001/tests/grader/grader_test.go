package subject_test

import (
	"task.local/reserved-copy-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	in := []int{1}
	out := subject.ReservedCopy(in, 0)
	if len(out) != 1 || cap(out) != 1 || out[0] != 1 {
		t.Fatalf("zero extra got%v cap%d", out, cap(out))
	}
	out[0] = 7
	if in[0] != 1 {
		t.Error("zero extra shares input")
	}
}
