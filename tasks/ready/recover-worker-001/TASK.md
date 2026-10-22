# Recover at a worker boundary

## Context
Run one callback asynchronously and publish its outcome.

## API
```go
type Result struct { Value int; Err error; Panicked bool }
func Launch(work func() int) <-chan Result
```

## Requirements
- Return a nonnil result channel promptly, even while work is blocked.
- Invoke work exactly once inside the launched goroutine.
- Publish exactly one Result after work finishes or panics, then close the channel.
- Normal completion preserves the returned value and sets Err to nil and Panicked to false.
- Recover a string or int panic inside that worker. Publish Value 0, Panicked true and a nonnil error with exact text "worker panic: <value>". Empty strings and integer zero are panics too.
- Calls have independent state. Method review inspects the worker's deferred recovery and publication boundary.

## Constraints
- Work is nonnil and eventually returns or panics. Exclude panic(nil), runtime.Goexit and cancellation. The caller drains the output.
