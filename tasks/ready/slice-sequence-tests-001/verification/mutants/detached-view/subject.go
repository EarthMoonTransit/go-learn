package subject

import (
	"slices"
)

func Sequences() ([]int, []string, []string) {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}
	words := []string{"a", "b", "c", "d", "e"}
	return numbers, words, slices.Clone(words[1:3])
}
