package subject_test

import (
	"task.local/encode-scalar-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if a, b, c := subject.EncodeScalar('A'); a != "A" || b != 1 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "A", 1, true)
	}
}

func TestPublicCase02(t *testing.T) {
	if a, b, c := subject.EncodeScalar('é'); a != "é" || b != 2 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "é", 2, true)
	}
}

func TestPublicCase03(t *testing.T) {
	if a, b, c := subject.EncodeScalar('中'); a != "中" || b != 3 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "中", 3, true)
	}
}

func TestPublicCase04(t *testing.T) {
	if a, b, c := subject.EncodeScalar('🔍'); a != "🔍" || b != 4 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "🔍", 4, true)
	}
}

func TestPublicCase05(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0); a != "\x00" || b != 1 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "\x00", 1, true)
	}
}

func TestPublicCase06(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0xfffd); a != "�" || b != 3 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "�", 3, true)
	}
}

func TestPublicCase07(t *testing.T) {
	if a, b, c := subject.EncodeScalar(-1); a != "" || b != 0 || c != false {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "", 0, false)
	}
}

func TestPublicCase08(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0xd800); a != "" || b != 0 || c != false {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "", 0, false)
	}
}

func TestPublicCase09(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0x110000); a != "" || b != 0 || c != false {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "", 0, false)
	}
}
