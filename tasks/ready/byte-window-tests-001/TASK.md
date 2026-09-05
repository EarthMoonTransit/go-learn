# Test byte windows

## Context
Add tests in `subject/subject_test.go` for the provided `ByteWindow` API. The checker supplies the implementation; submit tests rather than implementing it.

## API
```go
func ByteWindow(s string, start, end int) (string, bool)
```

## Requirements
- Valid bounds satisfy `0 <= start <= end <= len(s)`.
- For valid bounds, return the bytes at offsets `[start, end)` and `true`.
- Empty windows are valid, including at the end of a string and in an empty string.
- Invalid bounds return an empty string and `false`, without clamping or reordering the bounds.
- Preserve the selected bytes exactly, including zero and bytes above 127; offsets count bytes, not characters.
- Your tests must accept the correct implementation and reject implementations that violate these rules.

## Constraints
- Use `package subject_test` and the supplied import. `\xNN` encodes one byte. Assert both returned values; table organization and diagnostic style are not graded.
