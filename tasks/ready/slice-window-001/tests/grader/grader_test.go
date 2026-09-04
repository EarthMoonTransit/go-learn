package subject_test

import (
	"reflect"
	"task.local/slice-window-001/subject"
	"testing"
)

func TestHeadersAndSharing(t *testing.T) {
	for _, n := range []int{0, 1, 4} {
		for _, capacity := range []int{n, n + 3} {
			for low := -1; low <= n+1; low++ {
				for high := -1; high <= capacity+1; high++ {
					input := make([]int, n, capacity)
					for i := range input {
						input[i] = i + 11
					}
					before := append([]int{}, input...)
					oldLen, oldCap := len(input), cap(input)
					out, ok := subject.Window(input, low, high)
					valid := low >= 0 && high >= low && high <= n
					if ok != valid {
						t.Fatalf("invalid success flag")
					}
					if !reflect.DeepEqual(append([]int{}, input...), before) || len(input) != oldLen || cap(input) != oldCap {
						t.Fatal("input changed")
					}
					if !valid {
						if out != nil {
							t.Fatal("invalid window must be nil")
						}
						continue
					}
					if len(out) != high-low || cap(out) != capacity-low {
						t.Fatal("wrong slice header")
					}
					if len(out) > 0 {
						out[0] = 777
						if input[low] != 777 {
							t.Fatal("result must share storage")
						}
						input[low] = 888
						if out[0] != 888 {
							t.Fatal("sharing must work both ways")
						}
					}
				}
			}
		}
	}
	out, ok := subject.Window(nil, 0, 0)
	if !ok || out != nil {
		t.Fatal("nil empty window")
	}
}
