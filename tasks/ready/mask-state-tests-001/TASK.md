# Test selected-bit operations

## Context

Write tests for the supplied API to expose incorrect selected-bit operations.

## API

```go
type State struct { Any, All bool; Set, Clear, Toggle uint8 }
func Observe(bits, mask uint8) State
```

## Requirements

- Any reports whether at least one selected bit is present.
- All reports whether every selected bit is present. With mask zero, Any is false and All is true.
- Set turns selected bits on.
- Clear turns selected bits off.
- Toggle inverts selected bits.
- All three updates preserve unselected bits.

## Constraints

- Test zero, full, overlapping, already-set and disjoint masks. Values are eight-bit unsigned integers. Write tests in package subject_test using the standard testing package. Correct implementations must pass; each declared kind of error must be exposed by an assertion. Do not modify the implementation or assume a particular evaluation strategy.
