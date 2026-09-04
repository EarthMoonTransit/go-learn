# Progressive hints

1. ASCII alone makes byte indexes and rune indexes look identical.
2. Include a valid multibyte sequence and compare exact offsets.
3. Construct invalid UTF-8 from a byte slice so the test source remains explicit.
