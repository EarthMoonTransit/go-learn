# Progressive hints

1. A blocked callback must not block Start itself.
2. Receiving from a closed completion channel reports ok=false.
3. Release every test gate even on failure, and wait for successful workers to finish.
