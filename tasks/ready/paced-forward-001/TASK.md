# Pace a finite stream

## Context
Release tokens permit pending values.

## API
```go
func Pace(input <-chan int, release <-chan struct{}) <-chan int
```

## Requirements
- Return a private nonnil unbuffered output promptly.
- Use one select-loop worker. Set local channels nil to disable input while pending, release when no unpermitted value is pending, and output until permitted.
- Forward each value once in order, including zero/duplicates; hold at most one unsent value.
- Consume exactly one token per pending value needing permission. After permission, preserve further tokens until that value is sent.
- Closing release permits every remaining value; preserve pending data and caller channels.
- Close output only after input closes and pending data is sent.

## Constraints
- Inputs are nonnil; at most 8 values and 12 tokens. Caller eventually closes both and drains output. Review actual select/nil method.
