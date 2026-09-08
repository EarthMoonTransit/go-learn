# Test complete slice aggregation

## Context

Write tests for the specified contract.

## API

```go
func TotalCounts(counts []int) int
```

## Requirements

- Write tests for the sum of all values, including the first and final elements.
- Nil and empty inputs return zero; zero-valued elements do not stop processing.
- A single element returns its value.
- Do not change input elements; values are nonnegative and the total fits int.

## Constraints

- Use only the Go standard library.
