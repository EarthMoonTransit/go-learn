# Test deposits through a shared pointer receiver

## Context

Write tests for the specified contract.

## API

```go
type Balance struct{Units int}
func (b *Balance) Deposit(amount int)
```

## Requirements

- Add amount to the receiver Units, visible through every alias of the same pointer.
- Repeated calls accumulate from the current balance; separate Balance values remain independent.
- Receiver is non-nil, initial Units and amounts are nonnegative and sums fit int.
- Zero deposits leave Units unchanged.

## Constraints

- Use only the Go standard library.
