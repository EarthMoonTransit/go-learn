# Collect every indexed job result and error

## Context

Collect a finite batch concurrently. Jobs are non-nil and return normally.

## API

```go
type Result struct { Index int; Value int; Err error }
func CollectResults(jobs []func() (int, error)) []Result
```

## Requirements

- Invoke every job exactly once. All jobs must be able to start while the others remain blocked.
- Return one record per job, in input order. Record i contains Index i and that job's exact value and error; preserve a nonzero value returned with an error.
- Wait for every job to return, including after another job fails. Do not use a fail-fast policy.
- Nil and empty input return zero records; either nil or empty output is valid.
- Keep returned result records from separate calls independent, including concurrent calls.

## Constraints

- Use only the Go standard library. No cancellation or panic-recovery policy is required.
