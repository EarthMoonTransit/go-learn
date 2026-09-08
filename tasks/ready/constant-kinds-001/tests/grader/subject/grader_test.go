package subject_test

import (
	"task.local/constant-kinds-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	a, b, c := subject.Values()
	if a != "localhost" || b != 8080 || c != 1.5 {
		t.Errorf("repeat got%q,%d,%g", a, b, c)
	}
}
