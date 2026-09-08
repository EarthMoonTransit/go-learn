package subject_test

import (
	"task.local/decode-prefix-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if a, b, c := subject.DecodePrefix("\xc0\xaf"); a != 0xfffd || b != 1 || c != false {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0xfffd, 1, false)
	}
}

func TestGraderCase02(t *testing.T) {
	if a, b, c := subject.DecodePrefix("\xe2\x82"); a != 0xfffd || b != 1 || c != false {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0xfffd, 1, false)
	}
}

func TestGraderCase03(t *testing.T) {
	if a, b, c := subject.DecodePrefix("\xed\xa0\x80"); a != 0xfffd || b != 1 || c != false {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0xfffd, 1, false)
	}
}

func TestGraderCase04(t *testing.T) {
	if a, b, c := subject.DecodePrefix("\x00"); a != 0 || b != 1 || c != true {
		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0, 1, true)
	}
}
