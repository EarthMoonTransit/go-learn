package subject_test

import (
	"task.local/typed-nil-result-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	r := &subject.Result{Code: 6}
	n, ok := subject.ReadCode(r)
	if n != 6 || !ok {
		t.Errorf("code got%d,%v want6,true", n, ok)
	}
}
