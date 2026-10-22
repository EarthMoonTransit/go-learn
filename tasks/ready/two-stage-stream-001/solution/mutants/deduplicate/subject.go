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
		seen := map[int]bool{}
		for v := range middle {
			if v >= 0 {
				if !seen[v] {
					output <- v
					seen[v] = true
				}
			}
		}
	}()
	return output
}
