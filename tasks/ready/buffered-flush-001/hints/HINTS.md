# Progressive hints

1. Small writes may remain only in the buffer until Flush.
2. Check errors from both writing and flushing.
3. Flushing does not transfer ownership of the underlying writer.
