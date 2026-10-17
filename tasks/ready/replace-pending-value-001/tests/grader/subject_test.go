package subject_test

import "testing"

func TestReplacementMatrix(t *testing.T) {
	for old := -4; old <= 4; old++ {
		for next := -3; next <= 3; next++ {
			checkReplacement(t, old, next)
		}
	}
}
