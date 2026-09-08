# Progressive hints

1. A float32 rate carries its existing rounded value when widened.
2. Keep the input value in float64 throughout the multiplication.
3. Do not reconstruct a different rate from decimal literals.
