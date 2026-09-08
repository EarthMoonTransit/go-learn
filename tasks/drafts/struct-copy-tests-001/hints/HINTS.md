# Progressive hints

1. The struct value and its slice backing array have different copy behavior.
2. Check fields first, then mutate through both slice views.
3. An empty slice can retain nonzero capacity.
