package subject_test

import (
	"task.local/token-values-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.TokenValue("ace"); got != 11 {
		t.Errorf("TokenValue: got %v, want %v", got, 11)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.TokenValue("two"); got != 2 {
		t.Errorf("TokenValue: got %v, want %v", got, 2)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.TokenValue("three"); got != 3 {
		t.Errorf("TokenValue: got %v, want %v", got, 3)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.TokenValue("four"); got != 4 {
		t.Errorf("TokenValue: got %v, want %v", got, 4)
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.TokenValue("five"); got != 5 {
		t.Errorf("TokenValue: got %v, want %v", got, 5)
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.TokenValue("six"); got != 6 {
		t.Errorf("TokenValue: got %v, want %v", got, 6)
	}
}

func TestPublicCase07(t *testing.T) {
	if got := subject.TokenValue("seven"); got != 7 {
		t.Errorf("TokenValue: got %v, want %v", got, 7)
	}
}

func TestPublicCase08(t *testing.T) {
	if got := subject.TokenValue("eight"); got != 8 {
		t.Errorf("TokenValue: got %v, want %v", got, 8)
	}
}

func TestPublicCase09(t *testing.T) {
	if got := subject.TokenValue("nine"); got != 9 {
		t.Errorf("TokenValue: got %v, want %v", got, 9)
	}
}

func TestPublicCase10(t *testing.T) {
	if got := subject.TokenValue("ten"); got != 10 {
		t.Errorf("TokenValue: got %v, want %v", got, 10)
	}
}

func TestPublicCase11(t *testing.T) {
	if got := subject.TokenValue("jack"); got != 10 {
		t.Errorf("TokenValue: got %v, want %v", got, 10)
	}
}

func TestPublicCase12(t *testing.T) {
	if got := subject.TokenValue("queen"); got != 10 {
		t.Errorf("TokenValue: got %v, want %v", got, 10)
	}
}

func TestPublicCase13(t *testing.T) {
	if got := subject.TokenValue("king"); got != 10 {
		t.Errorf("TokenValue: got %v, want %v", got, 10)
	}
}

func TestPublicCase14(t *testing.T) {
	if got := subject.TokenValue("joker"); got != 0 {
		t.Errorf("TokenValue: got %v, want %v", got, 0)
	}
}

func TestPublicCase15(t *testing.T) {
	if got := subject.TokenValue("ACE"); got != 0 {
		t.Errorf("TokenValue: got %v, want %v", got, 0)
	}
}

func TestPublicCase16(t *testing.T) {
	if got := subject.TokenValue(""); got != 0 {
		t.Errorf("TokenValue: got %v, want %v", got, 0)
	}
}
