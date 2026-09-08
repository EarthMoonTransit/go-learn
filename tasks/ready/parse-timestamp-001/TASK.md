# Parse a fixed timestamp layout in UTC

## Context

Complete the specified operations.

## API

```go
var ErrDate error
func ParseTimestamp(text string) (time.Time, error)
```

## Requirements

- Parse text with Go layout "1/2/2006 15:04:05", retaining time.Parse acceptance rules, including optional fractional seconds.
- Successful values use UTC and preserve date, hour, minute, second and fractional second. Do not read the current time or local timezone.
- On any parse error return the zero time and supplied non-nil ErrDate. Preserve that sentinel.
- Month/day order follows the layout; invalid calendar dates and trailing text are errors.

## Constraints

- Use only the Go standard library.
