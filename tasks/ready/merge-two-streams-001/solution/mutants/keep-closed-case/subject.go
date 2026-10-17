package subject

func MergeTwo(left, right <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		remaining := 2
		for remaining > 0 {
			select {
			case value, open := <-left:
				if !open {
					remaining--
					continue
				}
				out <- value
			case value, open := <-right:
				if !open {
					remaining--
					continue
				}
				out <- value
			}
		}
	}()
	return out
}
