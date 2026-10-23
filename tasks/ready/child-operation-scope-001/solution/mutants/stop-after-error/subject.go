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
	results := make([]Result, 0, len(operations))
	for _, operation := range operations {
		result := invokeScoped(parent, operation)
		results = append(results, result)
		if result.Err != nil {
			break
		}
	}
	return results
}
