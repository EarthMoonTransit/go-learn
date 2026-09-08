# Progressive hints

1. The input index counts runes, not bytes.
2. Changing one occurrence should not replace other equal runes.
3. Use a separate rune representation, validate its index, then form the new string.
