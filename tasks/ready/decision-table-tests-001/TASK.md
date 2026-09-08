# Test ordered decisions over weighted tokens

## Context

Write tests for the specified contract.

## API

```go
func SelectAction(first, second, visible string) string
```

## Requirements

- Weights: ace=11; two through nine=2 through 9; ten, jack, queen, king=10; every other exact token=0.
- Sum the first two weights. At 22 return P. At 21 return S if the visible weight is at least 10, otherwise W.
- For sums 17..20 return S. For sums 12..16 return S when the visible weight is below 7, otherwise H. For sums below 12 return H.
- Codes are literal uppercase strings. Token matching is case-sensitive and preserves whitespace; unknown and empty tokens are allowed.

## Constraints

- Use only the Go standard library.
