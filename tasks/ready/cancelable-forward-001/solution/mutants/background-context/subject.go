package subject

import "context"

func Forward(ctx context.Context, input chan int) <-chan int {
	ctx = context.Background()
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			var value int
			var open bool
			select {
			case <-ctx.Done():
				return
			case value, open = <-input:
				if !open {
					return
				}
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
