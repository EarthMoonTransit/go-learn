# Apply stored functions until a result is rejected

## Context

Implement the supplied operation.

## API

```go
func RunStages(initial int, steps []func(int) (int, bool)) (int, int)
```

## Requirements

- Start with initial as the committed value and zero accepted stages.
- Visit entries in slice order. Skip nil functions without increasing the count.
- Call each encountered non-nil entry exactly once with the current committed value.
- A true result replaces that value and increases the accepted count.
- A false result stops immediately: discard its proposed value, exclude it from the count, and call no later entries.
- Return the final committed value and accepted count. Preserve every entry in steps; repeated entries are separate stages.

## Constraints

- Functions return normally and do not modify steps. Nil and empty slices are valid. Use only the standard library; equivalent implementations are accepted.
