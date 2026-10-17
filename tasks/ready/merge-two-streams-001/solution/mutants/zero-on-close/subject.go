package subject

func MergeTwo(left, right <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for left != nil || right != nil {
			select {
			case value, open := <-left:
				if !open {
					out <- 0
					left = nil
					continue
				}
				out <- value
			case value, open := <-right:
				if !open {
					out <- 0
					right = nil
					continue
				}
				out <- value
			}
		}
	}()
	return out
}
