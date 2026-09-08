# Repair character-limit checks

## Context

Repair the supplied Unicode operation.

## API

```go
func WithinLimit(log string, limit int) bool
```

## Requirements

- Input is valid UTF-8; each Unicode code point counts once, including combining marks.
- Return true when the code-point count is at most limit.
- Empty input fits limit zero; no input fits a negative limit.
- The result depends on code points, not UTF-8 bytes or displayed grapheme clusters.

## Constraints

- Use only the Go standard library.
