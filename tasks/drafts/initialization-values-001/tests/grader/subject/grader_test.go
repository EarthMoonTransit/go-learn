package subject_test

import (
	"task.local/initialization-values-001/subject"
	"testing"
)

func TestGraderBoundary(t *testing.T) {
	if a, b := subject.Values(); a != 3 || b != 6 {
		t.Errorf("package values got%d,%d", a, b)
	}
}
