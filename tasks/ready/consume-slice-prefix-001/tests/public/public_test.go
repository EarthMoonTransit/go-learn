package subject_test

import (
	"fmt"
	"reflect"
	subject "task.local/consume-slice-prefix-001/subject"
	"testing"
)

func TestTakePrefixHeaders(t *testing.T) {
	for _, n := range []int{0, 1, 3, 4} {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			storage := []int{7, -2, 11, 5, 41, 43}
			input := storage[:4]
			saved := input
			got, ok := subject.TakePrefix(&input, n)
			if !ok || len(got) != n || cap(got) != 6 || len(input) != 4-n || cap(input) != 6-n {
				t.Fatalf("n=%d: ok=%t prefix len/cap=%d/%d remaining=%d/%d; want true %d/6 %d/%d", n, ok, len(got), cap(got), len(input), cap(input), n, 4-n, 6-n)
			}
			if !reflect.DeepEqual(got, storage[:n]) || !reflect.DeepEqual(input, storage[n:4]) || !reflect.DeepEqual(storage, []int{7, -2, 11, 5, 41, 43}) {
				t.Fatalf("n=%d: prefix=%v remaining=%v backing=%v; values changed or wrong window", n, got, input, storage)
			}
			if len(saved) != 4 || cap(saved) != 6 {
				t.Fatalf("copied header changed: len/cap=%d/%d; want 4/6", len(saved), cap(saved))
			}
			if n > 0 {
				got[0] = 89
				if saved[0] != 89 {
					t.Fatalf("prefix write invisible through saved view: got %d; want 89", saved[0])
				}
				saved[n-1] = -79
				if got[n-1] != -79 {
					t.Fatalf("saved write invisible through prefix: got %d; want -79", got[n-1])
				}
			}
			if len(input) > 0 {
				input[0] = 97
				if saved[n] != 97 {
					t.Fatalf("remaining view has wrong storage: saved[%d]=%d; want 97", n, saved[n])
				}
			}
		})
	}
}
func TestTakePrefixInvalidAndEmpty(t *testing.T) {
	for _, n := range []int{-1, 5, 6, int(^uint(0) >> 1)} {
		storage := []int{3, 5, 7, 11, 13, 17}
		input := storage[:4]
		got, ok := subject.TakePrefix(&input, n)
		if ok || got != nil || len(input) != 4 || cap(input) != 6 || &input[0] != &storage[0] || !reflect.DeepEqual(storage, []int{3, 5, 7, 11, 13, 17}) {
			t.Fatalf("invalid n=%d: prefix=%v ok=%t remaining=%v cap=%d backing=%v shared=%t; want nil,false and unchanged shared input", n, got, ok, input, cap(input), storage, len(input) > 0 && &input[0] == &storage[0])
		}
	}
	var input []int
	got, ok := subject.TakePrefix(&input, 0)
	if !ok || got != nil || input != nil {
		t.Fatalf("nil slice n=0: prefix=%v nil=%t remaining nil=%t ok=%t; want nil,nil,true", got, got == nil, input == nil, ok)
	}
	storage := []int{31, 37}
	input = storage[:0]
	got, ok = subject.TakePrefix(&input, 0)
	if !ok || len(got) != 0 || cap(got) != 2 || len(input) != 0 || cap(input) != 2 {
		t.Fatalf("empty with spare: prefix len/cap=%d/%d remaining=%d/%d ok=%t; want 0/2 0/2 true", len(got), cap(got), len(input), cap(input), ok)
	}
}
func TestTakePrefixSuccessiveCalls(t *testing.T) {
	input := []int{2, 4, 6, 8, 10}
	original := input
	a, okA := subject.TakePrefix(&input, 2)
	b, okB := subject.TakePrefix(&input, 1)
	if !okA || !okB || !reflect.DeepEqual(a, []int{2, 4}) || !reflect.DeepEqual(b, []int{6}) || !reflect.DeepEqual(input, []int{8, 10}) || cap(b) != 3 || cap(input) != 2 {
		t.Fatalf("successive prefixes=%v,%v remaining=%v capacities=%d,%d; want [2 4],[6],[8 10],3,2", a, b, input, cap(b), cap(input))
	}
	if len(original) != 5 || cap(original) != 5 {
		t.Fatal("original copied header changed")
	}
}
