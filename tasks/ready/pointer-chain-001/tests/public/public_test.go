package subject_test

import (
	"task.local/pointer-chain-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	p, pp := subject.PointerChain()
	if p == nil || pp == nil || *pp != p || *p != 20 || **pp != 20 {
		t.Fatalf("invalid chain p%v pp%v", p, pp)
	}
	**pp = 7
	if *p != 7 {
		t.Error("dereference paths disagree")
	}
	other := 33
	*pp = &other
	if *p != 7 || **pp != 33 {
		t.Errorf("pointer replacement changed target got%d,%d", *p, **pp)
	}
}

func TestPublicCase02(t *testing.T) {
	a, ap := subject.PointerChain()
	b, bp := subject.PointerChain()
	if a == nil || b == nil || ap == nil || bp == nil {
		t.Fatal("nil chain")
	}
	*a = 9
	if *b != 20 || *ap != a || *bp != b || ap == bp {
		t.Error("calls share pointer or value storage")
	}
}
