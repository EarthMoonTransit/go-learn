# Reduce a stored quantity and delete exactly at zero

## Context

Complete the specified operations.

## API

```go
func RemoveQuantity(items, units map[string]int, item, unit string) bool
```

## Requirements

- Return false without changing either map if item or unit is missing or subtraction would be negative.
- Otherwise subtract the unit quantity, delete the item when its new quantity is zero, and return true.
- Item quantities are nonnegative; unit quantities are positive. Maps do not share storage; nil maps contain no entries.
- Preserve unrelated items and all unit entries. Keys are exact strings, including empty strings.

## Constraints

- Use only the Go standard library.
