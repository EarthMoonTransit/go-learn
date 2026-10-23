# Hints

- Reaching the byte cap ends the consumer even when the producer has more bytes.
- Keep the producer callback result separate from the error that interrupted its write.
- A consumer may finish before a blocked producer; decide which endpoint releases that write before joining.
