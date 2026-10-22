# Find first positions of equal ordered pairs

## Context

Implement the supplied classification operation.

## API

```go
func FirstPositions(pairs [][2]string) []int
```

## Requirements

- Return one index for each input pair, in input order.
- Each index is the smallest input position containing exactly the same ordered pair of strings.
- Every byte of both components participates, including empty strings, delimiters and NUL bytes.
- Swapped components identify a different pair unless both components are equal. Equal values at different addresses identify the same pair.
- Preserve input values and order. Calls are independent.
- Nil/empty input returns a slice of length zero; nil identity is unspecified.

## Constraints

- All string values are valid inputs. Use only the standard library; equivalent implementations are accepted.
