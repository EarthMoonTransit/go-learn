package subject_test

import (
	"task.local/allocated-counter-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	p := subject.NewCounter(2)
	if p == nil {
		t.Fatal("NewCounter(2): got nil want pointer")
	}
	if *p != 2 {
		t.Errorf("counter got %d want 2", *p)
	}
}

func TestPublicCase02(t *testing.T) {
	p := subject.NewCounter(0)
	if p == nil {
		t.Fatal("zero counter: got nil want pointer")
	}
	if *p != 0 {
		t.Errorf("counter got %d want 0", *p)
	}
}

func TestPublicCase03(t *testing.T) {
	p := subject.NewCounter(-3)
	if p == nil {
		t.Fatal("negative counter: got nil want pointer")
	}
	if *p != -3 {
		t.Errorf("counter got %d want -3", *p)
	}
}

func TestPublicCase04(t *testing.T) {
	a, b := subject.NewCounter(4), subject.NewCounter(4)
	if a == nil || b == nil {
		t.Fatal("got nil counters")
	}
	*a = 7
	if *a != 7 || *b != 4 {
		t.Errorf("independent counters got (%d,%d) want (7,4)", *a, *b)
	}
}
