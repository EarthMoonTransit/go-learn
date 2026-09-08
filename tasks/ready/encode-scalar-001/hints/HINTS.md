# Progressive hints

1. A rune is an integer type; not every value is a Unicode scalar.
2. Byte width is different from the number of code points.
3. Validate the scalar first, then encode it and measure the encoded string.
