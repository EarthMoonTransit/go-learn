# Transform a finite stream

## Context
Compose two communicating stages.

## API
```go
func Transform(input <-chan int, delta int) <-chan int
```

## Requirements
- Return the output channel promptly, before input closes.
- Stage one adds delta to every input; stage two forwards only nonnegative transformed values.
- Preserve order and duplicates; eligible values must stream before input exhaustion.
- Close final output after all input is processed. Never close the caller's input.
- Author two communicating stages with channel range loops and separate intermediate/final send-close ownership. Explain each termination path.

## Constraints
- Input is non-nil, carries at most16 integers in[-20,20], then its sole producer closes it; delta is in[-5,5].
- The caller drains output; there is no early abandonment or cancellation.
- Actual stage composition, range loops and ownership require method review. A behavior-correct fused stage may pass tests but fails this method requirement.
