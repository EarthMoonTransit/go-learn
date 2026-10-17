# Repair deferred registration and argument capture

## Context

Repair the supplied runTrace body so recorded values show registration-time arguments and deferred execution order.

## API

```go
const Checkpoint = 1000001
func Trace(seed int) []int // supplied wrapper
func runTrace(seed int, record func(int), checkpoint func())
```

## Requirements

- Defer the first two record calls within runTrace's existing function boundary.
- Save each call's argument at registration, before the later seed changes.
- Execute the third record call and checkpoint immediately.
- Produce exactly [seed+9, Checkpoint, seed+4, seed], with deferred calls in LIFO order.
- Keep calls independent and preserve the supplied wrapper and callbacks.

## Constraints

- Seeds range from -1000 through 1000. Edit only subject/implementation.go. Two deferred calls are required; public checks name a missing defer. For seed 2 the trace is [11, 1000001, 6, 2]. Do not add I/O or concurrency.
