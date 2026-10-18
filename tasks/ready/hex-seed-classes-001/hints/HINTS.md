# Progressive hints

1. `f.Add` registers concrete values; merely writing a branch in the callback does not exercise it.
2. Compare what the corpus exercises by length, byte range and order. Repeating one shape cannot establish variety.
3. Try short boundary values as well as a longer repeated value and an ordered sequence. Preserve the supplied value/status assertions; nil and empty represent the same byte value.
