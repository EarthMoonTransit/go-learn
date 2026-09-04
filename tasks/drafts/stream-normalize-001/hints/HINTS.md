# Progressive hints

1. The fixed pipeline package only delegates; the editable transform owns stream behavior.
2. A scanner exposes completed tokens before reporting a later input error.
3. Check both the byte count and error returned by each write.
