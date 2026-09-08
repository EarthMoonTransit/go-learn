package subject_test

import (
	"task.local/decode-prefix-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if a, b, c := subject.DecodePrefix(""); a != 0 || b != 0 || c != false {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0, 0, false)
	}
}

func TestPublicCase02(t *testing.T) {
	if a, b, c := subject.DecodePrefix("A\xff"); a != 'A' || b != 1 || c != true {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 'A', 1, true)
	}
}

func TestPublicCase03(t *testing.T) {
	if a, b, c := subject.DecodePrefix("éx"); a != 'é' || b != 2 || c != true {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 'é', 2, true)
	}
}

func TestPublicCase04(t *testing.T) {
	if a, b, c := subject.DecodePrefix("🔍"); a != '🔍' || b != 4 || c != true {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, '🔍', 4, true)
	}
}

func TestPublicCase05(t *testing.T) {
	if a, b, c := subject.DecodePrefix("\xffx"); a != 0xfffd || b != 1 || c != false {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0xfffd, 1, false)
	}
}

func TestPublicCase06(t *testing.T) {
	if a, b, c := subject.DecodePrefix("�"); a != 0xfffd || b != 3 || c != true {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0xfffd, 3, true)
	}
}
