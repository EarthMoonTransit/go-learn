# Strengthen identity tests

## Context

Write tests that distinguish a correct identity function from required incorrect implementations.

## API

```go
func Identity(value int) int
```

## Requirements

- Keep the test suite deterministic.
- Exercise a nonzero input.
- Call the exported API.
- Detect every required incorrect implementation.

## Constraints

- Submit only Go test files.
