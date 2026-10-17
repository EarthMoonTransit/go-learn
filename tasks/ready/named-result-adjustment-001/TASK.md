# Repair a deferred named-result adjustment

## Context

The supplied deferred adjustment changes caller-owned data instead of the named result. Repair it while preserving callback behavior.

## API

```go
func Adjust(base, delta int, caller *int, work func()) (result int)
```

## Requirements

- Call work exactly once before returning.
- Evaluate the return expression base+3, then adjust the named result by delta in a deferred call.
- On normal return, produce base+3+delta.
- Never change the value pointed to by caller.
- Propagate a panic from work with its original value; do not recover or replace it.

## Constraints

- base and delta range from -1000 through 1000; caller and work are non-nil. Edit subject/subject.go. Keep the named result and deferred mechanism; public form checks report their absence. Adjust(4, 7, caller, work) returns 14 when work returns normally. No I/O or concurrency.
