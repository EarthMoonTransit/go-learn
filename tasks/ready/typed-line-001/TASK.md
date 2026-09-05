# Format a status line

## Context
Repair `FormatStatus` in `subject/subject.go`. It must return one precisely formatted string from the supplied typed values.

## API
```go
func FormatStatus(label string, count int, enabled bool) string
```

## Requirements
- Return exactly `label=<label>; count=<count>; enabled=<enabled>`.
- Insert `label` unchanged: do not quote, escape, trim, or interpret its percent signs as formatting directives.
- Write `count` in decimal, keeping a minus sign when negative and adding no padding or plus sign.
- Write the boolean as lowercase `true` or `false`.
- Use exactly one space after each semicolon and add no newline or other suffix.

## Constraints
- Keep the signature. An empty label is valid. `fmt.Sprintf` returns formatted text; its verbs select how each argument is represented. No console output is needed.
