package subject

func Window(input <-chan int, limit int) <-chan int {
	output := make(chan int, limit)
	go func() {
		defer close(output)
		for value := range input {
			output <- value
		}
	}()
	return output
}
