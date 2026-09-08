# Declare a consecutive month enumeration

## Context

Complete the specified operations.

## API

```go
// months.go: declare jan through dec as constants
func GetMonth(month int) int
```

## Requirements

- In months.go declare jan through dec in one const block using iota, with consecutive values 1 through 12.
- The supplied API exports these as Jan through Dec and requires actual constant expressions.
- GetMonth returns its argument unchanged; callers pass one of the declared month constants.
- Public checks verify all twelve values and the use of iota in the month constant block.

## Constraints

- Use only the Go standard library.
