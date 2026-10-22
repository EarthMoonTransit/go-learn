package subject

func Transform(input <-chan int, delta int) <-chan int { out := make(chan int); close(out); return out }
