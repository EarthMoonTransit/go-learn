package subject

import (
	"context"
	"sync"
)

func TransformUntil(parent context.Context, input <-chan int, step func(int) (int, error)) (output <-chan int, terminal <-chan error) {
	values := make(chan int)
	result := make(chan error, 1)
	child, cancel := context.WithCancel(parent)
	if parent.Err() != nil {
		cancel()
		close(values)
		result <- parent.Err()
		close(result)
		return values, result
	}
	intermediate := make(chan int)
	var joined sync.WaitGroup
	joined.Add(2)
	var reason error
	go func() {
		defer joined.Done()
		defer close(intermediate)
		for {
			select {
			case <-child.Done():
				return
			case value, ok := <-input:
				if !ok {
					return
				}
				select {
				case <-child.Done():
					return
				case intermediate <- value:
				}
			}
		}
	}()
	go func() {
		defer joined.Done()
		defer close(values)
		for {
			select {
			case <-child.Done():
				reason = parent.Err()
				return
			case value, ok := <-intermediate:
				if !ok {
					if parent.Err() != nil {
						reason = parent.Err()
					}
					return
				}
				mapped, err := step(value)
				if err != nil {
					reason = err
					continue
				}
				select {
				case <-child.Done():
					reason = parent.Err()
					return
				case values <- mapped:
				}
			}
		}
	}()
	go func() {
		joined.Wait()
		cancel()
		result <- reason
		close(result)
	}()
	return values, result
}
