# Bound in-flight stream values

## Context
Bound a live stream's queued and held values.

## API
```go
func Window(input <-chan int, limit int) <-chan int
```

## Requirements
- Return a private non-nil output promptly, before EOF.
- Use one worker: receive a value, finish its send, then receive again. Never reserve output space before receiving.
- Forward every value exactly once in order, including zero, negatives and duplicates. Stream before EOF.
- While consumption pauses, at most limit values may have left input without reaching the consumer. Include the worker's held value.
- Choose the largest safe output capacity; explain its calculation and blocking point.
- Close output after EOF and forwarding completion. Calls have independent state/channels.

## Constraints
- Input is non-nil, carries at most 12 integers, then closes; limit is 1–4.
- Caller eventually drains output; no cancellation. Standard library only. Worker structure and reasoning require method review.
