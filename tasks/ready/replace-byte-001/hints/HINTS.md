# Progressive hints

1. Separate index validation from the replacement. The upper bound is exclusive.
2. A string cannot be assigned through s[index]. A byte slice can hold an editable copy.
3. Convert the string to []byte, change the indexed element, and convert the copy back to string. A direct string(value) conversion interprets a numeric value as a Unicode code point, not an arbitrary byte.
