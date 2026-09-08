# Repair context without losing the original error

## Context

Repair the supplied operation.

## API

```go
func Wrap(operation string, err error) error
```

## Requirements

- Return nil when err is nil, regardless of operation text.
- Otherwise return an error with text operation + ": " + err.Error(). Preserve spaces and empty operation text.
- The returned error must unwrap directly to the original err, preserving its identity.
- Repeated wrapping adds one context layer each time; do not flatten earlier layers or modify the original error.

## Constraints

- Use only the Go standard library.
