# Progressive hints

1. A plain map lookup cannot distinguish absence from a stored zero value.
2. The two-result lookup form reports whether the key exists.
3. `delete` is safe for missing keys and nil maps; use the lookup result to decide what to return.
