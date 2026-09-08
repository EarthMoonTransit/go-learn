package subject_test

import (
	"task.local/value-average-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := (subject.Sample{Label: "other", Attempts: 1000000, Hits: 1000000}).Average(); got != 1 {
		t.Errorf("all hits got %g want1", got)
	}
}
