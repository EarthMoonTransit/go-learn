# Write table-driven tests for integer clamping

## Context

Replace the weak test submission with tests that distinguish the contract from all required defective implementations.

## API

```go
func Clamp(value, min, max int) int
```

## Requirements

- Keep values inside the inclusive range unchanged.
- Return `min` for values below the range.
- Return `max` for values above the range.
- Cover both exact boundaries and values beyond them.
- Use a named table of cases with actionable failure output.
- Detect every required mutant.

## Constraints

- Inputs always satisfy `min <= max`.
- Submit tests in package `subject_test`.
- Do not use timing, randomness, environment, network, or filesystem state.
