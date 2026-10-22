package subject

import (
	"fmt"
	"sync/atomic"
)

var priorPanic atomic.Bool

type Result struct {
	Value    int
	Err      error
	Panicked bool
}

func Launch(work func() int) <-chan Result {
	output := make(chan Result, 1)
	go func() {
		result := Result{}
		defer func() {
			if value := recover(); value != nil {
				priorPanic.Store(true)
				result = Result{Value: 0, Err: fmt.Errorf("worker panic: %v", value), Panicked: true}
			}
			if priorPanic.Load() && !result.Panicked {
				result.Panicked = true
				result.Err = fmt.Errorf("previous worker panic")
			}
			output <- result
			close(output)
		}()
		result.Value = work()
	}()
	return output
}
