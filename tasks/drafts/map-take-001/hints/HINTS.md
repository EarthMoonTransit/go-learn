# Progressive hints

1. A plain map lookup cannot distinguish absence from a stored zero value.
2. The two-result lookup form reports whether the key exists.
3. Deleting from a nil map is safe, but the existence decision still matters.
