package subject_test

import (
	"task.local/lexical-choice-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.ChooseLabel("Beta", "Delta"); got != "selected: Beta" {
		t.Errorf("result got %q want %q", got, "selected: Beta")
	}
}

func TestContractCase02(t *testing.T) {
	if got := subject.ChooseLabel("Delta", "Beta"); got != "selected: Beta" {
		t.Errorf("result got %q want %q", got, "selected: Beta")
	}
}

func TestContractCase03(t *testing.T) {
	if got := subject.ChooseLabel("abz", "aba"); got != "selected: aba" {
		t.Errorf("result got %q want %q", got, "selected: aba")
	}
}

func TestContractCase04(t *testing.T) {
	if got := subject.ChooseLabel("z", "A"); got != "selected: A" {
		t.Errorf("result got %q want %q", got, "selected: A")
	}
}

func TestContractCase05(t *testing.T) {
	if got := subject.ChooseLabel("x", "x"); got != "selected: x" {
		t.Errorf("result got %q want %q", got, "selected: x")
	}
}

func TestContractCase06(t *testing.T) {
	if got := subject.ChooseLabel("x", ""); got != "selected: " {
		t.Errorf("result got %q want %q", got, "selected: ")
	}
}

func TestContractCase07(t *testing.T) {
	if got := subject.ChooseLabel(" a", "a"); got != "selected:  a" {
		t.Errorf("result got %q want %q", got, "selected:  a")
	}
}

func TestContractCase08(t *testing.T) {
	if got := subject.ChooseLabel("a", "Z"); got != "selected: Z" {
		t.Errorf("case-sensitive choice got %q want selected: Z", got)
	}
}
