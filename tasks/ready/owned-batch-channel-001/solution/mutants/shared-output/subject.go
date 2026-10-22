package subject

var shared = func() <-chan int { out := make(chan int); close(out); return out }()

func Publish(values []int, ready <-chan struct{}) <-chan int { return shared }
