package subject

import "context"

type Result struct {
	Value int
	Err   error
}

func invoke(ctx context.Context, operation func(context.Context) (int, error)) Result {
	value, err := operation(ctx)
	return Result{value, err}
}
func RunScoped(parent context.Context, operations []func(context.Context) (int, error)) []Result {
	var results []Result
	previous := parent
	for _, operation := range operations {
		child, cancel := context.WithCancel(previous)
		results = append(results, invoke(child, operation))
		cancel()
		previous = child
	}
	return results
}
