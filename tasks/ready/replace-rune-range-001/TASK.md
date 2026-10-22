# Replace a half-open Unicode rune range

## Context

Implement the supplied string operation.

## API

```go
func ReplaceRuneRange(s string, lo, hi int, replacement string) (string, bool)
```

## Requirements

- Valid bounds satisfy 0 <= lo <= hi <= the rune count of s.
- Replace the half-open range [lo,hi) with replacement and return the resulting string,true.
- Preserve all code points outside that range, in their original order.
- An empty range inserts, including at the end; an empty replacement removes the selected range.
- Indices count code points, not bytes or grapheme clusters. Combining code points count separately.
- Invalid bounds return the exact original string,false without panic. Empty strings are valid.

## Constraints

- Both strings are valid UTF-8. Use only the standard library; equivalent implementations are accepted.
