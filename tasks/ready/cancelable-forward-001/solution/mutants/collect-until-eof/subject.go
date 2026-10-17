package subject

import "context"

func Forward(ctx context.Context, input chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var values []int
		for {
			select {
			case <-ctx.Done():
				return
			case value, open := <-input:
				if !open {
					for _, v := range values {
						select {
						case out <- v:
						case <-ctx.Done():
							return
						}
					}
					return
				}
				values = append(values, value)
			}
		}
	}()
	return out
}
