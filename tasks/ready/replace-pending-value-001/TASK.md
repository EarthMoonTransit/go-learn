# Receive a pending value before sending its replacement

## Context

Replace one pending value in an open full channel without blocking on its existing value.

## API

```go
func ReplacePending(ch chan int, next int) int
```

## Requirements

- Receive the pending old value before sending next.
- Return exactly the received old value, including zero and negative values.
- Leave exactly next pending in the same channel.
- Keep the caller's channel open.

## Constraints

- ch is open with capacity one and exactly one pending integer. No other caller sends or receives during the operation. Edit subject/subject.go; preserve the API. Use only the standard library. The supplied harness handles goroutine scheduling and cleanup.
