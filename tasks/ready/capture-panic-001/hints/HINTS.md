# Progressive hints

1. Recover must execute directly in a deferred function on the panicking goroutine.
2. Named results can be updated by the deferred boundary.
3. A returned error is different from a panic.
