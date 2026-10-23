package subject

func Window(input <-chan int, limit int) <-chan int {
	output := make(chan int, max(0, limit-2))
	go func() {
		defer close(output)
		for value := range input {
			output <- value
		}
	}()
	return output
}
