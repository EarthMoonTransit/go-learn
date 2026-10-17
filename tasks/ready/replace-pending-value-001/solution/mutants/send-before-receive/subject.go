package subject

func ReplacePending(ch chan int, next int) int {
	ch <- next
	return <-ch
}
