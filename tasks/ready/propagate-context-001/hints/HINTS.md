# Progressive hints

1. The caller already owns the context lifetime.
2. Delegation can forward the same context and both return values directly.
3. An already canceled context is still passed to the operation.
