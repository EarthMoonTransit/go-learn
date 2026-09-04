# Progressive hints

1. Validate both bounds and their order before slicing.
2. Go permits a two-index slice high bound up to capacity. This Window API deliberately restricts it to the current length.
3. Two-index slicing keeps capacity through the original backing array. Allocating a copy changes the required sharing.
