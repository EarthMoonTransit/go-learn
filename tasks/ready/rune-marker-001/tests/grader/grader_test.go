package subject_test

import (
	"task.local/rune-marker-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.Application("é🔍❗☀"); got != "search" {
		t.Errorf("Application: got %q, want %q", got, "search")
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.Application("中❗☀"); got != "recommendation" {
		t.Errorf("Application: got %q, want %q", got, "recommendation")
	}
}

func TestGraderCase03(t *testing.T) {
	if got := subject.Application("☀☀"); got != "weather" {
		t.Errorf("Application: got %q, want %q", got, "weather")
	}
}
