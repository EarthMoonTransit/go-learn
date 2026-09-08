# Distinguish an unknown category from an empty period

## Context

Complete the specified operations.

## API

```go
type Record struct{Day int;Amount float64;Category string}
type Period struct{From,To int}
var ErrCategory=errors.New("unknown category")
func CategoryTotal(in []Record, p Period, category string) (float64, error)
```

## Requirements

- Sum Amount for exact matching Category and Day within inclusive From..To.
- If the category occurs anywhere in input, return the sum and nil even when no matching day lies in the period or From exceeds To.
- If no input record has that category, return zero and the original non-nil ErrCategory sentinel; preserve the sentinel and every input record.
- Empty categories, negative and zero amounts are valid; finite amounts have magnitude <=1000 and at most 100 records. Allow absolute sum error 0.000000001.

## Constraints

- Use only the Go standard library.
