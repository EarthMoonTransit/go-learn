# Implement increments in one owned critical section

## Context

Complete the specified operations.

## API

```go
type Counter struct{mu sync.Mutex;value int}
func NewCounter()*Counter{return &Counter{}}
func(c *Counter)Inc()
func(c *Counter)Value()int
```

## Requirements

- Complete Inc. A new or zero-value Counter starts at zero; each call adds exactly one, including simultaneous calls on the same Counter.
- Use the Counter-owned mutex to protect its shared value. Do not copy the Counter after first use. Different counters retain independent state.
- The provided Value is called only after all concurrent increments have completed in this lesson; leave its behavior unchanged.
- Public checks include the source three-increment and 1000-worker scenarios. Check and Complete run the race detector; a numerically correct result with a data race is incorrect.

## Constraints

- Use only the Go standard library.
