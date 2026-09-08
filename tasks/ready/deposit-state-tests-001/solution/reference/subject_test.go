package subject_test

import (
	"task.local/deposit-state-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	b := subject.Balance{}
	b.Deposit(10)
	if b.Units != 10 {
		t.Errorf("first got%d want10", b.Units)
	}
}

func TestContractCase02(t *testing.T) {
	b := subject.Balance{Units: 20}
	alias := &b
	other := subject.Balance{Units: 20}
	for _, amount := range []int{3, 0, 7} {
		before := b.Units
		b.Deposit(amount)
		if b.Units != before+amount || alias.Units != before+amount || other.Units != 20 {
			t.Errorf("deposit%d got%d alias%d other%d", amount, b.Units, alias.Units, other.Units)
		}
	}
	other.Deposit(1)
	if other.Units != 21 || b.Units != 30 {
		t.Errorf("independent got%d,%d", b.Units, other.Units)
	}
}
