# Combine boolean conditions

## Context

Implement Allowed(enabled, blocked, override bool) bool.

## API

```go
func Allowed(enabled, blocked, override bool) bool
```

## Requirements

- An override always permits the operation.
- Without an override, a disabled gate rejects the operation.
- Without an override, a blocked operation is rejected.
- Without an override, an enabled, unblocked operation is permitted.

## Constraints

- All eight combinations of the three boolean inputs are valid. Return a boolean, not a string. Use the existing signature. Go writes logical AND as &&, OR as || and NOT as !. Parentheses can make grouping explicit. No input/output code is required.
