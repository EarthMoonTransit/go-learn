# Finish an idle stream

## Context
Collect values until closure or inactivity.

## API
```go
func UntilIdle(input <-chan int, quiet time.Duration) (values []int, exhausted bool)
```

## Requirements
- Own one time.Timer whose first quiet interval begins when UntilIdle starts.
- Append every received value in order, preserving zero and duplicates; reset the timer to quiet after each accepted value.
- Return collected values and false when a full quiet interval elapses.
- Return collected values and true promptly when input closes, after its queued values.
- Stop the owned timer on every exit. Never close input or launch a background producer.
- When multiple events are simultaneously ready, either select outcome is permitted.

## Constraints
- Input is non-nil, quiet is positive, and input contains at most 16 values. Empty output may be nil or empty.
- The caller stops its own producer. Timer lifecycle requires method review; ticker behavior is outside this task.
