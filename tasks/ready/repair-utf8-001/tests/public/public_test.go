package subject_test

import (
	"task.local/repair-utf8-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.RepairUTF8("plain é"); got != "plain é" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "plain é")
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.RepairUTF8("A\xffB"); got != "A�B" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "A�B")
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.RepairUTF8("\xff\xfe"); got != "��" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "��")
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.RepairUTF8("�\x00"); got != "�\x00" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "�\x00")
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.RepairUTF8(""); got != "" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "")
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.RepairUTF8("\xe2\x82"); got != "��" {
		t.Errorf("RepairUTF8: got %q, want %q", got, "��")
	}
}
