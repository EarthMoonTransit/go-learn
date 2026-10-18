# Progressive hints

1. If two functions use the same wrong convention, composing them can still recover the input. Observe the encoded bytes separately.
2. Expected digits must come from the input and the stated canonical mapping, rather than from Decode.
3. Compare with `hex.AppendEncode(nil, input)`, or check each byte's high and low hexadecimal digits using a lowercase digit table. Decode the subject's wire, not the oracle result, and retain the value/status assertions.
