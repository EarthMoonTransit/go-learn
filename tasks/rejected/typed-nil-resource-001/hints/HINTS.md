# Progressive hints

1. An interface value carries a dynamic type and a dynamic value.
2. Comparing the interface with `nil` only covers the case where both are absent.
3. Reflection can inspect nil-capable kinds, but `IsNil` is not valid for every kind.
