# Initialize dependent package values before init

## Context

Complete the specified operations.

## API

```go
// values.go declares base and doubled, and registers one init function.
func Values() (int, int)
// Supplied Trace() []string exposes initialization events.
```

## Requirements

- In values.go declare doubled before base. Initialize base with record("base",3) and doubled with record("double",base*2).
- Add an init function calling record("ready",0) after variable initialization. Do not call record from Values.
- Values returns base and doubled: 3 and 6. Supplied Trace must observe base, double, ready exactly once.
- The supplied record helper returns its value unchanged. Calls to Values and Trace do not change initialization state.

## Constraints

- Use only the Go standard library.
