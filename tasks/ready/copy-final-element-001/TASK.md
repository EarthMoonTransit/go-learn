# Repair an update through a slice parameter

## Context

Repair the supplied operation.

## API

```go
func CopyLast(source, destination []string)
```

## Requirements

- Replace the final destination element with the final source element.
- Preserve every earlier destination element and the destination length.
- Changes are visible through existing destination views; preserve every source element.
- Both slices are nonempty and their backing storage does not overlap; lengths may differ.

## Constraints

- Use only the Go standard library.
