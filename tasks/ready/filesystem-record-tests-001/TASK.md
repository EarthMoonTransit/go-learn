# Test directory records and file cleanup on every path

## Context

Write tests for the specified contract.

## API

```go
type Post struct{Title,Description string;Tags []string;Body string}
var ErrFormat=errors.New("invalid record format")
func ReadRecords(fileSystem fs.FS) ([]Post, error)
```

## Requirements

- Read root files in lexical name order, skipping directories. Each starts with Title: , Description: , Tags: headers and a fourth line ---.
- Strip those exact prefixes, split nonempty tags on comma-space, and join remaining ScanLines body lines with newline. Preserve other spaces; empty tags/body are allowed. Lines are at most 2048 bytes.
- Close every successfully opened file exactly once before returning. Read/format errors take precedence over close errors; otherwise return the original close error.
- Any directory/open/read/format/close failure returns nil records and the original error; malformed headers use the supplied non-nil ErrFormat. Preserve that sentinel. Empty directories return no records.

## Constraints

- Use only the Go standard library.
