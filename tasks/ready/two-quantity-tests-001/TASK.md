# Test two independently accumulated results

## Context

Write tests for the specified contract.

## API

```go
func Quantities(kinds []string) (int, float64)
```

## Requirements

- Each exact solid entry contributes 50 to the integer result; each exact liquid entry contributes 0.2 to the floating result.
- Other strings contribute nothing; matching is case-sensitive.
- Return integer quantity first and floating quantity second; nil or empty input returns two zeros.
- Preserve the input. At most 10000 entries are supplied; allow absolute floating error 0.000000001.

## Constraints

- Use only the Go standard library.
