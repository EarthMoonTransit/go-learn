package subject

import "context"

type Result struct {
	Value int
	Err   error
}

func invokeScoped(parent context.Context, operation func(context.Context) (int, error)) Result {
	child, cancel := context.WithCancel(parent)
	value, err := operation(child)
	cancel()
	return Result{Value: value, Err: err}
}

func RunScoped(parent context.Context, operations []func(context.Context) (int, error)) []Result {
	go func() {
		for _, operation := range operations {
			invokeScoped(parent, operation)
		}
	}()
	return nil
}
