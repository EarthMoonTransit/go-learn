package subject_test

import (
	"task.local/duration-rounding-001/subject"
	"testing"
	"time"
)

func TestGraderCase01(t *testing.T) {
	n, e := subject.WholeMinutes(23*time.Hour + 59*time.Minute)
	if n != 1439 || e != nil {
		t.Errorf("last whole minute got%d,%v", n, e)
	}
}
