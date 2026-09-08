package subject_test

import (
	"task.local/repair-utf8-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.RepairUTF8("\xed\xa0\x80"); got != "���" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "���")
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.RepairUTF8("\xc0\xaf"); got != "��" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "��")
	}
}

func TestGraderCase03(t *testing.T) {
	if got := subject.RepairUTF8("🔍\xffé"); got != "🔍�é" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "🔍�é")
	}
}
