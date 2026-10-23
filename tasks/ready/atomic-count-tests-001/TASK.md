# Test atomic count updates

## Context
Write tests for atomic updates.

## API
```go
type Delta struct{Key string;Amount int}
var ErrCount=errors.New("invalid counts")
func ApplyCounts(counts map[string]int,deltas []Delta) error
```

## Requirements
- Apply deltas in order: missing keys start at zero; retain zero-valued keys and unrelated entries.
- If any intermediate count is negative, return original ErrCount and leave counts unchanged; otherwise commit updates and return nil. Preserve deltas.
- Author a named table covering empty updates, existing/missing keys, repeated keys, zero, and early/late rejection.
- Run every row as a serial t.Run child with fresh state. Use unique descriptive names.
- Author shared assertion helpers calling t.Helper before reporting. Pass each child t; compare exact errors, full maps and unchanged delta inputs with actual/expected diagnostics.
- Public checks compatibility; full checks defects. Review checks table, child isolation and helpers.

## Constraints
- Non-nil maps contain counts 0..100; at most 8 deltas have nonempty keys and amounts -10..10.
