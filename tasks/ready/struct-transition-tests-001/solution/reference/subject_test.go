package subject_test

import (
	"task.local/struct-transition-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	m := subject.Motor{Speed: 5, Drain: 2, Battery: 100}
	want := subject.Motor{Speed: 5, Drain: 2, Battery: 98, Distance: 5}
	if got := subject.Advance(m); got != want {
		t.Errorf("advance got %+v want %+v", got, want)
	}
}

func TestContractCase02(t *testing.T) {
	m := subject.Motor{Speed: 18, Drain: 10, Battery: 10}
	want := subject.Motor{Speed: 18, Drain: 10, Battery: 0, Distance: 18}
	if got := subject.Advance(m); got != want {
		t.Errorf("exact battery got %+v want %+v", got, want)
	}
}

func TestContractCase03(t *testing.T) {
	m := subject.Motor{Speed: 5, Drain: 7, Battery: 3, Distance: 9}
	if got := subject.Advance(m); got != m {
		t.Errorf("insufficient got %+v want %+v", got, m)
	}
}

func TestContractCase04(t *testing.T) {
	m := subject.Motor{Speed: 5, Drain: 3, Battery: 100, Distance: 1}
	before := m
	want := m
	want.Battery = 97
	want.Distance = 6
	if got := subject.Advance(m); got != want {
		t.Errorf("nonzero distance got %+v want %+v", got, want)
	}
	if m != before {
		t.Errorf("caller value changed from %+v to %+v", before, m)
	}
}

func TestContractCase05(t *testing.T) {
	m := subject.Motor{Speed: 2, Drain: 1, Battery: 0, Distance: 5}
	if got := subject.Advance(m); got != m {
		t.Errorf("empty battery got %+v want %+v", got, m)
	}
}
