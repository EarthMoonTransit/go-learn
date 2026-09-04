package subject_test

import (
	"reflect"
	"task.local/append-value-001/subject"
	"testing"
)

func TestGrowthAndPreservation(t *testing.T) {
	for _, n := range []int{0, 1, 4, 100} {
		for _, capacity := range []int{n, n + 7} {
			for _, v := range []int{-1000, 0, 1000} {
				input := make([]int, n, capacity)
				for i := range input {
					input[i] = i - 50
				}
				before := append([]int{}, input...)
				want := append(append([]int{}, input...), v)
				got := subject.Push(input, v)
				if !reflect.DeepEqual(got, want) || len(input) != n || !reflect.DeepEqual(append([]int{}, input...), before) {
					t.Fatal("append contract")
				}
			}
		}
	}
	if got := subject.Push(nil, 3); !reflect.DeepEqual(got, []int{3}) {
		t.Fatal("nil append")
	}
}
