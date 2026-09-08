package subject_test

import (
	"task.local/checked-withdrawal-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	b := subject.Balance{Units: 1000000}
	if e := b.Withdraw(999999); e != nil || b.Units != 1 {
		t.Errorf("large withdrawal got%d,%v want1,nil", b.Units, e)
	}
}
