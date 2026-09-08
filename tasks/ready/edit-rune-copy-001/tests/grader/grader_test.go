package subject_test

import (
	"task.local/edit-rune-copy-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if a, b := subject.EditRune("e\u0301", 1, '!'); a != "e!" || b != true {
		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "e!", true)
	}
}

func TestGraderCase02(t *testing.T) {
	if a, b := subject.EditRune("🔍", 0, 0); a != "\x00" || b != true {
		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "\x00", true)
	}
}
