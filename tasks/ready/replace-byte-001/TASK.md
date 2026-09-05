# Replace one byte

## Context
Implement `ReplaceByte` in `subject/subject.go`. Strings contain immutable bytes; this function returns an edited value.

## API
```go
func ReplaceByte(s string, index int, value byte) (string, bool)
```

## Requirements
- A valid index satisfies `0 <= index < len(s)`; `len(s)` counts bytes.
- For a valid index, return a string with exactly that byte replaced and `true`.
- Keep every other byte in its original position, including repeated byte values.
- For an invalid index, return the original string and `false`; an empty string has no valid index.
- Accept every byte value, including zero and 255. Preserve all unaffected bytes without interpreting characters.
- Leave the caller's original string unchanged.

## Constraints
- Keep the signature. `byte` is an alias for `uint8`; `\xNN` denotes one byte in a string literal. No input parsing or console output is needed.
