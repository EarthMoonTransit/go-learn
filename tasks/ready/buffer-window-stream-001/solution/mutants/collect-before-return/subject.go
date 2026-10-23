package subject

func Window(input <-chan int, limit int) <-chan int {
	var values []int
	for value := range input {
		values = append(values, value)
	}
	output := make(chan int, len(values))
	for _, value := range values {
		output <- value
	}
	close(output)
	return output
}
