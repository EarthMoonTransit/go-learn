# Calculate grouped and individual unit cost

## Context

Complete the specified operations.

## API

```go
func BatchCost(count int) uint
```

## Requirements

- Each complete group of ten units costs 95000.
- Each remaining individual unit costs 10000.
- Use as many complete groups as possible and return their combined cost.
- count is between 0 and 100000; zero units cost zero and the result fits uint.

## Constraints

- Use only the Go standard library.
