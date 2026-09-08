package subject_test

import (
	"task.local/mutex-counter-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	c := subject.NewCounter()
	for range 7 {
		c.Inc()
	}
	if c.Value() != 7 {
		t.Error("later increments lost")
	}
}
