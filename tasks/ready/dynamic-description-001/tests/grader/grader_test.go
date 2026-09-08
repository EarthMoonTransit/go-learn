package subject_test

import (
	"task.local/dynamic-description-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.Describe(subject.FancyNumber{Text: "9999999999999999999999999999999"}); got != "fancy 0.0" {
		t.Errorf("overflow got %q want fancy0", got)
	}
}
