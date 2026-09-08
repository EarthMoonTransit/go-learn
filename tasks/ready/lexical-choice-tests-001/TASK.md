# Test a symmetric lexical string choice

## Context

Write tests for the specified contract.

## API

```go
func ChooseLabel(a, b string) string
```

## Requirements

- Choose the string that comes first under Go string comparison.
- Return selected: followed by one space and the chosen string.
- Swapping the arguments does not change the result; equal strings return that same string after the prefix.
- Empty strings are valid; preserve exact case and whitespace and compare the complete strings.

## Constraints

- Use only the Go standard library.
