package subject

func TryReceive(input <-chan int, cancel <-chan struct{}) (int, string) {
	select {
	case _, ok := <-cancel:
		if ok {
			return 0, "pending"
		}
		return 0, "canceled"
	case value, ok := <-input:
		if !ok {
			return 0, "closed"
		}
		return value, "value"
	default:
		return 0, "pending"
	}
}
