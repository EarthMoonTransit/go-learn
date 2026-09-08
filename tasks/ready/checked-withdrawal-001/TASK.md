# Implement withdrawal that preserves state on failure

## Context

Complete the specified operations.

## API

```go
type Balance struct{Units int}
var ErrInsufficient=errors.New("insufficient units")
func (b *Balance) Withdraw(amount int) error
```

## Requirements

- For amount <= Units subtract amount and return nil. Withdrawing exactly the balance is allowed.
- If amount exceeds Units return the original non-nil ErrInsufficient sentinel and leave Units unchanged.
- The receiver is non-nil; initial Units and amount are nonnegative integers <=1000000. Zero withdrawals are valid.
- Update the shared receiver visible through aliases, preserve the sentinel, and use the current balance on repeated calls.

## Constraints

- Use only the Go standard library.
