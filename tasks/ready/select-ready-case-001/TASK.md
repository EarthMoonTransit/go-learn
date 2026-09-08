# Select a ready value cancellation or default result

## Context

Complete the specified operations.

## API

```go
func TryReceive(input <-chan int, cancel <-chan struct{}) (int, string)
```

## Requirements

- Perform one nonblocking select: a received input value returns that value and value; a closed empty input returns zero and closed; ready cancellation returns zero and canceled; otherwise return zero and pending. Result labels are literal strings.
- Nil channels disable their cases. Do not close or send to input channels. Consume at most one selected input value.
- If input and cancellation are both ready, either ready result is valid; no priority is promised. The default result is only valid when no case is ready.
- Buffered zeros and negative numbers remain values. Cancellation is signaled by a closed or readable channel.

## Constraints

- Use only the Go standard library.
