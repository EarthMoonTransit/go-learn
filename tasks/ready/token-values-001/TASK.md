# Map tokens with grouped switch cases

## Context

Complete the specified operations.

## API

```go
func TokenValue(token string) int
```

## Requirements

- Map two through nine to their corresponding integers, 2 through 9.
- Map ace to 11.
- Map ten, jack, queen and king to 10.
- All other strings return zero; matching is exact and case-sensitive.

## Constraints

- Use only the Go standard library.
