# Advance one slice header while preserving shared views

## Context

Repair the supplied operation.

## API

```go
func TakePrefix(input *[]int, n int) ([]int, bool)
```

## Requirements

- Let old be the slice stored at input before the call. Valid n means 0 <= n <= len(old).
- On success, return old's first n elements and true, with length n and capacity cap(old).
- Advance the caller's slice to the remaining elements: length len(old)-n, capacity cap(old)-n.
- Both results share original storage. Preserve every backing value during the call and every independently copied header.
- Invalid n returns nil,false without changing the caller's header or storage.
- A nil slice accepts n=0 and leaves both slices nil. Successive calls use the current header.

## Constraints

- input is non-nil. Use only the standard library; equivalent implementations are accepted.
