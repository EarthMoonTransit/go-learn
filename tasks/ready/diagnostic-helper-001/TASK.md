# Implement a caller-aware assertion helper

## Context

Complete the specified operations.

## API

```go
type Reporter interface{Helper();Errorf(format string,args ...any)}
func AssertInt(t Reporter, got, want int)
```

## Requirements

- Call Helper exactly once before any assertion reporting; a *testing.T satisfies the supplied Reporter interface.
- When got equals want report nothing. Otherwise call Errorf once with format "got %d, want %d" and the two integer arguments in that order.
- Do not panic or terminate the caller. Zero and negative values are valid.
- The public recording reporter makes helper registration, ordering and diagnostics observable.

## Constraints

- Use only the Go standard library.
