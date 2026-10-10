# Cancellation and operation lifetime

Read this topic after learning functions, goroutines, channels and the basic
select statement. Cancellation asks work to stop. Completion means the work has
actually stopped and released the resources it owns. Treat those as different
observations when designing a concurrent operation.

A context carries a cancellation signal across a call boundary. Code must observe
that signal, normally through Done or Err, and choose how to return. Cancel does
not forcibly interrupt arbitrary code. A blocked operation that cannot observe
cancellation can remain blocked even after the context is canceled.

The example is deliberately synchronous: it shows that signaling cancellation
and executing a response are separate steps, without giving the worker exercise's
start/stop implementation.

```go
package example

import (
    "context"
    "fmt"
)

func Example() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    acknowledged := false
    cancel()
    fmt.Println(ctx.Err() == context.Canceled, acknowledged)
    if ctx.Err() != nil {
        acknowledged = true
    }
    fmt.Println(acknowledged)
    // Output:
    // true false
    // true
}
```

For a real worker, identify who starts it, who requests cancellation and how the
caller waits for completion. Identify who sends on each channel and who may close
it. A sender blocked on a channel needs a specified exit path if its consumer
stops early. Closing a result channel and canceling a context have different
purposes; choose ownership explicitly rather than having every caller close
whatever channel it can access.

Call the cancellation function to release resources associated with a derived
context when the operation ends. For errors and partial results, follow the
exercise's stated policy. Cancellation is not permission to return an arbitrary
partial answer or hide an unrelated error.

Tests should establish the needed ordering with channels or other synchronization.
A random sleep does not prove that another goroutine reached a particular state.
A race detector can expose unsynchronized access; a passing race run alone does
not prove correct shutdown or freedom from logical deadlock.

Use the linked exercise's public Check after implementing its lifetime contract.
Returning to this page does not create an attempt, hint or completion. Advanced
cancellation composition belongs later in the route; this page does not move it
to the beginning of the course.
