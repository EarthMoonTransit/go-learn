# Write a buffered word stream

## Context
Normalize scanned words.

## API
```go
func WriteWords(dst io.Writer, src io.Reader) error
```

## Requirements
- Configure Scanner with ScanWords. Accept words up to 64 bytes; reject an unbroken 1024-byte word. Scanner capacity may vary.
- Own a fresh bufio.Writer with capacity at least 16 bytes, even when dst is already buffered. Flush only your buffer.
- For each scanned word, write its unchanged bytes followed by newline. Preserve order and the final unterminated word; do not collect all words first.
- Stop scanning after a buffered write fails. Attempt final Flush even after scanner failure, preserving the accepted prefix where writes succeed.
- Return the original write/flush error in preference to the scanner error; otherwise return the scanner error, or nil at clean EOF.
- Never close either stream. Method review checks scanner setup and buffer/flush ownership.

## Constraints
- Nonnil streams obey io contracts. Input is at most 16KiB of UTF-8 text.
