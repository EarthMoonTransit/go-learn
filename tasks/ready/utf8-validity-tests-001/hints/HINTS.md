# Progressive hints

1. Valid UTF-8 includes multibyte characters and the replacement character.
2. Test malformed bytes after a valid prefix as well as at the beginning.
3. Separate truncation, overlong encodings, surrogates and out-of-range code points.
