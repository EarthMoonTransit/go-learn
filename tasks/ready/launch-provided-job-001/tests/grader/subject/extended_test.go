package subject_test

import "testing"

func TestSingleAndMultipleJobs(t *testing.T) {
	for _, count := range []int{1, 4} {
		checkJobs(t, count)
	}
}
