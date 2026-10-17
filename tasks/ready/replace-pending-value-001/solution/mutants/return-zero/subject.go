package subject

func ReplacePending(ch chan int, next int) int {
	<-ch
	ch <- next
	return 0
}
