# Progressive hints

1. A channel buffer can hold the finite snapshot before any receiver runs.
2. Expose only the receiving direction in the returned type.
3. Close after queueing every value; no producer goroutine is needed.
