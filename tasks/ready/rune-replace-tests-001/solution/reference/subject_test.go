package subject_test

import (
	"task.local/rune-replace-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.Replace("❗ recommended product", '❗', '?'); got != "? recommended product" {
		t.Errorf("Replace: got %q, want %q", got, "? recommended product")
	}
}

func TestContractCase02(t *testing.T) {
	if got := subject.Replace("❗x❗", '❗', '?'); got != "?x?" {
		t.Errorf("Replace: got %q, want %q", got, "?x?")
	}
}

func TestContractCase03(t *testing.T) {
	if got := subject.Replace("aéa", 'a', '🔍'); got != "🔍é🔍" {
		t.Errorf("Replace: got %q, want %q", got, "🔍é🔍")
	}
}

func TestContractCase04(t *testing.T) {
	if got := subject.Replace("", 'a', 'b'); got != "" {
		t.Errorf("Replace: got %q, want %q", got, "")
	}
}

func TestContractCase05(t *testing.T) {
	if got := subject.Replace("中", 'a', 'b'); got != "中" {
		t.Errorf("Replace: got %q, want %q", got, "中")
	}
}

func TestContractCase06(t *testing.T) {
	if got := subject.Replace("éé", 'é', 'é'); got != "éé" {
		t.Errorf("Replace: got %q, want %q", got, "éé")
	}
}
