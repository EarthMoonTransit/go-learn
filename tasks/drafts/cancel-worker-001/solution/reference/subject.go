package subject

import (
	"context"
)

func StartCancelable(parent context.Context) (context.Context, context.CancelFunc, <-chan struct{}) {
	child, cancel := context.WithCancel(parent)
	done := make(chan struct{})
	go func() { defer close(done); <-child.Done() }()
	return child, cancel, done
}
