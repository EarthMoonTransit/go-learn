package subject_test

import (
	"task.local/checked-withdrawal-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	original := subject.ErrInsufficient
	if original == nil {
		t.Fatal("nil sentinel")
	}
	defer func() {
		if subject.ErrInsufficient != original {
			t.Error("sentinel identity changed")
		}
	}()
	b := subject.Balance{Units: 20}
	alias := &b
	for _, amount := range []int{10, 10, 0} {
		before := b.Units
		e := b.Withdraw(amount)
		if e != nil || b.Units != before-amount || alias.Units != before-amount {
			t.Errorf("withdraw%d got%d,%v want%d,nil", amount, b.Units, e, before-amount)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	original := subject.ErrInsufficient
	if original == nil {
		t.Fatal("nil sentinel")
	}
	defer func() {
		if subject.ErrInsufficient != original {
			t.Error("sentinel identity changed")
		}
	}()
	b := subject.Balance{Units: 20}
	sentinel := subject.ErrInsufficient
	if sentinel == nil {
		t.Fatal("nil sentinel")
	}
	e := b.Withdraw(100)
	if e != sentinel || b.Units != 20 || subject.ErrInsufficient != sentinel {
		t.Errorf("failed got%d,%v want20,original sentinel", b.Units, e)
	}
	if e = b.Withdraw(1); e != nil || b.Units != 19 {
		t.Errorf("after failure got%d,%v want19,nil", b.Units, e)
	}
}

func TestPublicCase03(t *testing.T) {
	original := subject.ErrInsufficient
	if original == nil {
		t.Fatal("nil sentinel")
	}
	defer func() {
		if subject.ErrInsufficient != original {
			t.Error("sentinel identity changed")
		}
	}()
	b := subject.Balance{}
	sentinel := subject.ErrInsufficient
	if e := b.Withdraw(1); e != sentinel || b.Units != 0 {
		t.Errorf("zero balance got%d,%v", b.Units, e)
	}
}
