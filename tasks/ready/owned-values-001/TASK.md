# Return a finite receive-only stream without a blocked producer

## Context

Complete the specified operations.

## API

```go
func OwnedValues(values []int) <-chan int
```

## Requirements

- Return a non-nil receive-only channel containing a snapshot of values in order, followed by closure. All values and closure are available when this function returns.
- The function owns the sending side and closes it. Use sufficient buffering so no receiver or background producer is needed to complete creation.
- Preserve input storage and keep separate calls independent. Changes to the input after return must not affect queued values.
- Empty input returns an already closed channel. Do not leave a goroutine waiting for a receiver; an unused result must not retain a worker.

## Constraints

- Use only the Go standard library.
