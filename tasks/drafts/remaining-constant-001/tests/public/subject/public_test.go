package subject_test

import (
	"task.local/remaining-constant-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	copied := subject.Target()
	if copied != 40 {
		t.Errorf("targetMinutes got %d want 40", copied)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.RemainingMinutes(15); got != 25 {
		t.Errorf("RemainingMinutes: got %v, want %v", got, 25)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.RemainingMinutes(30); got != 10 {
		t.Errorf("RemainingMinutes: got %v, want %v", got, 10)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.RemainingMinutes(0); got != 40 {
		t.Errorf("RemainingMinutes: got %v, want %v", got, 40)
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.RemainingMinutes(40); got != 0 {
		t.Errorf("RemainingMinutes: got %v, want %v", got, 0)
	}
}
