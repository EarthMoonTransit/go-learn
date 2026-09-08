package subject_test

import (
	"task.local/rune-boundary-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if a, b := subject.RuneBoundary("abc", 1); a != 1 || b != true {
		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 1, true)
	}
}

func TestContractCase02(t *testing.T) {
	if a, b := subject.RuneBoundary("Aé🔍中", 3); a != 7 || b != true {
		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 7, true)
	}
}

func TestContractCase03(t *testing.T) {
	if a, b := subject.RuneBoundary("é", 1); a != 2 || b != true {
		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 2, true)
	}
}

func TestContractCase04(t *testing.T) {
	if a, b := subject.RuneBoundary("", 0); a != 0 || b != true {
		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 0, true)
	}
}

func TestContractCase05(t *testing.T) {
	if a, b := subject.RuneBoundary("é", 2); a != 0 || b != false {
		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 0, false)
	}
}

func TestContractCase06(t *testing.T) {
	if a, b := subject.RuneBoundary("x", -1); a != 0 || b != false {
		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 0, false)
	}
}

func TestContractCase07(t *testing.T) {
	if a, b := subject.RuneBoundary("e\u0301", 2); a != 3 || b != true {
		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 3, true)
	}
}
