package subject

func Transform(input <-chan int, delta int) <-chan int {
	middle := make(chan int)
	output := make(chan int)
	go func() {
		defer close(middle)
		for v := range input {
			middle <- v + delta
		}
	}()
	go func() {
		defer close(output)
		for v := range middle {
			if v-delta >= 0 {
				output <- v
			}
		}
	}()
	return output
}
