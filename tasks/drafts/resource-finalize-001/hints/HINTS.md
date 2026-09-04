# Progressive hints

1. Cleanup must be registered before any return caused by the write result.
2. A named return value lets deferred cleanup contribute an error.
3. `errors.Join` can preserve both classifications when write and close fail.
