package subject

import "context"

func Forward(ctx context.Context, input chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			var value int
			var open bool
			value, open = <-input
			if !open {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- value:
			}
		}
	}()
	return out
}
