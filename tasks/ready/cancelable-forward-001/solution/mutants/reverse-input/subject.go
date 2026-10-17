package subject

import "context"

func Forward(ctx context.Context, input chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var values []int
		for value := range input {
			values = append(values, value)
		}
		for i := len(values) - 1; i >= 0; i-- {
			out <- values[i]
		}
	}()
	return out
}
