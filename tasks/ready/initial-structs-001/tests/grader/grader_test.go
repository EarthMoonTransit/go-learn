package subject_test

import (
	"task.local/initial-structs-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.NewTrack(360); got.Distance != 360 {
		t.Errorf("track got %+v want distance360", got)
	}
}
