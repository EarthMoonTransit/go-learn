# Implement value and pointer speakers through one interface

## Context

Complete the specified operations.

## API

```go
type Speaker interface{Speak()string}
type English struct{}
type Chinese struct{}
func (English) Speak() string
func (*Chinese) Speak() string
func SayHello(s Speaker) string
func BuiltinGreetings() [2]string
```

## Requirements

- English.Speak returns "Hello World"; (*Chinese).Speak returns "你好世界".
- SayHello calls Speak once on any supplied non-nil Speaker and returns its exact result.
- BuiltinGreetings returns English then Chinese output by using values or pointers that satisfy Speaker.
- Keep the declared receiver forms: English values and pointers qualify; Chinese requires a pointer. Calls are independent.

## Constraints

- Use only the Go standard library.
