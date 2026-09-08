package subject_test

import (
	"task.local/named-numeric-constraint-001/subject"
	"testing"
)

func TestGraderBoundary(t *testing.T) {
	if got := subject.Double(-.125); got != -.25 {
		t.Errorf("fraction got%g want-.25", got)
	}
}
