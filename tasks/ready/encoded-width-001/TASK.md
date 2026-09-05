# Repair switch dispatch

## Context

Repair the starter so each exact tag selects its own byte width.

## API

```go
func Width(kind string) (int, bool)
```

## Requirements

- The tags "byte", "word", "long" and "quad" return 1, 2, 4 and 8 respectively.
- Each recognized tag returns true as its second result.
- Every other string returns (0, false), including the empty string.
- Matching is exact: preserve case and whitespace distinctions.

## Constraints

- Edit subject/subject.go. Keep the signature. The starter uses a switch; Go does not fall through to the next case unless explicitly requested. Equivalent correct dispatch implementations are accepted.
