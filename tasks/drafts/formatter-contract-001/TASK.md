# Render values through an implicit interface

## Context

Implement a small formatting contract that accepts any type with the required method.

## API

```go
type Formatter interface {
    Format() string
}

type Pair struct {
    Left  string
    Right string
}

func (Pair) Format() string
func Render(value Formatter) string
```

## Requirements

- Format a `Pair` as `Left + ":" + Right`.
- Make the value type `Pair` satisfy `Formatter`.
- Have `Render` call the interface method.
- Accept caller-defined formatter types without registration or type switches.
- Return the formatter's result unchanged.

## Constraints

- Preserve the declared API.
- Do not add a closed list of supported concrete types.
- Use only the standard library.
