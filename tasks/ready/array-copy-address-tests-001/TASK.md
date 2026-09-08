# Test array copies and element address identities

## Context

Write tests for the specified contract.

## API

```go
func CopyArray(source *[5]string) (*[5]string, [5]*string)
```

## Requirements

- Copy all five source strings into an independent array and return its pointer.
- Return the address of each copied array element in matching index order; addresses must refer to that returned array.
- Changing source or copied elements must not change the other array. Separate calls produce independent arrays.
- Source is non-nil; empty, repeated and Unicode strings are valid. Compare pointer identities, not numeric memory addresses.

## Constraints

- Use only the Go standard library.
