package subject_test

import (
	"task.local/independent-policies-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	a, b := subject.Classify("beta", 19)
	if !a || !b {
		t.Errorf("accepted got%v,%v", a, b)
	}
}
