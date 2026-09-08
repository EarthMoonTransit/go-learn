package subject_test

import (
	"task.local/detached-prefix-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	out := subject.DetachedPrefix(nil, 0)
	if len(out) != 0 || cap(out) != 0 {
		t.Errorf("nil got%v cap%d", out, cap(out))
	}
}
