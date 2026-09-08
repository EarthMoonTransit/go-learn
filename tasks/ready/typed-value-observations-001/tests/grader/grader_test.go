package subject_test

import (
	"task.local/typed-value-observations-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	_, _, v, d := subject.LiteralValues()
	if v != true || d != float32(3.14) {
		t.Errorf("repeat got%v,%g", v, d)
	}
}
