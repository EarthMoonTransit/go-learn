package subject

import (
	"sync"
)

func JoinAll(work []func()) <-chan struct{} {
	done := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(len(work))
	for _, run := range work {
		func(f func()) { defer workers.Done(); f() }(run)
	}
	go func() { workers.Wait(); close(done) }()
	return done
}
