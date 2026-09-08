package subject_test

import (
	"task.local/exact-dynamic-number-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.Extract(subject.FancyNumber{Text: "-007"}); got != -7 {
		t.Errorf("signed decimal got %d want-7", got)
	}
}
