# Return a value address and the address of its pointer

## Context

Complete the specified operations.

## API

```go
func PointerChain() (*int, **int)
```

## Requirements

- Create an integer initialized to 20 and a pointer variable referring to it.
- Return the integer pointer and the address of that pointer variable; neither result is nil and the second dereferences to the first.
- Mutating the integer through either path is shared; replacing the pointer variable must not change the original integer.
- Calls allocate independent integer storage and pointer variables. Compare identities without numeric addresses.

## Constraints

- Use only the Go standard library.
