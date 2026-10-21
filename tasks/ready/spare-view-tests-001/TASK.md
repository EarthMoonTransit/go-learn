# Test a shared view with a limited append reserve

## Context

Write tests for the supplied API.

## API

```go
func SpareView(in []int, low, high, spare int) []int
```

## Requirements

- Return elements in[low:high], with length high-low and capacity high-low+spare.
- Selected elements share input storage: writes through either view affect the other.
- Appending up to spare elements reuses the following backing cells, including cells beyond len(in).
- Appending beyond that capacity produces independent storage; that append must not change input backing values.
- Preserve all input storage during the call. Empty views and nil input are valid; nil identity of empty results is unspecified.

## Constraints

- Inputs satisfy 0 <= low <= high <= len(in) and 0 <= spare <= cap(in)-high. Write subject_test tests using testing. Correct implementations must pass; defects must fail assertions. Do not require a particular capacity after growth.
