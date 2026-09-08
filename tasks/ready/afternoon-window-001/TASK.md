# Classify a parsed time in a half-open hour window

## Context

Complete the specified operations.

## API

```go
var ErrDate error
func IsAfternoon(text string) (bool, error)
```

## Requirements

- Parse with layout "Monday, January 2, 2006 15:04:05" using time.Parse rules, including its weekday handling and optional fractional seconds.
- Return true only for parsed times at or after 12:00 and before 18:00; return false for other valid times.
- On invalid input return false and the supplied non-nil ErrDate. Preserve its identity.
- The result depends only on text; do not use the current date or machine timezone.

## Constraints

- Use only the Go standard library.
