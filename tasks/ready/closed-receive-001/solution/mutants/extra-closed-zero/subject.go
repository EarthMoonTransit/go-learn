package subject

func CollectClosed(input <-chan int) []int {
	var values []int
	for value := range input {
		values = append(values, value)
	}
	return append(values, 0)
}
