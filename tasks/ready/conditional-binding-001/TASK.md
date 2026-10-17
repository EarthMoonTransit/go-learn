# Repair assignment across a conditional block

## Context

The supplied conditional update is lost before the function returns. Repair the local-name collision.

## API

```go
func SelectValue(initial, next int, useNext bool) int
```

## Requirements

- Return next when useNext is true.
- Return initial when useNext is false.
- Keep the supplied if useNext { ... } block and repair the collision within it.
- Handle zero, negative and other int values independently on each call.

## Constraints

- Edit only subject/implementation.go. The supplied API file and tests are fixed. All int values are valid. SelectValue(7, -4, true) returns -4; SelectValue(7, -4, false) returns 7. The public form check reports removal of the required conditional block.
