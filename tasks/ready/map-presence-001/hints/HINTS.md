# Progressive hints

1. A stored zero and a missing key can produce the same value in a one-result lookup.
2. Base the fallback decision on whether the key exists.
3. Reading from a nil map is allowed. This operation never needs to insert or delete anything.
