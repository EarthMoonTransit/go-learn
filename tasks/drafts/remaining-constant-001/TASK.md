# Declare a duration constant and compute the remainder

## Context

Complete the specified operations.

## API

```go
const targetMinutes = 40
func RemainingMinutes(elapsed int) int
```

## Requirements

- Declare targetMinutes as an integer constant with value 40.
- Return targetMinutes minus elapsed.
- elapsed is between 0 and 40 inclusive; zero elapsed returns the full target and 40 returns zero.
- The constant must remain usable in constant expressions.

## Constraints

- Edit subject/constants.go; subject/api.go provides the fixed Target accessor and constant-expression check.
