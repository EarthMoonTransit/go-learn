# Repair a fresh pointer and nil error result

## Context

Repair the supplied operation.

## API

```go
type Record struct{Label string;Count int}
func NewRecord(label string, count int) (*Record, error)
```

## Requirements

- Return a non-nil pointer to a Record containing Label and Count exactly.
- Return a nil error for every valid input; negative counts and empty labels are valid.
- Each call returns independent storage, even for identical inputs.
- Changing a returned record must not affect earlier or later results.

## Constraints

- Use only the Go standard library.
