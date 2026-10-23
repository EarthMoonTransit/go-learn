package subject

func Window(input <-chan int, limit int) <-chan int {
	output := make(chan int, limit-1)
	go func() {
		defer close(output)
		for value := range input {
			if value != 0 {
				output <- value
			}
		}
	}()
	return output
}
