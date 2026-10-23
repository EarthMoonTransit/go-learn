# Cancel pipeline stages
## Context
Coordinate finite stream termination.
## API
```go
func TransformUntil(parent context.Context, input <-chan int,
    step func(int)(int,error)) (output <-chan int, terminal <-chan error)
```
## Requirements
- Return non-nil channels promptly: unbuffered output, capacity-one terminal.
- Author feeder and mapper workers connected by an unbuffered channel. Every blocking pipeline send/receive observes an owned child context.
- Map sequentially once per processed input; stream successful results in order.
- First step error omits its value, preserves earlier results, and cancels feeder without EOF.
- Parent cancellation stops both workers; pre-canceled entry performs no reads/calls.
- Close output; report one result after both workers exit, then close terminal: nil on EOF, exact step error, or parent.Err.
## Constraints
- Non-nil inputs; at most 12 integers; step eventually returns. Caller owns input/parent.
- Normal completion requires EOF/drain; early completion requires error/cancellation. Already-winning sends and either co-ready termination cause are allowed. Explain worker termination. Standard library only.
