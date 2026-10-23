package subject_test

import (
	"task.local/selected-window-copies-001/subject"
	"testing"
)

func TestFiniteWindowExtensions(t *testing.T) {
	for size := 1; size <= 8; size++ {
		input := make([]int, size)
		for i := range input {
			input[i] = i*3 - size
		}
		for start := 0; start < size; start++ {
			windowsCheck(t, input, []subject.Window{{start, size}, {start, start}, {0, start}, {start, size}})
		}
	}
	long := make([]int, 256)
	for i := range long {
		long[i] = i - 100
	}
	windowsCheck(t, long, []subject.Window{{128, 130}, {129, 131}, {0, 1}, {255, 256}, {100, 100}, {42, 44}, {128, 130}, {12, 13}})
}
