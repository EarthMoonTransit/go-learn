# Test a private representation through exported operations

## Context

Write tests for the specified contract.

## API

```go
type Window struct{min,max int}
func NewWindow(min,max int)Window{return Window{min:min,max:max}}
func (w Window) Contains(value int) bool
```

## Requirements

- Write tests in package subject_test using only NewWindow and Contains; the private fields are not part of the external API.
- Contains reports membership in [min,max); constructor inputs satisfy min<=max and may be negative.
- An empty window contains nothing. Separate windows retain independent bounds.
- Test the public behavior at both boundaries without reproducing or reading its private storage.

## Constraints

- Use only the Go standard library.
