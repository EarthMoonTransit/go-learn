# Update elements through range

## Context

Repair a function whose loop changes only iteration copies.

## API

```go
func OffsetByIndex(values []int)
```

## Requirements

- Increase each element at index i by i + 1, using its original value.
- Apply updates to the supplied slice so callers sharing those elements observe them.
- Nil and empty slices require no changes and must not panic.
- Change only elements inside the supplied slice window.

## Constraints

- Edit subject/subject.go. Length is at most 64; original elements are between -1000 and 1000. Keep the signature. A range value is a copy; the range index can address the original element.
