# Progressive hints

1. Both examples call the same API. You may write individual prefix arguments after the original values slice, or expand a prefix slice with `prefix...`. The supplied `slices.Equal` compares the resulting elements. Passing `prefix` without `...` does not expand it into variadic arguments.
2. Distinct prefix values expose reversal or truncation. Also check a call with no prefix arguments; expanding an empty prefix supplies zero arguments.
3. Extend the cases to nil and empty original input. Check the required prefix-then-original order. With zero total elements, nil and an empty result are both valid; the contract does not require a particular allocation strategy.
