# Progressive hints

1. Check interface nil separately from the concrete pointer.
2. A comma-ok assertion lets you inspect the supported pointer type safely.
3. Other implementations still dispatch through the interface.
