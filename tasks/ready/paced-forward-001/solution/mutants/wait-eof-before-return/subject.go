package subject

func Pace(input <-chan int, release <-chan struct{}) <-chan int {
	values := []int{}
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
