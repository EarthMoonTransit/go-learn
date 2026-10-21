package subject_test

import (
	"reflect"
	subject "task.local/consume-slice-prefix-001/subject"
	"testing"
)

func TestPrefixEveryBoundary(t *testing.T) {
	for size := 0; size <= 6; size++ {
		for n := -2; n <= size+2; n++ {
			storage := []int{2, 3, 5, 7, 11, 13, 17, 19}
			input := storage[:size]
			before := append([]int(nil), storage...)
			got, ok := subject.TakePrefix(&input, n)
			valid := n >= 0 && n <= size
			if ok != valid {
				t.Fatalf("size=%d n=%d: ok=%t; want %t", size, n, ok, valid)
			}
			if !reflect.DeepEqual(storage, before) {
				t.Fatalf("size=%d n=%d changed backing: %v; want %v", size, n, storage, before)
			}
			if !valid {
				if got != nil || len(input) != size || cap(input) != 8 {
					t.Fatalf("invalid n=%d changed header", n)
				}
				continue
			}
			if !reflect.DeepEqual(got, before[:n]) || !reflect.DeepEqual(input, before[n:size]) || cap(got) != 8 || cap(input) != 8-n {
				t.Fatalf("size=%d n=%d wrong headers: prefix=%v cap=%d remaining=%v cap=%d", size, n, got, cap(got), input, cap(input))
			}
		}
	}
}
