# Integrate a streaming line transformer

## Context

Complete the editable transform package used by the fixed pipeline adapter.

## API

```go
func NormalizeLines(r io.Reader, w io.Writer) error
func pipeline.Run(r io.Reader, w io.Writer) error
```

## Requirements

- Read newline-delimited input incrementally.
- Trim surrounding Unicode whitespace from each line.
- Omit lines that become empty.
- Uppercase retained lines and write each with one trailing newline.
- Preserve output already written before a later read failure.
- Return read, write, and short-write errors without hiding them.

## Constraints

- Edit only `transform/normalize.go`.
- Each input line is at most 32 KiB.
- Do not use network, filesystem, or whole-input buffering.
