package subject

func ReplacePending(ch chan int, next int) int {
	return <-ch
}
