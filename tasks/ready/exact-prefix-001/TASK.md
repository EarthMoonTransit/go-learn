# Remove one exact prefix

## Context
Repair `RemovePrefix` in `subject/subject.go`. The current trimming call confuses a prefix with a set of characters.

## API
```go
func RemovePrefix(s, prefix string) (string, bool)
```

## Requirements
- If `s` begins with the complete, exact byte sequence `prefix`, remove that sequence once and return `true`.
- Preserve the rest unchanged, even when it begins with another copy of `prefix`.
- If there is no match at the beginning, return the original `s` and `false`.
- An empty prefix matches every string, including the empty string, and removes nothing.
- Match case and whitespace exactly. Do not search elsewhere, trim a character set, or normalize the input.

## Constraints
- Keep the signature. Inputs may contain any bytes. Removing a prefix equal to the entire input returns an empty string with `true`.
