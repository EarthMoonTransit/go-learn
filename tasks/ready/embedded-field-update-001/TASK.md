# Update promoted and shadowed fields independently

## Context

Implement ApplyCounts for the supplied types:

## API

```go
type Counters struct { Accepted, Rejected, Kept int }
type Layer struct { Counters; Rejected int; Tag string }
func ApplyCounts(in Layer, accepted, rejected, audit int) Layer
```

## Requirements

- Add accepted to the promoted Accepted field.
- Add rejected to Counters.Rejected.
- Add audit to the outer Layer.Rejected field. This field and Counters.Rejected are independent.
- Preserve Kept and Tag exactly.
- Return the updated value without changing the caller's input value.

## Constraints

- Increments can be negative or zero; all sums fit int. Preserve the supplied type definitions. Use only the standard library. Any equivalent implementation is valid.
