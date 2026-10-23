package subject

func Window(input <-chan int, limit int) <-chan int {
	output := make(chan int, limit-1)
	go func() {
		defer close(output)
		for first := range input {
			second, ok := <-input
			if ok {
				output <- second
			}
			output <- first
		}
	}()
	return output
}
