# Repair workspace behavior

## Context

Implement the editable package file so the adapter preserves its integer input.

## API

```go
func Apply(value int) int
```

## Requirements

- Preserve positive inputs.
- Preserve zero.
- Preserve negative inputs.
- Keep package boundaries unchanged.

## Constraints

- Change only the declared editable path.
