# Progressive hints

1. A declaration without an initializer provides the zero value automatically.
2. Check each result slot separately. Empty text and the text "0" are different values.
3. For slices and maps, an allocated empty value is not the zero value. Compare nilness, not just length.
