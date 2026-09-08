package subject

import (
	"sync"
)

func JoinAll(work []func()) <-chan struct{} {
	done := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(len(work))
	for _, run := range work {
		go func(f func()) { defer workers.Done(); f() }(run)
	}
	go func() { close(done); workers.Wait() }()
	return done
}
