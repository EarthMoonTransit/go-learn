package subject

func MergeTwo(left, right <-chan int) <-chan int {
	right = nil
	out := make(chan int)
	go func() {
		defer close(out)
		for left != nil || right != nil {
			select {
			case value, open := <-left:
				if !open {
					left = nil
					continue
				}
				out <- value
			case value, open := <-right:
				if !open {
					right = nil
					continue
				}
				out <- value
			}
		}
	}()
	return out
}
