# Repair mutation through an expanded variadic slice

## Context

Repair the supplied operation.

## API

```go
func ReplaceFirst(replacement int, values ...int) bool
```

## Requirements

- If values is empty return false. Otherwise replace only its first element and return true.
- When a caller passes a slice with values..., the caller must observe the change through shared backing storage.
- Preserve length, capacity and all other elements, including storage outside a passed subslice. Zero and negative integers are valid.
- Direct variadic arguments also work. Do not append, clone the slice or mutate a different position.

## Constraints

- Use only the Go standard library.
