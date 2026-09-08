package subject_test

import (
	"task.local/shared-counter-update-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	n := 2
	subject.AddCounter(&n, 2)
	if n != 4 {
		t.Errorf("stored value got %d want 4", n)
	}
}

func TestPublicCase02(t *testing.T) {
	n := 5
	p := &n
	q := p
	subject.AddCounter(p, 7)
	if *q != 12 {
		t.Errorf("alias got %d want 12", *q)
	}
}

func TestPublicCase03(t *testing.T) {
	n := 4
	subject.AddCounter(&n, 0)
	if n != 4 {
		t.Errorf("zero delta got %d want 4", n)
	}
}

func TestPublicCase04(t *testing.T) {
	n := 4
	subject.AddCounter(&n, -6)
	if n != -2 {
		t.Errorf("negative delta got %d want -2", n)
	}
}

func TestPublicCase05(t *testing.T) {
	n := 1
	subject.AddCounter(&n, 2)
	subject.AddCounter(&n, 3)
	if n != 6 {
		t.Errorf("repeated update got %d want 6", n)
	}
}
