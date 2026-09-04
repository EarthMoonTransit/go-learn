# Progressive hints

1. `append` may reuse the input backing array when capacity permits.
2. Copying only a slice header does not copy its elements.
3. Build independent storage before appending the additional values.
