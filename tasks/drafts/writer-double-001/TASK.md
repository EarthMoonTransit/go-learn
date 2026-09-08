# Build a recording writer with controlled partial failures

## Context

Complete the specified operations.

## API

```go
type RecordingWriter struct{limit int;failure error;calls int;data []byte}
func NewWriter(limit int, failure error) *RecordingWriter
func (w *RecordingWriter) Write(data []byte) (int, error)
func (w *RecordingWriter) Snapshot() ([]byte, int)
```

## Requirements

- NewWriter creates an empty independent writer with a nonnegative total byte limit and a supplied non-nil failure error.
- Each Write records one call and accepts the prefix fitting the remaining total capacity. If all bytes fit return count,nil, otherwise count and the original failure.
- An empty Write still counts and succeeds. Reaching the limit exactly succeeds; later nonempty writes fail with zero accepted bytes.
- Snapshot returns recorded bytes as an independent copy and the call count. Do not retain mutable input storage or share state between instances.

## Constraints

- Use only the Go standard library.
