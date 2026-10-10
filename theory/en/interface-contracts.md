# Program against an interface contract

An interface lets a consumer request behavior without naming one concrete
implementation. A type satisfies a basic interface implicitly when its method
set supplies the required methods. There is no separate implements declaration.
Method names, parameter types and result types must match.

Start from the consumer: what is the smallest behavior it needs? A narrow
interface makes it easier to supply another implementation or a test double.
It also limits the assumptions the consumer may make. The consumer should not
need a switch over every concrete type that happened to exist when it was written.

This example uses a standard-library implementation of a small output contract:

```go
package example

import (
    "fmt"
    "strings"
)

type byteSink interface {
    WriteByte(byte) error
}

func Example() {
    var builder strings.Builder
    var sink byteSink = &builder
    err := sink.WriteByte('Q')
    fmt.Println(builder.String(), err)
    // Output: Q <nil>
}
```

The interface value holds a concrete pointer here. Pointer and value receiver
method sets differ; do not assume that assigning a value and its address to an
interface are interchangeable. Check the required methods rather than relying
on a similarly named field or a method with a different signature.

An interface value is nil only when it has neither a dynamic type nor a dynamic
value. A typed nil pointer stored in an interface does not make the interface
nil. Whether calling a method on such a pointer is valid depends on that method's
contract and implementation. Do not use `value != nil` as universal proof that a
method cannot encounter nil receiver state.

In the linked practice, compare the observable consumer contract with several
possible implementations. Assertions should verify calls and results required
by the task, not assume one struct layout. This page introduces the boundary;
it does not supply the consumer implementation or the learner's test cases.
