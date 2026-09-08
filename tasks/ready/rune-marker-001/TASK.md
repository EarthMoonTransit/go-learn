# Identify the first Unicode marker

## Context

Complete the specified Unicode operation.

## API

```go
func Application(log string) string
```

## Requirements

- Input is valid UTF-8; inspect Unicode code points from left to right.
- Map ❗ to recommendation, 🔍 to search and ☀ to weather.
- When multiple markers occur, return the result for the earliest marker.
- Return default for empty input or text without a recognized marker.

## Constraints

- Use only the Go standard library.
