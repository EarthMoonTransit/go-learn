package subject_test

import (
	"task.local/text-border-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.AddBorder("é", 3); got != "***\n\u00e9\n***" {
		t.Errorf("result got %q want %q", got, "***\n\u00e9\n***")
	}
}
