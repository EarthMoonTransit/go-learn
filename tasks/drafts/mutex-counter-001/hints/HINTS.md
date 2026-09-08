# Progressive hints

1. Each Counter owns the mutex that protects its value.
2. Lock before the read-modify-write operation; unlock after it.
3. Waiting for goroutines later does not synchronize their earlier writes with one another.
