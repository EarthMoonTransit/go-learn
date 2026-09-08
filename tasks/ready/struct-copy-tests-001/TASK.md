# Test value fields and shared slice storage in a struct copy

## Context

Write tests for the specified contract.

## API

```go
type Snapshot struct{Label string;Count int;Items []int}
func CopySnapshot(input Snapshot) Snapshot
```

## Requirements

- Return an ordinary value copy of input: Label and Count are copied and Items retains the same slice view.
- Changing copied scalar fields does not affect the caller. Changing an existing Items element is visible through both structs.
- Preserve nil versus allocated-empty slices and the original slice length and capacity.
- Do not deep-copy, normalize or alter the input. Do not compare a struct containing a slice with ==.

## Constraints

- Use only the Go standard library.
