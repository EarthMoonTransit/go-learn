# Progressive hints

1. Register cleanup only after acquisition succeeds.
2. A defer would close too early when the helper returns.
3. Capture the resource belonging to each acquisition.
