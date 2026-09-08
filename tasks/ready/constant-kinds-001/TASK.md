# Repair constant kinds and fractional constant arithmetic

## Context

Repair the supplied operation.

## API

```go
// constants.go: server and port constants
func Values() (string, int, float64)
```

## Requirements

- Declare server as an untyped string constant "localhost" and port as a typed int constant 8080 in constants.go.
- Values returns server, port and the result of the constant division 3 / 2.0 as float64.
- Preserve the fractional result: integer-only division has a different value.
- Public checks inspect constant kinds with go/types and verify all three returned values; named aliases of int are allowed.

## Constraints

- Use only the Go standard library.
