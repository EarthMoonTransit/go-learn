# Capture each worker argument behind a start gate

## Context

Repair the batch launcher. The caller releases the start gate after the launcher returns.

## API

```go
func StartVisits(values []int, start <-chan struct{}, visit func(int)) <-chan struct{}
```

## Requirements

- Return a non-nil completion channel without waiting for start.
- Invoke visit exactly once per input element, passing that element's value after start closes.
- Let all callbacks start while others are blocked; their order is unspecified.
- Close completion after every callback returns; send no values.
- Nil/empty input invokes no callback and completes without releasing start.
- Preserve input elements and keep batches independent, including concurrent calls.

## Constraints

- start is initially open and receives no sends. The caller closes it after return and preserves input elements until completion. start and visit are non-nil; callbacks return normally. Use only the standard library.
