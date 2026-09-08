package subject_test

import (
	"slices"
	"task.local/owned-values-001/subject"
	"testing"
	"testing/synctest"
)

func TestPublicCase01(t *testing.T) {
	input := []int{1, 0, -2}
	c := subject.OwnedValues(input)
	if c == nil {
		t.Fatal("nil stream")
	}
	if !slices.Equal(input, []int{1, 0, -2}) {
		t.Error("input changed")
	}
	input[0] = 9
	for _, want := range []int{1, 0, -2} {
		select {
		case got, ok := <-c:
			if !ok || got != want {
				t.Errorf("queued got%d,%v want%d,true", got, ok, want)
			}
		default:
			t.Fatal("stream value must be ready at return")
		}
	}
	select {
	case _, ok := <-c:
		if ok {
			t.Error("stream not closed after values")
		}
	default:
		t.Error("closure must be ready at return")
	}
}

func TestPublicCase02(t *testing.T) {
	for _, v := range [][]int{nil, {}} {
		c := subject.OwnedValues(v)
		if c == nil {
			t.Fatal("nil empty stream")
		}
		select {
		case _, ok := <-c:
			if ok {
				t.Error("empty stream sent value")
			}
		default:
			t.Error("empty stream not closed")
		}
	}
	a := subject.OwnedValues([]int{3})
	b := subject.OwnedValues([]int{4})
	if a == b {
		t.Error("streams share channel")
	}
	if <-a != 3 || <-b != 4 {
		t.Error("streams share values")
	}
}

func TestUnusedStreamRetainsNoWorker(t *testing.T) {
	synctest.Test(t, func(t *testing.T) { _ = subject.OwnedValues([]int{1, 2, 3}); synctest.Wait() })
}
