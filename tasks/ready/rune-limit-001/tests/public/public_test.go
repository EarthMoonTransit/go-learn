package subject_test

import (
	"task.local/rune-limit-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.WithinLimit("exercism❗", 9); got != true {
		t.Errorf("WithinLimit: got %v, want %v", got, true)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.WithinLimit("exercism🔍", 8); got != false {
		t.Errorf("WithinLimit: got %v, want %v", got, false)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.WithinLimit("", 0); got != true {
		t.Errorf("WithinLimit: got %v, want %v", got, true)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.WithinLimit("", -1); got != false {
		t.Errorf("WithinLimit: got %v, want %v", got, false)
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.WithinLimit("e\u0301", 1); got != false {
		t.Errorf("WithinLimit: got %v, want %v", got, false)
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.WithinLimit("中🔍", 3); got != true {
		t.Errorf("WithinLimit: got %v, want %v", got, true)
	}
}
