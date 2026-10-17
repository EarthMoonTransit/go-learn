package subject

func ReplacePending(ch chan int, next int) int {
	old := <-ch
	ch <- next
	return old
}
