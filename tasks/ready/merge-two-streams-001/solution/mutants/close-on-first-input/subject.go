package subject

func MergeTwo(left, right <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for left != nil || right != nil {
			select {
			case value, open := <-left:
				if !open {
					return
				}
				out <- value
			case value, open := <-right:
				if !open {
					return
				}
				out <- value
			}
		}
	}()
	return out
}
