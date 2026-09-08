# Test final lines and scanner error propagation

## Context

Write tests for the specified contract.

## API

```go
func ScanLines(src io.Reader) ([]string, error)
```

## Requirements

- Collect lines using bufio.ScanLines semantics: remove newline and an optional preceding carriage return; include a final nonempty line without a newline.
- Preserve spaces, Unicode and interior empty lines. Empty input returns no lines; nil and empty result slices are equivalent.
- Return collected lines together with the scanner error, preserving a supplied reader error identity. Do not close src.
- Accept lines up to 2048 bytes. Reject a 10000-byte unbroken token with a non-nil error; exact buffer capacity is unspecified.

## Constraints

- Use only the Go standard library.
