# Progressive hints

1. A lock around writers does not protect an unlocked reader.
2. Hold the same mutex while reading the shared value.
3. Release the lock without changing the counter.
