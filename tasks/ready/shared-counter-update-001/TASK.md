# Mutate an integer through a shared pointer

## Context

Complete the specified operations.

## API

```go
func AddCounter(counter *int, delta int)
```

## Requirements

- Add delta to the integer referenced by counter.
- The caller and every alias to that storage observe the updated value.
- A zero delta leaves the integer unchanged; negative deltas subtract.
- counter is non-nil and the sum fits int; repeated calls accumulate.

## Constraints

- Use only the Go standard library.
