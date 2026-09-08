# Synchronize reads with existing counter updates

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

- Complete Value so it reads the current counter safely while other goroutines call Inc. Inc is supplied and already locks the Counter mutex.
- Use the same synchronization for reads and writes. Value must not mutate the count; it returns a value observed during that call.
- A new/zero counter reads zero. Repeated reads after all increments return the exact count; separate counters remain independent.
- The original source reads only after Wait. This extension explicitly permits overlapping reads and writes; race-free execution is required in public and full checks.

## Constraints

- Use only the Go standard library.
