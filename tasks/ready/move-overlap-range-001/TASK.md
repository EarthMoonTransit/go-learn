# Move an overlapping range atomically

## Context

Implement the supplied operation.

## API

```go
func MoveRange(buf []int, dst, src, n int) bool
```

## Requirements

- A valid range has nonnegative dst, src and n, with both n-element regions entirely within len(buf).
- For valid ranges, write the original src-region values into the dst-region and return true.
- Source and destination may overlap in either direction or coincide; all reads use the pre-call values.
- Modify the supplied storage so aliases observe changes. Preserve all cells outside the destination region, including spare capacity.
- Invalid ranges return false without changing storage or panicking, including very large indexes or counts.
- Zero-length ranges succeed when both endpoints are within 0..len(buf), inclusive. Nil storage accepts only dst=src=n=0.

## Constraints

- Values may be negative, zero or repeated. Use only the standard library; equivalent implementations are accepted.
