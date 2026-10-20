# Repair indexed array window updates

## Context

Repair the supplied function; changing a range value alone does not update an array element.

## API

```go
func PatchWindow(a [8]int, lo, hi, value int) ([8]int, bool)
```

## Requirements

- Bounds are valid exactly when `0 <= lo <= hi <= 8`.
- For valid bounds, replace each element at an index in `[lo, hi)` with `value` and return `true`.
- Preserve every element outside that interval.
- An empty interval, including `[8, 8)`, succeeds without changing any element.
- For invalid bounds, return the unchanged array and `false`; do not panic or normalize the bounds.

## Constraints

- Use only the standard library. The supplied array is a value; return the updated value. Any implementation satisfying the contract is accepted.
