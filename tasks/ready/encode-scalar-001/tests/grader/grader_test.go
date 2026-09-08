package subject_test

import (
	"task.local/encode-scalar-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0x7ff); a != "\u07ff" || b != 2 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "\u07ff", 2, true)
	}
}

func TestGraderCase02(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0xd7ff); a != "\ud7ff" || b != 3 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "\ud7ff", 3, true)
	}
}

func TestGraderCase03(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0xe000); a != "\ue000" || b != 3 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "\ue000", 3, true)
	}
}

func TestGraderCase04(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0x10ffff); a != "\U0010ffff" || b != 4 || c != true {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "\U0010ffff", 4, true)
	}
}

func TestGraderCase05(t *testing.T) {
	if a, b, c := subject.EncodeScalar(0xdfff); a != "" || b != 0 || c != false {
		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "", 0, false)
	}
}
