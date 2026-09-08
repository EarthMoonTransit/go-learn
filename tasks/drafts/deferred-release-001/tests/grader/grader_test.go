package subject_test

import (
	"task.local/deferred-release-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	calls := 0
	n := subject.WithRelease(func() int { return -4 }, func() { calls++ })
	if n != -4 || calls != 1 {
		t.Errorf("negative got%d releases%d", n, calls)
	}
}
