package subject_test

import (
	"task.local/concurrent-counter-read-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	c := subject.NewCounter()
	for i := 1; i <= 8; i++ {
		c.Inc()
		if c.Value() != i {
			t.Errorf("fresh read got%d want%d", c.Value(), i)
		}
	}
}
