package subject

import "context"

type Result struct {
	Value int
	Err   error
}

// Write a private context-first helper and implement the operation scopes.
func RunScoped(parent context.Context, operations []func(context.Context) (int, error)) []Result {
	return nil
}
