package subject

import (
	"context"
)

func StartCancelable(parent context.Context) (context.Context, context.CancelFunc, <-chan struct{}) {
	child, cancel := context.WithCancel(parent)
	done := make(chan struct{}, 1)
	go func() { defer close(done); <-child.Done(); done <- struct{}{} }()
	return child, cancel, done
}
