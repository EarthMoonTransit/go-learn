package subject_test

import (
	"io"
	"task.local/cleanup-order-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	r := &closeProbe{}
	t.Run("real-scope", func(child *testing.T) {
		got, e := subject.AcquireForTest(child, func() (io.Closer, error) { return r, nil })
		if got != r || e != nil || r.calls != 0 {
			t.Errorf("real setup got%v,%v closes%d", got, e, r.calls)
		}
	})
	if r.calls != 1 {
		t.Errorf("real cleanup closes%d", r.calls)
	}
}
