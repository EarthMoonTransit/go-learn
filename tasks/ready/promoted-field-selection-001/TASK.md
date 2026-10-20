# Repair selection of promoted and shadowed fields

## Context

Repair the function using the supplied value types. Return the five requested fields in the order listed below.

## API

```go
type Header struct { Code int; Name, Label string }
type Left struct { Value int }
type Right struct { Value int }
type Record struct { Header; Left; Right; Code int }
func SelectFields(r Record) (outer, embedded int, name string, left, right int)
```

## Requirements

- The Code declared directly on Record.
- The Code stored in its Header.
- Header.Name, also available as a promoted field.
- Left.Value.
- Right.Value.

## Constraints

- Return the actual values, including zeros, negatives and empty strings. Keep the supplied type declarations unchanged. Explicit qualification is accepted wherever promotion would obscure ownership; do not use reflection.
