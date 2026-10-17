package subject

func ReplacePending(ch chan int, next int) int {
	old := <-ch
	ch <- next
	close(ch)
	return old
}
