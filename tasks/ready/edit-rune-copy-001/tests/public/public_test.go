package subject_test

import (
	"task.local/edit-rune-copy-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if a, b := subject.EditRune("abc", 1, 'X'); a != "aXc" || b != true {
		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "aXc", true)
	}
}

func TestPublicCase02(t *testing.T) {
	if a, b := subject.EditRune("Aé🔍", 1, '中'); a != "A中🔍" || b != true {
		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "A中🔍", true)
	}
}

func TestPublicCase03(t *testing.T) {
	if a, b := subject.EditRune("éé", 1, 'a'); a != "éa" || b != true {
		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "éa", true)
	}
}

func TestPublicCase04(t *testing.T) {
	if a, b := subject.EditRune("", 0, 'x'); a != "" || b != false {
		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "", false)
	}
}

func TestPublicCase05(t *testing.T) {
	if a, b := subject.EditRune("é", 1, 'x'); a != "é" || b != false {
		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "é", false)
	}
}

func TestPublicCase06(t *testing.T) {
	if a, b := subject.EditRune("ab", -1, 'x'); a != "ab" || b != false {
		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "ab", false)
	}
}
