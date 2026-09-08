package subject_test

import (
	"task.local/external-window-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	w := subject.NewWindow(2, 5)
	if !w.Contains(3) {
		t.Error("interior missing")
	}
}

func TestContractCase02(t *testing.T) {
	for _, c := range []struct{ min, max int }{{2, 5}, {-3, 0}, {0, 0}, {7, 7}} {
		w := subject.NewWindow(c.min, c.max)
		for _, v := range []int{c.min - 1, c.min, c.max - 1, c.max, c.max + 1} {
			want := v >= c.min && v < c.max
			if got := w.Contains(v); got != want {
				t.Errorf("window[%d,%d) value%d got%v want%v", c.min, c.max, v, got, want)
			}
		}
	}
}

func TestContractCase03(t *testing.T) {
	a := subject.NewWindow(1, 3)
	b := subject.NewWindow(7, 9)
	if !a.Contains(2) || a.Contains(8) || !b.Contains(8) || b.Contains(2) {
		t.Error("independent windows changed")
	}
	empty := subject.NewWindow(4, 4)
	if empty.Contains(4) || !a.Contains(2) {
		t.Error("empty constructor changed existing window")
	}
}
