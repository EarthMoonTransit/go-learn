# Choose variable types

## Context

Edit subject/bindings.go. The fixed Snapshot function returns the four variables as values; you do not need to change or implement that function.

## API

```go
// Edit subject/bindings.go; Snapshot is provided.
func Snapshot() (any, any, any, any)
```

## Requirements

- count must have Go type int and value 12.
- scale must have Go type float64 and value 1.25.
- label must have Go type string and value "G".
- enabled must have Go type bool and value true.

## Constraints

- Keep the variable names. You may use explicit types or type inference. A single-quoted character and a double-quoted string have different types. Preparation: recognize a package declaration and open the module in GoLand. Snapshot's any return slots are test plumbing that preserves each variable's type.
