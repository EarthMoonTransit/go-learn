# Progressive hints

1. A string index is a byte offset, not a rune number.
2. UTF-8 decoding reports how many source bytes belong to the next decoded rune.
3. Slice the original string at the accumulated byte offset rather than rebuilding it.
