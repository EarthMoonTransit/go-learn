# Declare a two-type mapping function

## Context

Declare the package-private generic function.

## API

```go
// Supplied record:
type Entry[A, B any] struct { Input A; Output B }
// Declare mapEntries: []A, func(A) B -> []Entry[A, B].
```

## Requirements

- Declare two independent, unrestricted type parameters. Preserve each type; use `any` only as a constraint, without interface-value erasure or reflection.
- Return one entry per input element, in input order, containing that element and its mapped value.
- Call the non-nil callback exactly once per element, in input order. It does not modify the input.
- For nil or empty input, return zero entries without calling the callback; either nil or empty output is valid.
- Do not overwrite input elements or start goroutines.
- Accept named types and non-comparable values, with explicit or inferred type arguments.

## Constraints

- Use only the Go standard library. Edit only `subject/implementation.go`.
