package subject_test

import (
	"slices"
	"task.local/generic-stack-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	s := subject.NewStack[int]()
	if s == nil {
		t.Fatal("nil new stack")
	}
	if !s.IsEmpty() {
		t.Error("new nonempty")
	}
	s.Push(123)
	s.Push(456)
	if s.IsEmpty() {
		t.Error("populated empty")
	}
	for _, want := range []int{456, 123} {
		got, ok := s.Pop()
		if got != want || !ok {
			t.Errorf("LIFO got%d,%v want%d,true", got, ok, want)
		}
	}
	for i := 0; i < 2; i++ {
		v, ok := s.Pop()
		if v != 0 || ok || !s.IsEmpty() {
			t.Errorf("empty got%d,%v", v, ok)
		}
	}
	s.Push(7)
	v, ok := s.Pop()
	if v != 7 || !ok {
		t.Error("reuse failed")
	}
}

func TestPublicCase02(t *testing.T) {
	var a, b subject.Stack[string]
	a.Push("a")
	b.Push("b")
	a.Push("")
	for _, want := range []string{"", "a"} {
		v, ok := a.Pop()
		if v != want || !ok {
			t.Errorf("first instance got%q,%v", v, ok)
		}
	}
	v, ok := b.Pop()
	if v != "b" || !ok {
		t.Errorf("second instance got%q,%v", v, ok)
	}
}

func TestPublicCase03(t *testing.T) {
	s := subject.NewStack[[]int]()
	if s == nil {
		t.Fatal("nil stack")
	}
	s.Push([]int{1, 2})
	v, ok := s.Pop()
	if !ok || !slices.Equal(v, []int{1, 2}) {
		t.Errorf("slice element got%v,%v", v, ok)
	}
	v, ok = s.Pop()
	if v != nil || ok {
		t.Errorf("slice zero got%v,%v", v, ok)
	}
}
