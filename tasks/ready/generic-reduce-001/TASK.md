# Repair a generic ordered reduction across two types

## Context

Repair the supplied operation.

## API

```go
func Reduce[A, B any](items []A, f func(B, A) B, initial B) B
```

## Requirements

- Starting with initial, call f once per input element in order and carry its returned accumulator to the next call.
- Return initial without calling f on nil or empty input. The callback is non-nil.
- Input and result types may differ. Preserve the input slice; no concurrent calls or environment access.
- Support named types and non-comparable accumulators; arithmetic inputs in tests do not overflow.

## Constraints

- Use only the Go standard library.
