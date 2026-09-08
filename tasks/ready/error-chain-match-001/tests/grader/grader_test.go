package subject_test

import (
	"errors"
	"task.local/error-chain-match-001/subject"
	"testing"
)

func TestGraderBoundary(t *testing.T) {
	a := errors.New("a")
	b := errors.New("b")
	if subject.Matches(a, b) {
		t.Error("distinct errors matched")
	}
}
