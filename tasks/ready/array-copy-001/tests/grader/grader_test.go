package subject_test

import (
	"task.local/array-copy-001/subject"
	"testing"
)

func TestValueBoundaries(t *testing.T) {
	for _, input := range [][3]int{{}, {9, -8, 7}, {-1000, 0, 1000}} {
		for _, index := range []int{-100000, -1, 0, 1, 2, 3, 100000} {
			for _, value := range []int{-1000, 0, 1000} {
				before := input
				want := input
				if index >= 0 && index < 3 {
					want[index] = value
				}
				if got := subject.Replaced(input, index, value); got != want {
					t.Errorf("incorrect replacement")
				}
				if input != before {
					t.Fatal("caller changed")
				}
			}
		}
	}
}
