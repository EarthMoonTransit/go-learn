# Repair a UTF-8 prefix function

## Context

Repair `PrefixRunes` so its limit counts decoded runes while preserving the selected source bytes.

## API

```go
func PrefixRunes(text string, limit int) string
```

## Requirements

- Return at most the first `limit` decoded runes.
- Return an empty string when `limit <= 0`.
- Return all input bytes when the limit reaches or exceeds the rune count.
- Do not split a valid UTF-8 encoding.
- Treat each invalid UTF-8 byte as one decoded `utf8.RuneError`.
- Preserve original bytes, including invalid UTF-8 bytes, in the returned prefix.

## Constraints

- Do not normalize or replace input bytes.
- Use only the standard library.
