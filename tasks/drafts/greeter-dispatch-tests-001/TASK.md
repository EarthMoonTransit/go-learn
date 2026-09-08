# Test composition through a two-method interface

## Context

Write tests for the specified contract.

## API

```go
type Greeter interface{LanguageName()string;Greet(string)string}
func SayHello(name string, g Greeter) string
```

## Requirements

- Return "I can speak " + g.LanguageName() + ": " + g.Greet(name).
- The non-nil Greeter may hold any implementing concrete type and state; dispatch both methods through its current dynamic value.
- Forward name exactly, including empty strings, Unicode and spaces. Preserve returned language and greeting text exactly.
- Write local implementing types in your tests and include different types and differently configured values of one type.

## Constraints

- Use only the Go standard library.
