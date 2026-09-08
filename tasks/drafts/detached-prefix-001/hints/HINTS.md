# Progressive hints

1. A capacity limit does not detach the backing array.
2. Allocate the requested length and copy only the prefix.
3. Zero-length results should not keep spare capacity from the input.
