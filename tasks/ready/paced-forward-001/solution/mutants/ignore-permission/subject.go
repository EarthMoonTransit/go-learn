package subject

func Pace(input <-chan int, release <-chan struct{}) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		incoming, tokens := input, release
		pending, permitted := false, false
		value := 0
		for incoming != nil || pending {
			receiving := incoming
			if pending {
				receiving = nil
			}
			granting := tokens
			if !pending || permitted {
				granting = nil
			}
			var sending chan<- int
			if pending && permitted {
				sending = output
			}
			select {
			case next, ok := <-receiving:
				if !ok {
					incoming = nil
					continue
				}
				value = next
				pending = true
				permitted = true
			case _, ok := <-granting:
				permitted = true
				if !ok {
					tokens = nil
				}
			case sending <- value:
				pending = false
				permitted = false
			}
		}
	}()
	return output
}
