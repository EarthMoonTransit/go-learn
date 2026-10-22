package subject

import "fmt"

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
				_ = fmt.Sprint(value)
				close(output)
				return
			}
			output <- result
			close(output)
		}()
		result.Value = work()
	}()
	return output
}
