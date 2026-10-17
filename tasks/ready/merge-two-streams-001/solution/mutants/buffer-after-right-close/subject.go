package subject

func MergeTwo(left, right <-chan int) <-chan int {
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
					var rest []int
					if left != nil {
						for value := range left {
							rest = append(rest, value)
						}
					}
					for _, value := range rest {
						out <- value
					}
					return
				}
				out <- value
			}
		}
	}()
	return out
}
