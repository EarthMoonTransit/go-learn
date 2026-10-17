package subject

func ReplacePending(ch chan int, next int) int {
	old := <-ch
	ch <- old
	return old
}
