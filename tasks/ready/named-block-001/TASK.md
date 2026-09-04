# Define a named array

## Context

Edit subject/bindings.go. Observe is supplied: it constructs block{2, 4, 6}, assigns it to a second variable, changes the second value at index zero to 99, and returns both.

## API

```go
// Edit subject/bindings.go; Observe is supplied.
func Observe() (any, any)
```

## Requirements

- Define a distinct Go type named block, not an alias.
- Its underlying type must be an array of exactly three int elements.
- The first observed value must remain [2, 4, 6].
- The second observed value must become [99, 4, 6].

## Constraints

- Keep the name block. A definition uses type Name Underlying; an alias adds an equals sign. Array assignment copies the elements. A slice would share element storage. The any result slots preserve types for the supplied checks.
