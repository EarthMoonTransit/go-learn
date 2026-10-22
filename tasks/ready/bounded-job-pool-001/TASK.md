# Run a bounded job pool

## Context
Execute independent callbacks with a worker limit.

## API
```go
type Result struct { Index int; Value int; Err error }
func RunPool(jobs []func() (int, error), workers int) []Result
```

## Requirements
- Invoke each callback exactly once.
- Exactly min(workers, len(jobs)) worker goroutines consume a shared job channel. With enough jobs, all workers can start while earlier callbacks remain blocked; never exceed that active limit.
- Return one result per job with its index, exact value and exact error, including a value accompanying an error.
- Errors do not stop other jobs. Return only after all callbacks and result collection complete.
- Result order is unspecified. Empty input permits nil or empty output.
- Merge results through a channel; use sync.WaitGroup to close it after every worker finishes. Method review inspects your implementation.

## Constraints
- 1–4 workers, 0–12 nonnil callbacks. Callbacks eventually return normally. No cancellation or panic handling.
