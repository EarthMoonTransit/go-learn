# Progressive hints

1. A slice's length belongs to the slice value, not directly to its backing array.
2. append returns the slice value that describes the result. It may reuse storage or allocate.
3. Check which slice value the function returns after calling append.
