# Build an integer list

## Context
Implement `FormatValues` in `subject/subject.go`. Construct one string from a prefix and the supplied integer slice.

## API
```go
func FormatValues(prefix string, values []int) string
```

## Requirements
- Begin with the exact, unchanged `prefix`, followed by `[`.
- Include every value once, in its original order, using decimal integers with their signs.
- Put one comma between adjacent values, with no inserted spaces or trailing comma.
- End with `]`; nil and empty slices both produce `prefix + "[]"`.
- Leave every element of the input slice unchanged.
- Add no newline or other output.

## Constraints
- Keep the signature. For example, prefix `v=` and values `3, -2, 0` produce `v=[3,-2,0]`. Concatenation or a string builder is acceptable; this task assesses output and input preservation, without an allocation or timing threshold.
