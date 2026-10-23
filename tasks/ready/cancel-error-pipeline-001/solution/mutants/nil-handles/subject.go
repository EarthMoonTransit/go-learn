package subject

import "context"

func TransformUntil(parent context.Context, input <-chan int, step func(int) (int, error)) (output <-chan int, terminal <-chan error) {
	return nil, nil
}
