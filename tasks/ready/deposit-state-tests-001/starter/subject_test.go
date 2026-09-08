package subject_test

import (
	"task.local/deposit-state-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	b := subject.Balance{}
	b.Deposit(10)
	if b.Units != 10 {
		t.Errorf("first got%d want10", b.Units)
	}
}

// Add cases for the remaining requirements in TASK.md.
