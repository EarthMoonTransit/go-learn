# Implement two independent classification rules

## Context

Complete the specified operations.

## API

```go
func Classify(identifier string, version int) (bool, bool)
```

## Requirements

- The first result is true exactly for identifier "alpha" or "beta".
- The second result is true exactly when version is at least 18, independently of the identifier.
- Identifiers are case-sensitive and are not trimmed; empty text is allowed.
- Version may be any int, including negative values. Do not combine the two decisions into one result.

## Constraints

- Use only the Go standard library.
