package subject_test

import (
	"task.local/rune-limit-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.WithinLimit("\x00", 1); got != true {
		t.Errorf("WithinLimit: got %v, want %v", got, true)
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.WithinLimit("☀", 0); got != false {
		t.Errorf("WithinLimit: got %v, want %v", got, false)
	}
}

func TestGraderCase03(t *testing.T) {
	if got := subject.WithinLimit("é", 1); got != true {
		t.Errorf("WithinLimit: got %v, want %v", got, true)
	}
}
