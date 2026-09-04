# Construct zero values

## Context

Implement Defaults in subject/subject.go. Its five result slots are already declared; each call must return their exact Go zero values.

## API

```go
func Defaults() (int, bool, string, []int, map[string]int)
```

## Requirements

- The integer result is 0.
- The boolean result is false.
- The string result is empty.
- The slice result is nil, not an allocated empty slice.
- The map result is nil, not an allocated empty map.

## Constraints

- Keep the function signature. An uninitialized variable receives the zero value of its declared type. A nil slice and a nil map have length zero; length alone does not distinguish them from allocated empty containers. No input parsing or main function is needed.
