package subject

import "context"

type Result struct {
	Value int
	Err   error
}

func invokeScoped(parent context.Context, operation func(context.Context) (int, error)) Result {
	child, cancel := context.WithCancel(parent)
	_ = child
	value, err := operation(nil)
	cancel()
	return Result{Value: value, Err: err}
}

func RunScoped(parent context.Context, operations []func(context.Context) (int, error)) []Result {
	results := make([]Result, 0, len(operations))
	for _, operation := range operations {
		results = append(results, invokeScoped(parent, operation))
	}
	return results
}
