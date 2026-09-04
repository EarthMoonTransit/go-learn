# Refactor a counter helper to mutate its receiver

## Context

Refactor the unexported helper so `Add` updates the original counter through its internal interface.

## API

```go
type Counter struct {
    Value int
}

func (c *Counter) Add(delta int)
```

## Requirements

- Preserve the exported API exactly.
- Make `Add` change the receiver's `Value`.
- Add positive, zero, and negative deltas.
- Accumulate changes across repeated calls.
- Keep the internal `counterAdder` interface assignment valid.
- Do not copy the receiver before applying the update.

## Constraints

- Modify implementation details, not the public contract.
- Use only the standard library.
