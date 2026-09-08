# Progressive hints

1. The replacement character also has a valid three-byte encoding.
2. Invalid decoding consumes one byte even when the intended sequence was longer.
3. Use both the decoded rune and the consumed width to distinguish the two cases.
