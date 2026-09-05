# Test inclusive interval comparisons

## Context

Extend the supplied test scaffold to check interval boundaries and value preservation. The function implementation is supplied by the checker.

## API

```go
func Clamp(value, min, max int) int
```

## Requirements

- Values below min return min; values above max return max.
- Values inside the inclusive interval remain unchanged.
- Equal endpoints form a valid interval and always return that endpoint.
- Check below, lower equality, interior, upper equality, above and equal-endpoint behavior.
- Your tests must pass a correct implementation and reject each required defective behavior.

## Constraints

- Inputs always satisfy min <= max. Edit subject/subject_test.go in package subject_test; keep tests deterministic. The starter supplies the import and test/assertion pattern; use it for additional cases. Any clear test organization is accepted. No table structure or diagnostic style is graded.
