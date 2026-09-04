# Repair tests for UTF-8 rune offsets

## Context

Strengthen the test submission for a function that reports source byte offsets of decoded runes.

## API

```go
func RuneOffsets(text string) []int
```

## Requirements

- Expect one byte offset for the start of each decoded rune.
- Expect a final offset equal to `len(text)`.
- Expect `[]int{0}` for empty input.
- Distinguish byte offsets from rune ordinals with multibyte UTF-8.
- Cover invalid UTF-8, where each invalid byte is one decoded `utf8.RuneError`.
- Detect every required mutant with deterministic tests.

## Constraints

- Submit tests in package `subject_test`.
- Do not normalize the input.
- Do not use timing, randomness, environment, network, or filesystem state.
