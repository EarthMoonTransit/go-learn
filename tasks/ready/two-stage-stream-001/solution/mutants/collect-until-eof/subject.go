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
		values := []int{}
		for v := range middle {
			if v >= 0 {
				values = append(values, v)
			}
		}
		for _, v := range values {
			output <- v
		}
	}()
	return output
}
