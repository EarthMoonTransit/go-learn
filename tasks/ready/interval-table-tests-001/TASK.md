# Choose a named test table

## Context

Write tests for the supplied contract.

## API

```go
func Contains(lo, hi, x int) bool
```

## Requirements

- Return true exactly when `lo <= x && x < hi`; empty or reversed intervals contain nothing.
- Build your own named table covering below, lower endpoint, interior, upper endpoint, above, empty, reversed and negative-coordinate cases.
- Share readable execution and assertions across the table; report the case name, input, actual result and expected result.
- Keep each case's setup and expected outcome explicit. Do not modify the supplied implementation.

## Constraints

- Use package subject_test and the standard testing package. All int inputs are admitted; do not rely on arithmetic that overflows.

Public checks require compatibility with the correct implementation. Full checks challenge your tests with defects. Method review separately inspects meaningful named partitions and shared readable execution; mutation detection alone does not prove a table was written.
