# Repair matching named and anonymous struct values

## Context

Repair the supplied operation.

## API

```go
type Contact struct{Name,Email string;Age int}
func BuildContacts(name, email string, age int) (Contact, struct { Name, Email string; Age int })
```

## Requirements

- Construct a Contact and an anonymous struct with the declared three fields, both storing the inputs exactly.
- Empty names/emails, spaces, Unicode and age zero are valid; age is 0..150.
- Returned values are independent value copies. Changing either result does not affect the other or later calls.
- Do not swap same-typed fields or discard the anonymous result.

## Constraints

- Use only the Go standard library.
