# Progressive hints

1. WithCancel preserves the parent cancellation relationship and values.
2. The worker should wait on the child Done channel.
3. Close completion only when that wait ends; always return the derived cancel function.
