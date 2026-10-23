# Stream job outcomes

## Context
Collect bounded concurrent jobs.

## API
```go
type Result struct { Index, Value int; Err error; Panicked bool; Panic any }
func StreamJobs(jobs []func() (int, error), workers int) <-chan Result
```

## Requirements
- Return nonnil output promptly, even while callbacks are blocked.
- Exactly min(workers,len(jobs)) worker goroutines consume one shared index queue; call each job once without callback goroutines.
- Recover each job inside its worker and continue subsequent jobs.
- Normal results preserve Value/Err with Panicked=false,Panic=nil. Panics produce Value=0,Err=nil,Panicked=true and the original Panic value.
- Merge one indexed result per job; publish completed results while other callbacks remain blocked. Ordering is unspecified.
- Close output after workers and results finish; empty input closes immediately. Method review checks recovery, queue and merge.

## Constraints
- 1–3 workers; 0–8 nonnil, eventually terminating callbacks. Panics are strings, ints or nonnil errors; exclude panic(nil)/Goexit. Caller drains output.
