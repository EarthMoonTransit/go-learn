package subject_test

import (
	"task.local/period-predicate-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	f := subject.WithinDays(subject.DaysPeriod{From: -3, To: -1})
	if f == nil || !f(subject.Record{Day: -2}) {
		t.Error("negative day range rejected")
	}
}
