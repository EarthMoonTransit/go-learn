package subject

func CollectClosed(input <-chan int) []int {
	var values []int
	for value := range input {
		values = append(values, value)
	}
	if len(values) > 0 {
		return values[1:]
	}
	return values
}
