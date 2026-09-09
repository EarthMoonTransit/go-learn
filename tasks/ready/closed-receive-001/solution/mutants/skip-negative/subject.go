package subject

func CollectClosed(input <-chan int) []int {
	var values []int
	for value := range input {
		if value >= 0 {
			values = append(values, value)
		}
	}
	return values
}
