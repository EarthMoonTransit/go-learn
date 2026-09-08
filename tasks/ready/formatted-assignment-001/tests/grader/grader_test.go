package subject_test

import (
	"task.local/formatted-assignment-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.Welcome(""); got != "Welcome, !" {
		t.Errorf("empty welcome got%q", got)
	}
}
