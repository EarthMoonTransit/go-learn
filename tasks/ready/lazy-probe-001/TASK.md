# Gate a supplied probe without eager calls

## Context

The supplied Probe records each call and returns a configured boolean. Implement a gate for that operation.

## API

```go
func CheckEnabled(enabled bool) bool
func Probe() bool
func ResetProbe(value bool)
func ProbeCalls() int
```

## Requirements

- When enabled is false, return false without calling Probe.
- When enabled is true, call Probe exactly once.
- Return that call's boolean result unchanged.
- Work on repeated sequential calls with either configured probe result.

## Constraints

- Edit only subject/implementation.go. Probe, ResetProbe and ProbeCalls are supplied; do not modify them or call the observation helpers from CheckEnabled. Calls are sequential. After ResetProbe(true), CheckEnabled(false) returns false with zero calls. After another reset to true, CheckEnabled(true) returns true with one call. Concurrency is outside this contract.
