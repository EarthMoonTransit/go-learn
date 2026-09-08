# Reset a record through its pointer receiver

## Context

Complete the specified operations.

## API

```go
type Entry struct{Name string;Age int;Attributes map[string]string}
func (e *Entry) Clear()
```

## Requirements

- Set Name to empty, Age to zero and Attributes to nil.
- Changes must be visible through every pointer alias to the same record.
- The previously referenced Attributes map itself remains unchanged.
- The receiver is non-nil; clearing an already-zero record is valid and repeatable.

## Constraints

- Use only the Go standard library.
