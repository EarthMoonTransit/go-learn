# Test paired iota expressions and block boundaries

## Context

Write tests for Layout. Its specified declarations are:

```go
const (
    code0, mask0 = iota+10, uint8(1<<iota)
    _, _
    code2, mask2
    code3, mask3
)
const (
    restart0 = iota
    restart1
)
```

## API

```go
func Layout() ([3]int, [3]uint8, [2]int)
```

## Requirements

- The first array contains code0, code2 and code3, in that order.
- The second array contains mask0, mask2 and mask3, in that order.
- The final array contains restart0 and restart1, in that order.
- Check all positions, including the omitted declaration and the new const block's counter.

## Constraints

- Submit tests in package subject_test using the standard library. Only returned values are checked; tests must accept equivalent implementations. The declarations are supplied, not learner-authored.
