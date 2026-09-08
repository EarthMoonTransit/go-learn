# Progressive hints

1. Use controlled callback gates to show all workers start before any finishes.
2. WaitGroup count must include every worker before Wait can observe zero.
3. Join workers before checking their separately owned results.
