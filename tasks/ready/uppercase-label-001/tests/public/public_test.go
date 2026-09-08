package subject_test

import (
	"task.local/uppercase-label-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.UpperLabel("Judy"); got != "INPUT: JUDY" {
		t.Errorf("result got %q want %q", got, "INPUT: JUDY")
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.UpperLabel("lars"); got != "INPUT: LARS" {
		t.Errorf("result got %q want %q", got, "INPUT: LARS")
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.UpperLabel("Peter-James"); got != "INPUT: PETER-JAMES" {
		t.Errorf("result got %q want %q", got, "INPUT: PETER-JAMES")
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.UpperLabel("MJ"); got != "INPUT: MJ" {
		t.Errorf("result got %q want %q", got, "INPUT: MJ")
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.UpperLabel("éж9"); got != "INPUT: \u00c9\u04169" {
		t.Errorf("result got %q want %q", got, "INPUT: \u00c9\u04169")
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.UpperLabel(" a "); got != "INPUT:  A " {
		t.Errorf("result got %q want %q", got, "INPUT:  A ")
	}
}

func TestPublicCase07(t *testing.T) {
	if got := subject.UpperLabel(""); got != "INPUT: " {
		t.Errorf("result got %q want %q", got, "INPUT: ")
	}
}
