package subject_test

import "testing"

func TestFiniteMergeExtended(t *testing.T) {
	checkBufferedMerge(t, []int{7, 0, 7}, nil)
	checkBufferedMerge(t, []int{104, 100, 102}, []int{204, 200, 202})
	checkBufferedMerge(t, []int{0, 0, 0}, []int{0, 0})
}
