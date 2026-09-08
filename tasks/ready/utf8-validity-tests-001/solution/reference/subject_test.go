package subject_test

import (
	"task.local/utf8-validity-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.ValidUTF8("ASCII"); got != true {
		t.Errorf("ValidUTF8: got %v, want %v", got, true)
	}
}

func TestContractCase02(t *testing.T) {
	if got := subject.ValidUTF8(""); got != true {
		t.Errorf("ValidUTF8: got %v, want %v", got, true)
	}
}

func TestContractCase03(t *testing.T) {
	if got := subject.ValidUTF8("\x00é中🔍�"); got != true {
		t.Errorf("ValidUTF8: got %v, want %v", got, true)
	}
}

func TestContractCase04(t *testing.T) {
	if got := subject.ValidUTF8("\x80"); got != false {
		t.Errorf("ValidUTF8: got %v, want %v", got, false)
	}
}

func TestContractCase05(t *testing.T) {
	if got := subject.ValidUTF8("\xff"); got != false {
		t.Errorf("ValidUTF8: got %v, want %v", got, false)
	}
}

func TestContractCase06(t *testing.T) {
	if got := subject.ValidUTF8("\xc0\xaf"); got != false {
		t.Errorf("ValidUTF8: got %v, want %v", got, false)
	}
}

func TestContractCase07(t *testing.T) {
	if got := subject.ValidUTF8("\xe2\x82"); got != false {
		t.Errorf("ValidUTF8: got %v, want %v", got, false)
	}
}

func TestContractCase08(t *testing.T) {
	if got := subject.ValidUTF8("\xed\xa0\x80"); got != false {
		t.Errorf("ValidUTF8: got %v, want %v", got, false)
	}
}

func TestContractCase09(t *testing.T) {
	if got := subject.ValidUTF8("\xf4\x90\x80\x80"); got != false {
		t.Errorf("ValidUTF8: got %v, want %v", got, false)
	}
}

func TestContractCase10(t *testing.T) {
	if got := subject.ValidUTF8("ok\xff"); got != false {
		t.Errorf("ValidUTF8: got %v, want %v", got, false)
	}
}
