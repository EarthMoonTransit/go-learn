# Progressive hints

1. The valid-index and append cases have different effects on the slice.
2. An existing view must observe a valid replacement.
3. Append may return a new slice header; return that result for invalid indexes.
