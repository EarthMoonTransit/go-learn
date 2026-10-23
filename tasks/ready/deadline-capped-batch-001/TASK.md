# Share a bounded batch deadline

## Context
Use one end-to-end budget.

## API
```go
type Result struct { Value int; Err error }
func RunBudgeted(parent context.Context, budget time.Duration, ceiling time.Time, operations []func(context.Context) (int, error)) []Result
```

## Requirements
- Start budget once at entry; share one derived operation context across all callbacks.
- Its deadline is the earliest of entry plus budget, nonzero ceiling, and any parent deadline; zero ceiling adds no limit.
- Preserve parent values/cancellation. Nonpositive budget or an already-past effective deadline means expired before invocation.
- Invoke all callbacks synchronously once in order, including after errors or expiry; never restart the budget.
- Return ordered Results preserving exact values/errors, including values with errors; nil/empty output is allowed for empty input.
- Wait for callback cleanup; release every owned context on return without canceling parent.

## Constraints
- Non-nil parent; 0–3 nonnil callbacks return normally after supplied cleanup gates are released.
