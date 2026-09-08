package subject_test

import (
	"task.local/rune-marker-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.Application("❗ recommended product"); got != "recommendation" {
		t.Errorf("Application: got %q, want %q", got, "recommendation")
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.Application("executed search 🔍"); got != "search" {
		t.Errorf("Application: got %q, want %q", got, "search")
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.Application("forecast: ☀ sunny"); got != "weather" {
		t.Errorf("Application: got %q, want %q", got, "weather")
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.Application("☀x❗🔍"); got != "weather" {
		t.Errorf("Application: got %q, want %q", got, "weather")
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.Application(""); got != "default" {
		t.Errorf("Application: got %q, want %q", got, "default")
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.Application("plain ⚠"); got != "default" {
		t.Errorf("Application: got %q, want %q", got, "default")
	}
}
