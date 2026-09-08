package subject

import (
	"sync"
)

func JoinAll(work []func()) <-chan struct{} {
	done := make(chan struct{}, 1)
	var workers sync.WaitGroup
	workers.Add(len(work))
	for _, run := range work {
		go func(f func()) { defer workers.Done(); f() }(run)
	}
	go func() { workers.Wait(); done <- struct{}{}; close(done) }()
	return done
}
