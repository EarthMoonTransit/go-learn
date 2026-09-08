package subject_test

import (
	"task.local/select-ready-case-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if n, s := subject.TryReceive(nil, nil); n != 0 || s != "pending" {
		t.Errorf("disabled got%d,%s", n, s)
	}
	open := make(chan int)
	if n, s := subject.TryReceive(open, nil); n != 0 || s != "pending" {
		t.Errorf("unready got%d,%s", n, s)
	}
	close(open)
	if n, s := subject.TryReceive(open, nil); n != 0 || s != "closed" {
		t.Errorf("closed got%d,%s", n, s)
	}
}

func TestPublicCase02(t *testing.T) {
	c := make(chan int, 2)
	c <- 0
	c <- -3
	close(c)
	for _, want := range []int{0, -3} {
		if n, s := subject.TryReceive(c, nil); n != want || s != "value" {
			t.Errorf("value got%d,%s want%d,value", n, s, want)
		}
	}
	stop := make(chan struct{})
	close(stop)
	if n, s := subject.TryReceive(nil, stop); n != 0 || s != "canceled" {
		t.Errorf("cancel got%d,%s", n, s)
	}
}

func TestPublicCase03(t *testing.T) {
	c := make(chan int, 1)
	c <- 7
	stop := make(chan struct{})
	close(stop)
	n, s := subject.TryReceive(c, stop)
	if !(s == "value" && n == 7 || s == "canceled" && n == 0) {
		t.Errorf("both ready got%d,%s", n, s)
	}
	if s == "canceled" && len(c) != 1 {
		t.Error("unselected input consumed")
	}
	if s == "value" && len(c) != 0 {
		t.Error("selected input not consumed")
	}
}
