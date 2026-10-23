# Run independent operation scopes

## Context
Give operations independent lifetimes.

## API
```go
type Result struct { Value int; Err error }
func RunScoped(parent context.Context, operations []func(context.Context) (int, error)) []Result
```

## Requirements
- Write a helper taking context first; use it to call all operations synchronously once, in order, even after errors or parent cancellation.
- Give each a fresh cancelable child of the original parent, preserving its values, deadline and cancellation.
- Release each child after its operation returns, before the next operation or return; preserve parent and sibling independence.
- Return ordered Results with exact values/errors, including values with errors. Empty input permits nil/empty results.
- Wait for every invoked operation's cleanup; no operation may remain running after return.

## Constraints
- Non-nil parent; 0–4 nonnil operations. Each returns normally after any supplied cleanup gate is released.
