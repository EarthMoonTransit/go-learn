# Progressive hints

1. A variadic slice argument can share the caller backing array.
2. An expanded subslice still has its own first element.
3. Update an existing element rather than append.
