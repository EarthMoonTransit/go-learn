package subject

func TryReceive(input <-chan int, cancel <-chan struct{}) (int, string) {
	select {
	case <-cancel:
		return 0, "canceled"
	case value, ok := <-input:
		_ = ok
		return value, "value"
	default:
		return 0, "pending"
	}
}
