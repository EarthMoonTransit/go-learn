package subject

func Window(input <-chan int, limit int) <-chan int {
	output := make(chan int, limit-1)
	go func() {

		for value := range input {
			output <- value
		}
	}()
	return output
}
