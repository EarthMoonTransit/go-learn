package subject

import (
	"slices"
)

func CollectClosed(input <-chan int) []int {
	var values []int
	for value := range input {
		values = append(values, value)
	}
	slices.Reverse(values)
	return values
}
