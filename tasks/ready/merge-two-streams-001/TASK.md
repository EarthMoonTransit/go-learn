# Merge two finite input streams

## Context

Merge two finite integer channels while each input can become ready independently.

## API

```go
func MergeTwo(left, right <-chan int) <-chan int
```

## Requirements

- Return a non-nil output channel and forward every input occurrence exactly once.
- Preserve each input's internal order; no total order between inputs is required.
- Make progress from either input without waiting for the other to produce or close.
- Use a select loop; disable a closed input by assigning nil to its local channel variable.
- Close output after both inputs close; emit no post-close zero values.

## Constraints

- Inputs are distinct non-nil channels that eventually close. The consumer fully drains output; early abandonment is outside this contract. Do not close inputs. Preserve the API; use only the standard library. Public checks include the required select/nil constructs.
