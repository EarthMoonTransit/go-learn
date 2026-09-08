package subject_test

import (
	"task.local/stateful-counter-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	f := subject.NewCounter(10)
	if f == nil {
		t.Fatal("nil closure")
	}
	if got := f(2); got != 12 {
		t.Errorf("first got%d want12", got)
	}
}

func TestContractCase02(t *testing.T) {
	a := subject.NewCounter(5)
	b := subject.NewCounter(5)
	c := subject.NewCounter(-2)
	if a == nil || b == nil || c == nil {
		t.Fatal("nil closure")
	}
	for _, x := range []struct {
		f    func(int) int
		d, w int
	}{{a, 3, 8}, {b, 0, 5}, {a, -4, 4}, {c, 1, -1}, {b, 2, 7}, {a, 0, 4}} {
		if got := x.f(x.d); got != x.w {
			t.Errorf("counter delta%d got%d want%d", x.d, got, x.w)
		}
	}
}
