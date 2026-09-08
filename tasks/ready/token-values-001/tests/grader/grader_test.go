package subject_test

import (
	"task.local/token-values-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.TokenValue("ten "); got != 0 {
		t.Errorf("TokenValue: got %v, want %v", got, 0)
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.TokenValue("1"); got != 0 {
		t.Errorf("TokenValue: got %v, want %v", got, 0)
	}
}
